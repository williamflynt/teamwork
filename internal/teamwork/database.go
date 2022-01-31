package teamwork

import (
	"context"
	"teamwork/internal/backends"
)

// Database defines the interface with the graph storage engine.
type Database interface {
	CreateEdge(ctx context.Context, s, p Vertex, o Edge) (Fetchable, error)
	CreateVertex(ctx context.Context, vtx Vertex) (Fetchable, error)
	GetEdge[T Edge](ctx context.Context, edge Fetchable) (*T, error)
	GetVertex[T Vertex](ctx context.Context, vtx Fetchable) (*T, error)
}

// NewDatabase returns a new Database using the specified underlying store.
func NewDatabase() (Database, error) {
	// TODO: Implement use of a connection string or similar. (wf 29 Jan 21)
	return backends.NewDagger()
}
