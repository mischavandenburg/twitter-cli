package main

import (
	"github.com/mischavandenburg/twitter-cli"
	"github.com/mischavandenburg/twitter-cli/internal/feed"
)

func main() {
	r := feed.GetLatestPost()
	twitter.Post(r)
	// fmt.Println(r)
}
