package anilist

import (
	"context"
	"time"
)

type Episode struct {
	Number   int
	AiringAt time.Time
}
type AnimeSchedule struct {
	Title    string
	Episodes []Episode
}

func GetUserCurrentAnimeSchedule(ctx context.Context, username string, includePlanning bool) (*[]AnimeSchedule, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()
	user, err := GetUserFromName(ctx, "kamuridesu")
	if err != nil {
		return nil, err
	}

	mediaList, err := GetUserMediaList(ctx, user, includePlanning)
	if err != nil {
		return nil, err
	}

	var anime []AnimeSchedule
	for _, entry := range *mediaList {
		if entry.Media.Status == "FINISHED" {
			continue
		}
		var eps []Episode
		for _, air := range entry.Media.AiringSchedule.Nodes {
			airingAt := time.Unix(int64(air.AiringAt), 0)

			eps = append(eps, Episode{
				Number:   air.Episode,
				AiringAt: airingAt,
			})
		}
		anime = append(anime, AnimeSchedule{
			Episodes: eps,
			Title:    entry.Media.Title.UserPreferred,
		})
	}

	return &anime, nil
}
