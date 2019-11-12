package storage

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/evan-buss/watch-together/metadata/scraper/data"
	"github.com/jmoiron/sqlx"

	// Use sqlite database driver
	_ "github.com/mattn/go-sqlite3"
)

// SQLite defines a generic structure for persisting data to an sqlite3 database
type SQLite struct {
	db         *sqlx.DB
	DataType   data.Parser //The datatype the database will expect, the *.db file is named after this
	Create     string      // The table generation query
	Insert     string      // The single row insertion query
	insertStmt *sqlx.Stmt
	mem        map[string]struct{}
	lastPurge  time.Time
}

// Init initializes the database
// Include your database creation statement as well as your insert statement
func (s *SQLite) Init() ([]string, error) {
	// We get the actual data type of the given data.Parser to use as DB file name
	name := reflect.TypeOf(s.DataType).String()
	name = name[strings.Index(name, ".")+1:]
	var err error
	s.db, err = sqlx.Open("sqlite3", name+".db")
	if err != nil {
		return nil, err
	}

	s.db.MustExec(s.Create)

	var urls []string
	s.db.Select(&urls, "SELECT url FROM movies")
	s.mem = make(map[string]struct{})
	for _, url := range urls {
		s.mem[url] = struct{}{}
	}

	fmt.Println("loaded", len(urls), "from db storage")

	var unvisited []string
	s.db.Select(&unvisited, "SELECT link FROM links WHERE link not in (SELECT url FROM movies);")
	fmt.Println("Found", len(unvisited), "unvisted links. Adding them to the queue.")

	s.insertStmt, err = s.db.Preparex(s.Insert)
	if err != nil {
		return nil, err
	}

	s.lastPurge = time.Now()

	return unvisited, nil
}

// Write inserts a single row into the database table
func (s *SQLite) Write(obj data.Parser) error {
	_, err := s.insertStmt.Exec(structToArray(obj)...)
	if err != nil {
		return err
	}

	for _, url := range obj.GetLinks() {
		_, err := s.db.Exec(`INSERT INTO links (url, link) VALUES (?,?)`, obj.GetKey(), url)
		if err != nil {
			return err
		}
	}

	// We want to purge duplicates every 5 minutes
	// TODO: TEST THIS THEN CHANGE IT
	if time.Now().Sub(s.lastPurge) > (time.Minute * 1) {
		log.Println("Purging duplicates from the database")
		// Remove any duplicates on close. Keeps the storage size down.
		s.db.Exec(`DELETE FROM links WHERE link in (SELECT url FROM movies);`)
		s.lastPurge = time.Now()
	}

	s.mem[obj.GetKey()] = struct{}{}
	return nil
}

// GetUnvisitedLinks queries the database for unvisited links
func (s *SQLite) GetUnvisitedLinks(links []string) []string {
	// We filter the urls to make sure they are unique
	output := make([]string, 0)

	for _, url := range links {
		_, pres := s.mem[url]
		if !pres {
			output = append(output, url)
		}
	}
	return output
}

// GetVisited returns true if the provided url has already been visited
func (s *SQLite) GetVisited(url string) bool {
	_, pres := s.mem[url]
	return pres
}

// Close cleans up and closes the database
func (s *SQLite) Close() {

	// Remove any duplicates on close. Keeps the storage size down.
	s.db.Exec(`DELETE FROM links WHERE link in (SELECT url FROM movies);`)

	err := s.db.Close()
	if err != nil {
		log.Println(err)
	}
}

// Create an array of interfaces from a struct's fields
func structToArray(obj data.Parser) []interface{} {
	values := make([]interface{}, 0)
	st := reflect.TypeOf(obj)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if field.Tag.Get("sql") != "-" {
			values = append(values, reflect.ValueOf(obj).Field(i).Interface())
		}
	}
	return values
}
