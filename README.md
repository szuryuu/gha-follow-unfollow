# Go GitHub Follow/Unfollow Bot

[![Go Build](https://github.com/szuryuu/gha-follow-unfollow/actions/workflows/follow_unfollow.yml/badge.svg)](https://github.com/szuryuu/gha-follow-unfollow/actions/workflows/follow_unfollow.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A GitHub Action, written in Go, that automates managing your followers. It can unfollow users who don't follow you back and follow back users who have recently followed you.

## Features

-   **Unfollow Automation**: Automatically unfollows users who are not following you back.
-   **Follow Back Automation**: Automatically follows back any new followers.
-   **Scheduled & Manual Execution**: Runs on a daily schedule but can also be triggered manually.
-   **Pure Go Implementation**: No external dependencies besides the official Go GitHub client.
<!---   **Safe List**: A configurable list of users who will never be unfollowed.-->

## How to Use This Template

You can easily set up this bot for your own account by using this repository as a template.

### Step 1: Create Your Repository from this Template

Click the **"Use this template"** button at the top of this repository's page and select **"Create a new repository"**. Give your new repository a name and create it.

### Step 2: Generate a Personal Access Token (PAT)

This action requires a GitHub Personal Access Token with specific permissions to follow and unfollow users on your behalf.

1.  Go to your GitHub **Settings** > **Developer settings** > **Personal access tokens** > **Tokens (classic)**.
2.  Click **"Generate new token"** and select **"Generate new token (classic)"**.
3.  Give it a descriptive name (e.g., `FOLLOW_UNFOLLOW_ACTION_TOKEN`).
4.  Set an expiration date (recommended for security).
5.  Select the following scopes:
    -   `repo` (to allow updates to the repository)
    -   `user` (to allow following and unfollowing users)
6.  Click **"Generate token"** and **copy the token immediately**. You will not be able to see it again.

### Step 3: Configure Repository Secrets

Add the required secrets to your repository:

1. Go to your repository → **Settings** → **Secrets and variables** → **Actions**
2. Add these secrets by clicking **"New repository secret"**:

| Secret Name | Value | Description |
|-------------|--------|-------------|
| `MY_PAT` | Your PAT from Step 2 | GitHub Personal Access Token |
| `GIT_EMAIL` | `your.email@example.com` | Your GitHub email for commits |
| `GIT_NAME` | `Your Name` | Your display name for commits |

### Step 4: Update Configuration

1. Open `main.go` in your repository
2. Change the username on line 13:
   ```go
   githubUsername = "your-github-username"  // Replace with your username
   ```
3. Commit the changes:
   ```bash
   git add main.go
   git commit -m "Update username configuration"
   git push
   ```

## Running the Action

-   **Automatic Schedule**: The action is configured to run automatically every day.
-   **Manual Trigger**: You can run it manually at any time by going to the **Actions** tab in your repository, selecting **"Follow Unfollow Bot"** from the sidebar, and clicking the **"Run workflow"** button.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
