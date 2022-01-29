package database

import "github.com/autom8ter/dagger"

// daggerDb implements the Database interface on top of Autom8ter's Dagger.
type daggerDb struct {
	Graph *dagger.Graph
}

func newDagger() (*daggerDb, error) {
	return &daggerDb{Graph: dagger.NewGraph()}, nil
}
