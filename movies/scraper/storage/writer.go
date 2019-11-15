package storage

import (
	"github.com/evan-buss/watch-together/metadata/scraper/data"
)

// Writer is an interface to save data to a specific location
// Provides a flexibile interface to save in variable formats
type Writer interface {
	// Init performs any initialization before the data can be written. This is always be called first
	// If the storage has any links that have not been visted, add those back to the queue
	Init() ([]string, error)

	// Write writes a single data.Parser object
	Write(obj data.Parser) error

	GetQueue() []string

	// Close performs and shutdown tasks. This must always be called last
	Close()
}
