package feed

import (
	"strings"

	"github.com/mmcdole/gofeed"
)

func GetLatestPost() string {
	fp := gofeed.NewParser()

	// TODO handle error
	feed, _ := fp.ParseURL("https://mischavandenburg.com/index.xml")
	slice := []string{feed.Items[0].Title, feed.Items[0].Link}
	result := strings.Join(slice, "\n")
	// fmt.Printf("Testing printf %q", result)
	return result
}
