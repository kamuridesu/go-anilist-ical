package anilist

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUserFromName(ctx context.Context, name string) (*User, error) {
	payload := map[string]any{
		"query": QueryGetUserId,
		"variables": map[string]string{
			"name": name,
		},
	}

	var userData UserInfoResponse
	err := sendRequest(ctx, payload, &userData)
	if err != nil {
		return nil, err
	}

	return &userData.Data.User, nil
}

func GetUserMediaList(ctx context.Context, user *User, includePlanning bool) (*[]MediaEntry, error) {
	statusIn := []string{"CURRENT"}
	if includePlanning {
		statusIn = append(statusIn, "PLANNING")
	}
	payload := map[string]any{
		"query": QueryGetMediaList,
		"variables": map[string]any{
			"usrId":    user.Id,
			"type":     "ANIME",
			"statusIn": statusIn,
		},
	}

	var mediaData MediaListResponse
	err := sendRequest(ctx, payload, &mediaData)
	if err != nil {
		return nil, err
	}

	return &mediaData.Data.MediaListCollection.Lists[0].Entries, nil
}

func sendRequest[T any](ctx context.Context, payload map[string]any, ds *T) error {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://graphql.anilist.co", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("error querying: status is %d and body is %s", res.StatusCode, string(resBody))
	}

	if err := json.Unmarshal(resBody, ds); err != nil {
		return err
	}
	return nil
}
