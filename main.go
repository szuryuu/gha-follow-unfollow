package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v74/github"
)

const (
	githubUsername = "szuryuu"
)

type GithubClient struct {
	gc *github.Client
}

func NewGithubClient() *GithubClient {
	token := os.Getenv("MY_PAT")
	client := github.NewClient(nil).WithAuthToken(token)
	return &GithubClient{gc: client}
}

func (gc *GithubClient) ListFollowers(ctx context.Context, username string, opts *github.ListOptions) ([]*github.User, *github.Response, error) {
	return gc.gc.Users.ListFollowers(ctx, username, opts)
}

func (gc *GithubClient) ListFollowing(ctx context.Context, username string, opts *github.ListOptions) ([]*github.User, *github.Response, error) {
	return gc.gc.Users.ListFollowing(ctx, username, opts)
}

func (gc *GithubClient) GetAllFollowers(ctx context.Context, username string) ([]*github.User, error) {
	var allFollowers []*github.User
	opts := &github.ListOptions{PerPage: 100}

	for {
		followers, resp, err := gc.ListFollowers(ctx, username, opts)
		if err != nil {
			return nil, err
		}
		allFollowers = append(allFollowers, followers...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allFollowers, nil
}

func (gc *GithubClient) GetAllFollowing(ctx context.Context, username string) ([]*github.User, error) {
	var allFollowing []*github.User
	opts := &github.ListOptions{PerPage: 100}

	for {
		following, resp, err := gc.ListFollowing(ctx, username, opts)
		if err != nil {
			return nil, err
		}
		allFollowing = append(allFollowing, following...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allFollowing, nil
}

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	ctx := context.Background()
	client := NewGithubClient()

	followers, err := client.GetAllFollowers(ctx, githubUsername)
	if err != nil {
		log.Fatalf("Failed to get followers: %v", err)
	}
	// log.Println("Followers:")
	// for _, follower := range followers {
	// 	log.Println(follower.GetLogin())
	// }

	following, err := client.GetAllFollowing(ctx, githubUsername)
	if err != nil {
		log.Fatalf("Failed to get following: %v", err)
	}
	// log.Println("Following:")
	// for _, following := range following {
	// 	log.Println(following.GetLogin())
	// }

	followingMap := make(map[string]bool)
	for _, f := range following {
		followingMap[f.GetLogin()] = true
	}

	var needFollow []string
	for _, f := range followers {
		if !followingMap[f.GetLogin()] {
			needFollow = append(needFollow, f.GetLogin())
		}
	}

	log.Println("Need to follow back (followers you don't follow):", len(needFollow))
	for _, user := range needFollow {
		log.Println(user)
	}

	var needUnfollow []string
	for _, f := range following {
		if !followingMap[f.GetLogin()] {
			needUnfollow = append(needUnfollow, f.GetLogin())
		}
	}

	log.Println("Need to unfollow (following you don't follow back):", len(needUnfollow))
	for _, user := range needUnfollow {
		log.Println(user)
	}

	log.Println("Followers:", len(followers))
	log.Println("Following:", len(following))
}
