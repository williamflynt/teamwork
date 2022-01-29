package database

// Database defines the interface with the graph storage engine.
type Database interface{}

// New returns a new Database using the specified underlying store.
func New() (Database, error) {
	// TODO: Implement use of a connection string or similar. (wf 29 Jan 21)
	return newDagger(), nil
}
