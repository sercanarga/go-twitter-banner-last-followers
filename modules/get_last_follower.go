package modules

import (
	"github.com/sercanarga/go-twitter-modded/twitter"
)

func GetLastFollowerID(client twitter.Client, user twitter.User, count int) ([]int64, error) {
	lastFollower, _, err := client.Followers.IDs(&twitter.FollowerIDParams{UserID: user.ID, Count: count})
	return lastFollower.IDs, err
}

func GetLastFollowerList(client twitter.Client, user twitter.User, count int) ([]twitter.User, error) {
	lastFollower, _, err := client.Followers.List(&twitter.FollowerListParams{UserID: user.ID, Count: count, SkipStatus: twitter.Bool(true)})
	return lastFollower.Users, err
}
