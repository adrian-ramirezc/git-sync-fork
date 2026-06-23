# git-sync-fork

`git-sync-fork` is a lightweight Git plugin that lets users update their fork with a simple `git sync-fork` command instead of manually syncing the fork repository.

## Install from the latest GitHub release

This repository publishes binary releases for Linux and Windows.

### Linux

1. Download the latest Linux release asset from GitHub.
2. Rename the downloaded file to `git-sync-fork`.
3. Make it executable:
   ```bash
   chmod +x git-sync-fork
   ```
4. Move it into a directory that is on your `PATH`, for example:
   ```bash
   mv git-sync-fork ~/.local/bin/
   ```
5. Verify the destination is in your `PATH`:
   ```bash
   echo "$PATH" | tr ':' '\n'
   ```
6. Run it as a git extension:
   ```bash
   git sync-fork
   ```

### Windows

1. Download the latest Windows release asset from GitHub.
2. Rename the downloaded file to `git-sync-fork.exe`.
3. Move it into a directory that is on your `PATH`, for example:
   ```powershell
   move .\git-sync-fork.exe $env:USERPROFILE\bin\git-sync-fork.exe
   ```
4. Verify the destination is in your `PATH`:
   ```powershell
   $env:PATH -split ';'
   ```
5. Run it as a git extension:
   ```powershell
   git sync-fork
   ```

## Notes

- Pick a local install directory you own, like `~/.local/bin` on Linux or `%USERPROFILE%\bin` on Windows.
- If the chosen directory is not already in your PATH, add it to your shell profile or system environment variables before running `git sync-fork`.
