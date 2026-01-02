package server

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kamuridesu/go-anilist-ical/internal/anilist"
	"github.com/kamuridesu/go-anilist-ical/internal/ical"
	"github.com/kamuridesu/gomechan/core/response"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	rw := response.New(&w, r)

	paths := strings.Split(r.URL.Path, "/")
	username := ""
	includePlanning := false
	for _, path := range paths {
		if strings.HasSuffix(path, ".ics") {
			username = strings.ReplaceAll(strings.TrimSuffix(path, ".ics"), "/", "")
		}
	}

	if username == "" {
		rw.AsJson(http.StatusBadRequest, map[string]any{"error": "user cannot be empty"})
		return
	}

	if r.URL.Query().Has("include_planning") && r.URL.Query().Get("include_planning") == "true" {
		includePlanning = true
	}

	userMedia, err := anilist.GetUserCurrentAnimeSchedule(r.Context(), username, includePlanning)
	if err != nil {
		rw.AsJson(http.StatusBadRequest, map[string]any{"error": fmt.Sprintf("an error happened: %s", err)})
		return
	}

	ics := ical.New()

	for _, anime := range *userMedia {
		for _, ep := range anime.Episodes {
			ics.AddEvent(
				ep.AiringAt,
				ep.AiringAt.Add(time.Minute*30),
				fmt.Sprintf("%s Episode %d aired", anime.Title, ep.Number),
				"",
			)
		}
	}

	headers := map[string]string{
		"Content-Type":        "text/calendar; charset=utf-8",
		"Content-Disposition": "attachment; filename=anime.ics",
	}
	rw.SetHeaders(headers).Build(http.StatusOK, ics.Build()).Send()

}

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)

	slog.Info("Listening at 0.0.0.0:8080")
	err := http.ListenAndServe("0.0.0.0:8080", mux)

	if errors.Is(err, http.ErrServerClosed) {
		slog.Error("Server closed")
	} else if err != nil {
		slog.Error(fmt.Sprintf("Unknown error: %s", err))
		os.Exit(1)
	}

}
