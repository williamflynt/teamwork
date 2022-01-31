package teamwork

import (
	"context"
	"teamwork/internal/backends"
)

// Database defines the interface with the graph storage engine.
type Database interface {
	CreateEdge(ctx context.Context, s, p Vertex, o Edge) (Fetchable, error)
	CreateVertex(ctx context.Context, vtx Vertex) (Fetchable, error)
	GetEdge(ctx context.Context, edge Fetchable) (Edge, error)
	GetVertex(ctx context.Context, vtx Fetchable) (Vertex, error)
}

// NewDatabase returns a new Database using the specified underlying store.
func NewDatabase() (Database, error) {
	// TODO: Implement use of a connection string or similar. (wf 29 Jan 21)
	return backends.NewDagger()
}

// --- EXPORTED HELPERS ---

// Fetchable is something that can return values for an Id and Type.
type Fetchable interface {
	Idable
	Typer
}

// fetchable allows us to get a Vertex or Edge from the database using Id and Type. The database backend may use type
// as the graph label in a quad, or as a table name in a SQL database, or even disregard it.
// It implements Fetchable.
type fetchable struct {
	Id    string `json:"id"`   // Id is the unique identifier for the Vertex.
	Type_ string `json:"type"` // Type is the type of Vertex, which may be used as SQL table, label in a quad, or other.
}

func NewFetchable(id, type_ string) *fetchable {
	if id == "" && type_ == "" {
		return nil
	}
	return &fetchable{
		Id:    id,
		Type_: type_,
	}
}

func (f fetchable) GetId() string {
	return f.Id
}

func (f fetchable) Type() string {
	return f.Type_
}
