package teamwork

import "time"

// Vertex is any entity in our graph representation for TeamWork.
type Vertex interface {
	Idable
	Typer
	Created() time.Time
	Modified() time.Time
}
