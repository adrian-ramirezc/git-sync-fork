package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"fmt"
)

var rootCmd = &cobra.Command{
	Use:   "git-sync-fork",
	Short: "Sync your current github fork with the main repository",
	Long:  `Sync your current github fork with the main repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello %s!\n", cmd.Flag("name").Value.String())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("name", "n", "World", "Say hello")
}
