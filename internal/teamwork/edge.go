package teamwork

import (
	"github.com/segmentio/ksuid"
	"time"
)

// Edge is a connection between two Vertex. It's used to encode relationships.
type Edge interface {
	Idable
	Typer
	Subject() Fetchable
	Predicate() string
	Object() Fetchable
}

// edge implements Edge and Vertex.
type edge struct {
	Id          string    `json:"id"`
	CreatedUtc  time.Time `json:"created_utc"`
	ModifiedUtc time.Time `json:"modified_utc"`
	Subject_    Fetchable `json:"subject"`
	Predicate_  string    `json:"predicate"`
	Object_     Fetchable `json:"object"`
}

func NewEdge(s, o Vertex, predicate string) *edge {
	return &edge{
		Id:          ksuid.New().String(),
		CreatedUtc:  time.Now(),
		ModifiedUtc: time.Now(),
		Subject_:    GetFetchable(s),
		Predicate_:  predicate,
		Object_:     GetFetchable(o),
	}
}

// --- IMPLEMENT EDGE ---

func (e edge) GetId() string {
	return e.Id
}

func (e edge) Subject() Fetchable {
	return e.Subject_
}

func (e edge) Predicate() string {
	return e.Predicate_
}

func (e edge) Object() Fetchable {
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
