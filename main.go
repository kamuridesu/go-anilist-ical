package main

import (
	"os"

	"github.com/kamuridesu/go-anilist-ical/internal/server"
)

func init() {
	if _, err := os.Stat("/ca-certificates.crt"); err != nil {
		if os.IsNotExist(err) {
			return
		} else {
			panic(err)
		}
	}
	server.SetupTLS()
}

func main() {
	server.Start()
}
