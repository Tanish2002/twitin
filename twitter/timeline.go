package twitter

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Tanish2002/twitin/db"
	"github.com/sivchari/gotwtr"
)

func init() {
	// Just a onetime check to see if db is empty, if yes then update it with most recent tweet's ID
	if db.GetRecentID() == "" {
		recentID, _ := getRecentTweet()
		db.AddRecentID(recentID)
	}
}

var (
	bearer_token     string = os.Getenv("BEARER_TOKEN")
	twitter_username string = os.Getenv("TWITTER_USERNAME")
)

func getClient() *gotwtr.Client {
	return gotwtr.New(bearer_token)
}

func TwitterTimer() {
	for range time.Tick(time.Second * 10) {
		CheckTweets()
	}
}

func CheckTweets() {
	id, tweet := getRecentTweet()
	fmt.Println("The tweet is: ", tweet)
	if db.GetRecentID() != id {
		fmt.Println(id)
		db.UpdateRecentID(id)
	}
}

func getRecentTweet() (string, *gotwtr.UserTweetTimelineResponse) {
	client := getClient()
	user, err := client.RetrieveSingleUserWithUserName(context.Background(), twitter_username)
	if err != nil {
		fmt.Println(err)
	}
	tweets, err := client.UserTweetTimeline(context.Background(), user.User.ID, &gotwtr.UserTweetTimelineOption{
		Expansions:  []gotwtr.Expansion{gotwtr.ExpansionAttachmentsMediaKeys},
		MediaFields: []gotwtr.MediaField{gotwtr.MediaFieldURL},
	})
	if err != nil {
		fmt.Println(err)
	}
	return tweets.Meta.NewestID, tweets
}
