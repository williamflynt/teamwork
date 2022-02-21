package models

// Idable has a GetId method that returns a string.
type Idable interface {
	GetId() string
}

// Typer has a Type method that returns a string containing the type of thing it is.
type Typer interface {
	Type() string
}

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
