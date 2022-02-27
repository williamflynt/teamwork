package teamwork

import (
	"context"
	"teamwork/internal/models"
)

// Database defines the interface with the graph storage engine.
type Database interface {
	CreateEdge(ctx context.Context, s, p Vertex, o Edge) (models.Fetchable, error)
	CreateVertex(ctx context.Context, vtx Vertex) (models.Fetchable, error)
	GetEdge(ctx context.Context, edge models.Fetchable) (Edge, error)
	GetVertex(ctx context.Context, vtx models.Fetchable) (Vertex, error)
}
