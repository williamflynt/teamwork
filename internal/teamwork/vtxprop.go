package teamwork

import "time"

// VertexProperty is any arbitrary property on a Vertex that also implements Vertex.
type VertexProperty struct {
	Value string `json:"value"`
}

func NewVertexProperty(value string) *VertexProperty {
	return &VertexProperty{Value: value}
}

func (v VertexProperty) GetId() string {
	return v.Value
}

func (v VertexProperty) Type() string {
	return "vertexProperty"
}

func (v VertexProperty) Created() time.Time {
	return time.Unix(0, 0)
}

func (v VertexProperty) Modified() time.Time {
	return time.Unix(0, 0)
}

func (v VertexProperty) GetAttrs() map[string]interface{} {
	return nil
}
