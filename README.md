# git-sync-fork

`git-sync-fork` is a lightweight Git plugin that lets users update their fork and pull those changes to the local repository with a simple `git sync-fork` command instead of manually syncing the fork repository and the pulling the changes.

## What it does

When you run `git sync-fork`, the tool:

1. determines the fork repository from `git remote get-url origin`
2. finds the default branch name from `git remote show origin`
3. calls GitHub's fork sync API on `POST /repos/{owner}/{repo}/merge-upstream`
4. updates the local repository by checking out the default branch and pulling `origin/<branch>`

That means your fork is synced with upstream and your local checkout is updated in one command.

## Install from the latest GitHub release

This repository publishes binary releases for Linux and Windows.

### Linux

1. Download the latest Linux release asset from GitHub.
2. Move it into a directory that is on your `PATH`, and rename it to `git-sync-fork`, for example:
   ```bash
   mv git-sync-fork-linux ~/.local/bin/git-sync-fork
   ```
3. Make sure it is executable:
   ```bash
   chmod +x ~/.local/bin/git-sync-fork
   ```
4. Run it as a git extension:
   ```bash
   git sync-fork
   ```

### Windows

1. Download the latest Windows release asset from GitHub.
2. Move it into a directory that is on your `PATH`, and rename it to `git-sync-fork.exe`, for example:
   Powershell
   ```powershell
   move .\git-sync-fork-windows.exe $env:USERPROFILE\git-sync-fork.exe
   ```
   
   Bash
   ```bash
   move git-sync-fork-windows.exe ~/git-sync-fork.exe
   ```

3. Run it as a git extension:
   ```powershell
   git sync-fork
   ```

## Notes

- The tool expects `origin` to be configured for your GitHub fork.
- By default it uses the fork's default branch as reported by `git remote show origin`.
- If the install directory is not already in your `PATH`, add it to your shell profile or environment variables before running `git sync-fork`.

## Usage

### Prerequisites

You need a Github fine-grained personal access access token with "Contents" repository permissions (write).
Create one [here](https://github.com/settings/personal-access-tokens/) following this [doc](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token).

Recommendations:
- In Repository access, select "All repositories"
- In Persmissions, add new permission "Contents" with "Read and write access".

### Linux

Set the token as an environment variable:
```bash
export GITHUB_TOKEN=<your_github_personal_access_token>
```

Then run:
```bash
git sync-fork
```

Example output:
```
Synced fork with upstream
Switching to branch main to update the local repository
Updated local repository
```

To persist the token across sessions, add it to your shell profile (e.g. `~/.bashrc`, `~/.zshrc`):
```bash
export GITHUB_TOKEN=<your_github_personal_access_token>
```

### Windows

Set the token as an environment variable:
```powershell
$env:GITHUB_TOKEN = "<your_github_personal_access_token>"
```

Then run:
```powershell
git sync-fork
```

Example output:
```
Synced fork with upstream
Switching to branch main to update the local repository
Updated local repository
```

To persist the token permanently, set it as a system environment variable:
1. Press `Win + X` and select "System"
2. Click "Advanced system settings"
3. Click "Environment Variables"
4. Click "New" under "User variables" or "System variables"
5. Variable name: `GITHUB_TOKEN`
6. Variable value: `<your_github_personal_access_token>`
7. Click OK and restart your terminal

### Verify installation

Check that the command is working:

**Linux or Windows:**
```bash
git sync-fork -h
```

Note: Git may intercept `git sync-fork --help` and try to show a man page, so use `-h` or run the binary directly for built-in help.
