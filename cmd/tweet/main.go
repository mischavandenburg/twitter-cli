package main

import "github.com/mischavandenburg/twitter-cli"

var text string = "Hello World! Testing a tweet using my cli command.\n#go #coding #study #learning"

func main() {
	twitter.Post(text)
}
