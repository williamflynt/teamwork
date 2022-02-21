package backends

import (
	"context"
	"errors"
	"github.com/autom8ter/dagger"
	"teamwork/internal/models"
	"teamwork/internal/teamwork"
)

// daggerDb implements the Database interface on top of Autom8ter's Dagger library.
type daggerDb struct {
	Graph *dagger.Graph
}

func NewDagger() (*daggerDb, error) {
	return &daggerDb{Graph: dagger.NewGraph()}, nil
}

func (d daggerDb) CreateEdge(ctx context.Context, s, p teamwork.Vertex, o teamwork.Edge) (models.Fetchable, error) {
	// Dagger represents edges as Nodes, with attributes.
	edgeNode := dagger.Node{
		Path:       toDaggerPath(p),
		Attributes: teamwork.GetAttrs(p),
	}
	e, err := d.Graph.SetEdge(toDaggerPath(s), toDaggerPath(o), edgeNode)
	if err != nil {
		return models.NewFetchable("", ""), err
	}
	return toTeamworkFetchable(e.Path), nil
}

func (d daggerDb) CreateVertex(ctx context.Context, vtx teamwork.Vertex) (models.Fetchable, error) {
	v := d.Graph.SetNode(toDaggerPath(vtx), teamwork.GetAttrs(vtx))
	return toTeamworkFetchable(v.Path), nil
}

func (d daggerDb) GetEdge(ctx context.Context, edge models.Fetchable) (teamwork.Edge, error) {
	e, ok := d.Graph.GetEdge(toDaggerPath(edge))
	if !ok {
		return nil, errors.New("could not find requested edge")
	}
	return teamwork.NewEdge(
		toTeamworkFetchable(e.From),
		toTeamworkFetchable(e.To),
		e.XType,
	), nil
}

func (d daggerDb) GetVertex(ctx context.Context, vtx models.Fetchable) (teamwork.Vertex, error) {
	v, ok := d.Graph.GetNode(toDaggerPath(vtx))
	if !ok {
		return nil, errors.New("could not find requested vertex")
	}
	options := make([]teamwork.Option, 0)
	for k, val := range v.Attributes {
		options = append(options, map[string]any{k: val})
	}
	return teamwork.NewGenericVertex(v.XID, v.XType, options...), nil
}

// --- HELPERS ---

func toDaggerPath(f models.Fetchable) dagger.Path {
	return dagger.Path{
		XID:   f.GetId(),
		XType: f.Type(),
	}
}

func toTeamworkFetchable(p dagger.Path) models.Fetchable {
	return models.NewFetchable(p.XID, p.XType)
}
