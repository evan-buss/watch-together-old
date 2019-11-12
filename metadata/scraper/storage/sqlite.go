package storage

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/evan-buss/watch-together/metadata/scraper/data"
	"github.com/jmoiron/sqlx"

	// Use sqlite database driver
	_ "github.com/mattn/go-sqlite3"
)

// SQLite defines a generic structure for persisting data to an sqlite3 database
type SQLite struct {
	db              *sqlx.DB
	insertMovieStmt *sqlx.Stmt
	insertLinkStmt  *sqlx.Stmt
	purgeStmt       *sqlx.Stmt
	visited         map[string]struct{}
	lastPurge       time.Time
}

// Init initializes the database
// Include your database creation statement as well as your insert statement
func (s *SQLite) Init() ([]string, error) {
	// We get the actual data type of the given data.Parser to use as DB file name

	var err error
	s.db, err = sqlx.Open("sqlite3", "movies.db")
	if err != nil {
		return nil, err
	}

	// Create a table for results and a table for links
	s.db.MustExec(`CREATE TABLE IF NOT EXISTS movies (
		url TEXT PRIMARY KEY,
		title TEXT,
		year TEXT,
		rating TEXT,
		summary TEXT,
		poster TEXT
	);
	CREATE TABLE IF NOT EXISTS links (url, link TEXT, FOREIGN KEY(url) REFERENCES movies(url));`)

	// Load the previously saved links into memory so we can check that we don't visit them again
	var urls []string
	s.db.Select(&urls, "SELECT url FROM movies")
	s.visited = make(map[string]struct{})
	for _, url := range urls {
		s.visited[url] = struct{}{}
	}

	fmt.Println("DATABASE: CONTAINS", len(urls), " ITEMS")

	// Get any unvisted links and add them the parser queue
	var unvisited []string
	s.db.Select(&unvisited, "SELECT link FROM links WHERE link not in (SELECT url FROM movies);")
	fmt.Println("DATABSE: ", len(unvisited), "Unvisited Links Added to Queue")

	s.insertMovieStmt, err = s.db.Preparex(`
		INSERT OR IGNORE INTO movies (url, title, year, rating, summary, poster) 
		VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	s.insertLinkStmt, err = s.db.Preparex(`INSERT INTO links (url, link) VALUES (?,?)`)
	if err != nil {
		return nil, err
	}

	s.purgeStmt, err = s.db.Preparex(`DELETE FROM links WHERE link in (SELECT url FROM movies);`)
	if err != nil {
		return nil, err
	}

	s.lastPurge = time.Now()

	return unvisited, nil
}

// Write inserts a single row into the database table
func (s *SQLite) Write(obj data.Parser) error {

	_, err := s.insertMovieStmt.Exec(structToArray(obj)...)
	if err != nil {
		return err
	}
	s.visited[obj.GetKey()] = struct{}{}

	for _, url := range obj.GetLinks() {
		_, err := s.insertLinkStmt.Exec(obj.GetKey(), url)
		if err != nil {
			return err
		}
	}

	// We want to purge duplicates every 5 minutes to keep db size down
	if time.Now().Sub(s.lastPurge) > (time.Minute * 5) {
		log.Println("Purging duplicates from the database")
		// Remove any duplicates on close. Keeps the storage size down.
		_, err = s.purgeStmt.Exec()
		if err != nil {
			return err
		}
		s.lastPurge = time.Now()
	}

	return nil
}

// GetUnvisitedLinks queries the database for unvisited links
func (s *SQLite) GetUnvisitedLinks(links []string) []string {
	// We filter the urls to make sure they are unique
	output := make([]string, 0)

	for _, url := range links {
		_, pres := s.visited[url]
		if !pres {
			output = append(output, url)
		}
	}
	return output
}

// GetVisited returns true if the provided url has already been visited
func (s *SQLite) GetVisited(url string) bool {
	_, pres := s.visited[url]
	return pres
}

// Close cleans up and closes the database
func (s *SQLite) Close() {

	// Remove any duplicates on close. Keeps the storage size down.
	_, err := s.purgeStmt.Exec()
	if err != nil {
		log.Println(err)
	}
	err = s.db.Close()
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
