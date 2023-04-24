package main

import (
	"os"
	"strings"

	"github.com/mischavandenburg/twitter-cli"
	"github.com/mischavandenburg/twitter-cli/internal/feed"
)

var tags string

func main() {
	if len(os.Args) > 1 {
		tags = "\n" + strings.Join(os.Args[1:], " ")
	}
	r := feed.GetLatestPost()
	tweet := r + tags
	// fmt.Println(tweet)
	twitter.Post(tweet)
}
