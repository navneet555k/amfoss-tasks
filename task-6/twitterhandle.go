package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	var username string
	fmt.Println("Enter Twitter Handle.")
	fmt.Scanln(&username)
	tn := flag.String("twitter", username, "name of the twitter handle")

	flag.Parse()

	config := oauth1.NewConfig("", "")
	token := oauth1.NewToken("", "")
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	f, err := os.Create("followers.txt")

	// Name of Followers
	para := &twitter.FollowerListParams{ScreenName: *tn, Count: 90}
	followers, resp, err := client.Followers.List(para)
	var count int = 0
	fmt.Println(resp, err)
	f.WriteString("Followers - " + *tn)
	for _, follower := range followers.Users {
		count++
		f.WriteString("\n" + follower.ScreenName)
	}
	f.WriteString(fmt.Sprintf("\n", count))
	f.Close()
}
