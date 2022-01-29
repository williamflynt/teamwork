package teamwork

import (
	"teamwork/internal/database"
	"time"
)

// App is the interface for working with our TeamWork application.
type App interface{}

// Vertex is any entity in our graph representation for TeamWork.
type Vertex interface {
	Id() string
	CreatedUtc() time.Time
	ModifiedUtc() time.Time
}

// Edge is a connection between two Vertex. It's used to encode relationships.
type Edge interface {
	Id() string
	Subject() string
	Predicate() string
	Object() string
	Label() string
}

// app implements App.
type app struct {
	Db database.Database
}

// New returns a new instance of App.
func New() (App, error) {
	return app{Db: database.New()}, nil
}
