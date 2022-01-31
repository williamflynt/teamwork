package backends

import (
	"context"
	"github.com/autom8ter/dagger"
	"teamwork/internal/teamwork"
)

// daggerDb implements the Database interface on top of Autom8ter's Dagger library.
type daggerDb struct {
	Graph *dagger.Graph
}

func NewDagger() (*daggerDb, error) {
	return &daggerDb{Graph: dagger.NewGraph()}, nil
}

func (d daggerDb) CreateEdge(ctx context.Context, s, p teamwork.Vertex, o teamwork.Edge) (teamwork.Fetchable, error) {
	sPath := dagger.Path{
		XID:   s.GetId(),
		XType: s.Type(),
	}
	oPath := dagger.Path{
		XID:   o.GetId(),
		XType: o.Type(),
	}
	edgeAttrs := teamwork.GetAttrs(p)
	edgePath := dagger.Path{
		XID:   p.GetId(),
		XType: p.Type(),
	}
	edgeNode := dagger.Node{
		Path:       edgePath,
		Attributes: edgeAttrs,
	}
	e, err := d.Graph.SetEdge(sPath, oPath, edgeNode)
	if err != nil {
		return teamwork.Fetchable{}, nil
	}
	return teamwork.Fetchable{Id: e.XID, Type: e.XType}, nil
}

func (d daggerDb) CreateVertex(ctx context.Context, vtx teamwork.Vertex) (teamwork.Fetchable, error) {
	f := teamwork.GetFetchable(vtx)
	daggerPath := dagger.Path{
		XID:   f.Id,
		XType: f.Type,
	}
	v := d.Graph.SetNode(daggerPath, teamwork.GetAttrs(vtx))
	return teamwork.Fetchable{Id: v.XID, Type: v.XType}, nil
}
