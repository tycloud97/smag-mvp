package models

import (
	"time"

	"github.com/codeuniversity/smag-mvp/utils"
)

// TwitterUserRaw holds the follow graph info, only relating userNames
type TwitterUserRaw struct {
	// Meta
	ID   string `json:"id"`
	URL  string `json:"url"`
	Type string `json:"type"`

	// User info
	Name            string `json:"name"`
	Username        string `json:"username"`
	Bio             string `json:"bio"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`

	// Profile stats
	Location   string `json:"location"`
	JoinDate   string `json:"join_date"`
	JoinTime   string `json:"join_time"`
	IsPrivate  int    `json:"is_private"`
	IsVerified int    `json:"is_verified"`

	// Follows
	Following     int      `json:"following"`
	FollowingList []string `json:"following_list"`
	Followers     int      `json:"followers"`
	FollowersList []string `json:"followers_list"`

	// Usage stats
	Tweets     int `json:"tweets"`
	Likes      int `json:"likes"`
	MediaCount int `json:"media_count"`
}

// TwitterUser holds the follow graph info, only relating userNames
type TwitterUser struct {
	GormModelWithoutID

	// Meta
	ID   string
	URL  string
	Type string

	// User info
	Name            string
	Username        string
	Bio             string
	Avatar          string
	BackgroundImage string

	// Profile stats
	Location   string
	JoinDate   time.Time
	IsPrivate  bool
	IsVerified bool

	// Follows
	Following     int
	FollowingList []*TwitterUser
	Followers     int
	FollowersList []*TwitterUser

	// Usage stats
	Tweets     int
	Likes      int
	MediaCount int
}

func convertTwitterUser(raw *TwitterUserRaw) *TwitterUser {
	followingList := make([]*TwitterUser, len(raw.FollowingList))
	followersList := make([]*TwitterUser, len(raw.FollowersList))

	for index, item := range raw.FollowingList {
		followingList[index] = &TwitterUser{
			Username: item,
			//^^^ check if item is username using scraper output
		}
	}

	for index, item := range raw.FollowingList {
		followingList[index] = &TwitterUser{
			Username: item,
			//^^^ check if item is username using scraper output
		}
	}

	//joinDate := time.Unix() TODO

	isPrivate := utils.ConvertIntToBool(raw.IsPrivate)
	isVerified := utils.ConvertIntToBool(raw.IsVerified)

	return &TwitterUser{
		ID:   raw.ID,
		URL:  raw.URL,
		Type: raw.Type,

		Name:            raw.Name,
		Username:        raw.Username,
		Bio:             raw.Bio,
		Avatar:          raw.Avatar,
		BackgroundImage: raw.BackgroundImage,

		Location: raw.Location,
		//JoinDate: joinDate,
		IsPrivate:  isPrivate,
		IsVerified: isVerified,

		Following:     raw.Following,
		FollowingList: followingList,
		Followers:     raw.Followers,
		FollowersList: followersList,

		Tweets:     raw.Tweets,
		Likes:      raw.Likes,
		MediaCount: raw.MediaCount,
	}
}
