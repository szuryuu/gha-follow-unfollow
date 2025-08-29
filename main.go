package main

import (
	"context"
	"fmt"
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

func main() {
	ctx := context.Background()
	client := NewGithubClient()
	followers, _, err := client.ListFollowers(ctx, githubUsername, &github.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, follower := range followers {
		fmt.Println(follower.GetLogin())
		log.Println(follower.GetLogin())
	}

	following, _, err := client.ListFollowing(ctx, githubUsername, &github.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, following := range following {
		fmt.Println(following.GetLogin())
		log.Println(following.GetLogin())
	}
}
