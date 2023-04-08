package main

import (
	"context"
	"fmt"
	"os"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func main() {
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv("TWITTER_ACCESS_TOKEN"),
		OAuthTokenSecret:     os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.CreateInput{
		Text: gotwi.String("Hello World! This time I'm using os.Getenv in my Go program to load the credentials.\n#go #coding #study #learning"),
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
}
