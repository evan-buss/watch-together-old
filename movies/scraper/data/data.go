package data

import "io"

// Parser is used to implement custom site-specific parsing logic.
// Each Parser has only 2 rules for the contents of the struct
//  1) It should have some sort of unique string key
//  2) It should collect a slice of links to visit from each page
type Parser interface {
	GetKey() string     // Get the unique key
	GetLinks() []string // Get the list of links
	Parse(body *io.ReadCloser, url string) (Parser, error)
}
