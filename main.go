package main

import (
	"fmt"

	"github.com/google/go-github/v74/github"
)

func main() {
	fmt.Println("Init")
	client := github.NewClient(nil)
	fmt.Print(client)
}
