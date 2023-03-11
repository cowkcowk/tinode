// Package adapter contains the interfaces to be implemented by the database adapter
package adapter

// Adapter is the interface that must be implemented by a database
// adapter. The current schema supports a single connection by database type.
type Adapter interface {
	// General

	// Open and configure the adapter
	Open(config json.RawMessage) error
	
}