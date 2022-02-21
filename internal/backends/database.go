package backends

import (
	"context"
	"teamwork/internal/models"
	"teamwork/internal/teamwork"
)

// Database defines the interface with the graph storage engine.
type Database interface {
	CreateEdge(ctx context.Context, s, p teamwork.Vertex, o teamwork.Edge) (models.Fetchable, error)
	CreateVertex(ctx context.Context, vtx teamwork.Vertex) (models.Fetchable, error)
	GetEdge(ctx context.Context, edge models.Fetchable) (teamwork.Edge, error)
	GetVertex(ctx context.Context, vtx models.Fetchable) (teamwork.Vertex, error)
}

// NewDatabase returns a new Database using the specified underlying store.
func NewDatabase() (Database, error) {
	// TODO: Implement use of a connection string or similar. (wf 29 Jan 21)
	return NewDagger()
}
