package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/robfig/cron"
	"github.com/sercanarga/go-twitter-modded/twitter"
	"log"
	"os"
	"os/signal"
	"github.com/sercanarga/go-twitter-banner-last-followers/modules"
)

type info struct {
	LastFollowerID int64
}

var (
	flags          = flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey    = flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret = flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken    = flags.String("access-token", "", "Twitter Access Token")
	accessSecret   = flags.String("access-secret", "", "Twitter Access Secret")
	i              = info{}
	jsonFile       = "info.json"
	bgFile         = "img/bg.jpg"
)

func main() {
	log.Println("Bot is up!")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(false),
	}
	var user, _, _ = client.Accounts.VerifyCredentials(verifyParams)

	cron := cron.New()
	cron.AddFunc("@every 1m", func() {
		var lastFollower, err = modules.GetLastFollowerID(*client, *user, 1)
		if err != nil {
			log.Fatal(err)
		}

		modules.ReadJson(jsonFile, &i)
		if lastFollower[0] != i.LastFollowerID {
			fmt.Println("New follower found")

			modules.WriteJson(jsonFile, info{
				LastFollowerID: lastFollower[0],
			})

			var profilePhotos = modules.GetFollowerPhotos(*client, user)
			err = modules.GenerateBanner(profilePhotos, bgFile)
			if err != nil {
				log.Fatal(err)
			}

			var Img, _ = os.ReadFile(bgFile)
			_, _, err = client.Accounts.UpdateProfileBannerPhoto(&twitter.AccountUpdateProfileBannerPhotoParams{
				Banner: base64.StdEncoding.EncodeToString(Img),
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	})
	cron.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutdown")
}
