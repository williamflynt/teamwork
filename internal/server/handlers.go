package server

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func healthcheckHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`ok`))
		if err != nil {
			log.Error().Err(err).Msg("failed to write for healthcheck")
		}
		w.WriteHeader(http.StatusOK)
	}
}
