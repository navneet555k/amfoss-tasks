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

	config := oauth1.NewConfig("FaMHhif2wNxjcbzTKSX4jCem3", "OZ6M3O3e82Ms3MLuirDvUVqWkx1knaiSwjUmAOx5wJWu2OanMd")
	token := oauth1.NewToken("1212951844583903232-rWBbPESBpebIMOXRtkzmgw3suUZ0V5", "NmlJcogw86HBVKbqALtZO56jrhVRmEc4gbUUN5zlzVjT5")
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
