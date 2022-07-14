package modules

import (
	"github.com/sercanarga/go-twitter-modded/twitter"
	"log"
	"strings"
)

func GetFollowerPhotos(client twitter.Client, user *twitter.User) []string {
	var lastFollowers, err = GetLastFollowerList(client, *user, 5)
	if err != nil {
		log.Fatal(err)
	}

	var profilePhotos []string
	for _, follower := range lastFollowers {
		profilePhoto := strings.Replace(follower.ProfileImageURLHttps, "_normal", "", -1)
		profilePhotos = append(profilePhotos, profilePhoto)
	}

	return profilePhotos
}
