package teamwork

import (
	"github.com/rs/zerolog/log"
	"github.com/segmentio/ksuid"
	"teamwork/internal/models"
	"time"
)

// Edge is a connection between two Vertex. It's used to encode relationships.
type Edge interface {
	models.Idable
	models.Typer
	Subject() models.Fetchable
	Predicate() string
	Object() models.Fetchable
}

// edge implements Edge and Vertex.
type edge struct {
	Id          string           `json:"id"`
	CreatedUtc  time.Time        `json:"created_utc"`
	ModifiedUtc time.Time        `json:"modified_utc"`
	Subject_    models.Fetchable `json:"subject"`
	Predicate_  string           `json:"predicate"`
	Object_     models.Fetchable `json:"object"`
}

func NewEdge(s, o models.Fetchable, predicate string, options ...Option) *edge {
	e, err := withOptions(new(edge), options...)
	if err != nil {
		log.Error().Err(err).Msg("failed to apply options on new edge")
	}
	if e.Id == "" {
		e.Id = ksuid.New().String()
	}
	if e.CreatedUtc.IsZero() {
		e.CreatedUtc = time.Now()
	}
	if e.ModifiedUtc.IsZero() {
		e.ModifiedUtc = time.Now()
	}
	e.Subject_ = s
	e.Object_ = o
	e.Predicate_ = predicate
	return e
}

// --- IMPLEMENT EDGE ---

func (e edge) GetId() string {
	return e.Id
}

func (e edge) Subject() models.Fetchable {
	return e.Subject_
}

func (e edge) Predicate() string {
	return e.Predicate_
}

func (e edge) Object() models.Fetchable {
	return e.Object_
}

// --- IMPLEMENT VERTEX ---

func (e edge) Type() string {
	return e.Predicate_
}

func (e edge) Created() time.Time {
	return e.CreatedUtc
}

func (e edge) Modified() time.Time {
	return e.ModifiedUtc
}
