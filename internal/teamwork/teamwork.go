package teamwork

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
)

// App is the interface for working with our TeamWork application.
type App interface {
	AddEntity(ctx context.Context, vtx Vertex) (Vertex, error)                     // AddEntity inserts a new Vertex to the graph.
	AddProperty(ctx context.Context, vtx Vertex, name, value string) (Edge, error) // AddProperty adds the specified property value to the Vertex in the form of an Edge, where name is the edge/relation type and value is the Vertex id. The resulting Edge will have `Id` of `name`, and the new Vertex will have `Value` of `value`.
	Link(ctx context.Context, s Vertex, o Vertex, p string) (Edge, error)          // Link creates an Edge between two Vertex in the graph.
}

// app implements App.
type app struct {
	Db Database
}

// New returns a new instance of App.
func New() (App, error) {
	db, err := NewDatabase()
	if err != nil {
		return nil, err
	}
	return &app{Db: db}, nil
}

// --- INTERFACES ---

// Idable has a GetId method that returns a string.
type Idable interface {
	GetId() string
}

// Typer has a Type method that returns a string containing the type of thing it is.
type Typer interface {
	Type() string
}

// --- GENERIC FUNCTIONS ---

func GetAttrs[T Vertex | Edge](t T) map[string]interface{} {
	b, err := json.Marshal(t)
	if err != nil {
		log.Error().Err(err).Msg("failed to get attributes for graph insertion on marshal")
		return nil
	}
	attrs := new(map[string]interface{})
	if uErr := json.Unmarshal(b, attrs); uErr != nil {
		log.Error().Err(err).Msg("failed to get attributes for graph insertion on unmarshal")
		return nil
	}
	return *attrs
}

// --- IMPLEMENT APP ---

func (a *app) Link(ctx context.Context, s Vertex, o Vertex, linkType string) (Edge, error) {
	if s.GetId() == "" {
		return nil, errors.New("tried to link a subject without an Id - create the Vertex first")
	}
	if o.GetId() == "" {
		return nil, errors.New("tried to link an object without an Id - create the Vertex first")
	}
	if s.GetId() == o.GetId() {
		return nil, errors.New("could not create edge - subject and object have the same id")
	}
	e := NewEdge(s, o, linkType)
	eF, err := a.Db.CreateEdge(ctx, s, o, e)
	if err != nil {
		return nil, err
	}
	if eF.GetId() == "" {
		return nil, errors.New("failed to insert edge during Link - no edge ID returned")
	}
	return e, err
}

func (a *app) AddEntity(ctx context.Context, vtx Vertex) (Vertex, error) {
	if vtx.GetId() == "" {
		return nil, errors.New("tried to add a Vertex without an Id - set an Id first")
	}
	vF, err := a.Db.CreateVertex(ctx, vtx)
	if err != nil {
		return nil, err
	}
	if vF.GetId() != vtx.GetId() {
		return vtx, errors.New("output and input id do not match for new vertex")
	}
	return vtx, nil
}

func (a *app) AddProperty(ctx context.Context, vtx Vertex, name, value string) (Edge, error) {
	if name == "" {
		return nil, errors.New("name may not be empty")
	}
	if value == "" {
		return nil, errors.New("value may not be empty")
	}
	propertyVtx := NewVertexProperty(value)
	return a.Link(ctx, vtx, propertyVtx, name)
}
