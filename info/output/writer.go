package output

import (
	"github.com/evan-buss/watch-together/info/data"
)

// Writer is an interface to save data to a specific location
// Provides a flexibile interface to save in variable formats
type Writer interface {
	// Init performs any initialization before the data can be written. This must always be called first
	Init() error

	// WriteSingle writes a single data.Parser object
	WriteSingle(obj data.Parser) error

	// Close performs and shutdown tasks. This must always be called last
	Close()
}
