package anilist

type UserInfoResponse struct {
	Data UserInfo `json:"data"`
}

type UserInfo struct {
	User User `json:"User"`
}

type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type MediaListResponse struct {
	Data struct {
		MediaListCollection struct {
			Lists []MediaList `json:"lists"`
		} `json:"MediaListCollection"`
	} `json:"data"`
}

type MediaListCollection struct {
	Lists []MediaList `json:"data"`
}

type MediaList struct {
	Entries []MediaEntry `json:"entries"`
}

type MediaEntry struct {
	Media Media `json:"media"`
}

type Title struct {
	UserPreferred string `json:"userPreferred"`
}

type AiringSchedule struct {
	Nodes []AiringNode `json:"nodes"`
}

type AiringNode struct {
	AiringAt        int `json:"airingAt"`
	TimeUntilAiring int `json:"timeUntilAiring"`
	Episode         int `json:"episode"`
}

type Media struct {
	Title          Title          `json:"title"`
	AiringSchedule AiringSchedule `json:"airingSchedule"`
	Status         string         `json:"status"`
}
