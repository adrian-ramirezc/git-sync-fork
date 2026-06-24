package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git sync-fork",
	Short: "Sync your current github fork with upstream and update the local repository.",
	Long:  `Sync your current github fork with upstream and update the local repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		syncFork()
		pullLocalBranch()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

func runGit(args ...string) (string, error) {
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return "", fmt.Errorf("git %v: %w", args, err)
	}
	return strings.TrimSpace(string(out)), nil
}

func getMainBranchName() string {
	origin, err := runGit("remote", "show", "origin")
	if err != nil {
		panic(err)
	}

	branch := parseHeadBranch(origin)
	if branch == "" {
		panic("Could not determine the main branch name")
	}

	return branch
}

func parseHeadBranch(origin string) string {
	for _, line := range strings.Split(origin, "\n") {
		if strings.Contains(line, "HEAD branch:") {
			return strings.TrimSpace(strings.Split(line, ":")[1])
		}
	}
	return ""
}

func getOwnerAndForkRepoName() string {
	originURL, err := runGit("remote", "get-url", "origin")
	if err != nil {
		panic(err)
	}
	originURL = strings.TrimSpace(originURL)
	originURL = strings.TrimSuffix(originURL, ".git")

	if after, ok := strings.CutPrefix(originURL, "https://github.com/"); ok {
		return after
	}
	panic(fmt.Sprintf("Unsupported origin URL: %s", originURL))
}

func syncFork() {
	mainBranch := getMainBranchName()
	forkRepo := getOwnerAndForkRepoName()

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		panic("GITHUB_TOKEN is required")
	}

	req := createMergeUpstreamHttpRequest(forkRepo, mainBranch, token)
	executeHttpRequest(req)
	fmt.Printf("Synced fork with upstream\n")
}

func pullLocalBranch() {
	mainBranch := getMainBranchName()
	currentBranch, err := runGit("branch", "--show-current")
	if err != nil {
		panic(err)
	}

	if currentBranch != mainBranch {
		fmt.Printf("Switching to branch %s to update the local repository\n", mainBranch)
		if _, err := runGit("checkout", mainBranch); err != nil {
			panic(err)
		}
	}

	if _, err := runGit("pull", "origin", mainBranch); err != nil {
		panic(err)
	}
	fmt.Printf("Updated local repository\n")
}

func createMergeUpstreamHttpRequest(forkRepo, mainBranch, token string) *http.Request {
	body, err := json.Marshal(map[string]string{
		"branch": mainBranch,
	})
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("https://api.github.com/repos/%s/merge-upstream", forkRepo),
		bytes.NewReader(body),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Add("X-GitHub-Api-Version", "2026-03-10")
	req.Header.Set("Content-Type", "application/json")

	return req
}

func executeHttpRequest(req *http.Request) *http.Response {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		panic(fmt.Errorf("GitHub API returned %s", resp.Status))
	}
	return resp
}
