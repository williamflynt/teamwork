package main

import (
	"github.com/rs/zerolog/log"
	"os"
	"teamwork/internal/backends"
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
	// TODO: Get database type and config from our application config. (wf 5 Feb 22)
	db, err := backends.NewDagger()
	if err != nil {
		return err
	}
	tw, twErr := teamwork.New(db)
	if twErr != nil {
		return twErr
	}

	s, sErr := server.New(tw)
	if sErr != nil {
		return sErr
	}
	return s.Run()
}
