package main

import (
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")

	log = &logger{logrus.New()}
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("did you forget your keys? " + name)
	}
	return v
}

func tweetFeed() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	api.SetLogger(log)

	rand.Seed(time.Now().UnixNano())
	rss := "https://cdn.rawgit.com/freeCodeCampTO/fcc-motivation/5b3dde34/index.xml"

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rss)
	limit := len(feed.Items)
	pick := rand.Intn(limit)
	rssItem := feed.Items[pick]
	tweet := rssItem.Description + " #coding #programming #yyz"

	_, err := api.PostTweet(tweet, url.Values{})
	if err != nil {
		log.Critical(err)
	}
}

func main() {

	lambda.Start(tweetFeed)

}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
