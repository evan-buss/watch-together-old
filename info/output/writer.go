package output

import (
	"github.com/evan-buss/watch-together/info/data"
)

// Writer is an interface to save data to a specific location
// Provides a flexibile interface to save in variable formats
type Writer interface {
	// Init performs any initialization before the data can be written. This must always be called first
	Init() error

	// Write writes a single data.Parser object
	Write(obj data.Parser) error

	GetUnvisitedLinks([]string) []string

	// Close performs and shutdown tasks. This must always be called last
	Close()
}
