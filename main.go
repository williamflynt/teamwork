package main

import (
	"github.com/rs/zerolog/log"
	"os"
	"teamwork/internal/server"
	"teamwork/internal/teamwork"
)

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err).Msg("teamwork run returned with error")
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	tw, twErr := teamwork.New()
	if twErr != nil {
		return twErr
	}

	s, sErr := server.New(tw)
	if sErr != nil {
		return sErr
	}
	return s.Run()
}
