package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mischavandenburg/twitter-cli/internal/feed"
)

var tags string
var action string
var tweet string

func main() {
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	if action == "url" {
		url := os.Args[2]
		tweet = url + tags

		// tags have a different position if url is given
		tags = "\n" + strings.Join(os.Args[3:], " ")
		send()
	}

	if action == "latest" {
		latestpost := feed.GetLatestPost()

		// tags have a different position if url is given
		tags = "\n" + strings.Join(os.Args[2:], " ")
		tweet = latestpost + tags
		send()
	}

	fmt.Println("Usage: tweet url url tags or tweet latest tags")

}

func send() {
	fmt.Println(tweet)
	// twitter.Post(tweet)
	os.Exit(0)
}
