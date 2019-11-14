package storage

import (
	"log"
	"reflect"
	"sync"
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
	lastPurge       time.Time
	mux             sync.Mutex
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
	CREATE TABLE IF NOT EXISTS links (url, link TEXT UNIQUE, FOREIGN KEY(url) REFERENCES movies(url));`)

	// Load the previously saved links into memory so we can check that we don't visit them again
	row := s.db.QueryRowx("SELECT COUNT(url) FROM movies;")
	var count string
	row.Scan(&count)
	log.Println("DATABASE: CONTAINS", count, "ITEMS")

	// Get any unvisted links and add them the parser queue
	var unvisited []string
	s.db.Select(&unvisited, "SELECT DISTINCT link FROM links WHERE link not in (SELECT url FROM movies) LIMIT 1000;")
	log.Println("DATABASE:", len(unvisited), "LINKS ADDED TO PARSE QUEUE")

	s.insertMovieStmt, err = s.db.Preparex(`
		INSERT OR IGNORE INTO movies (url, title, year, rating, summary, poster) 
		VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	// If we have a duplicate link we just ignore the insertion as it already exists
	s.insertLinkStmt, err = s.db.Preparex(`INSERT OR IGNORE INTO links (url, link) VALUES (?,?)`)
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
	s.mux.Lock()
	defer s.mux.Unlock()

	_, err := s.insertMovieStmt.Exec(structToArray(obj)...)
	if err != nil {
		return err
	}

	for _, url := range obj.GetLinks() {
		_, err := s.insertLinkStmt.Exec(obj.GetKey(), url)
		if err != nil {
			return err
		}
	}

	// We want to purge duplicates every 5 minutes to keep db size down
	// The goal is to keep the links table items unique and as small as possible
	if time.Now().Sub(s.lastPurge) > (time.Minute * 5) {
		log.Println("DATABASE: PURGING DUPLICATES")
		// Remove any duplicates on close. Keeps the storage size down.
		_, err = s.purgeStmt.Exec()
		if err != nil {
			return err
		}
		s.lastPurge = time.Now()
	}

	return nil
}

// GetQueue loads any unvisited links from the database
func (s *SQLite) GetQueue() []string {
	// Get any unvisted links and add them the parser queue
	var unvisited []string
	err := s.db.Select(&unvisited, "SELECT DISTINCT link FROM links WHERE link not in (SELECT url FROM movies) LIMIT 1000;")
	if err != nil {
		log.Println(err)
	}
	return unvisited
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
		if field.Tag.Get("db") != "-" {
			values = append(values, reflect.ValueOf(obj).Field(i).Interface())
		}
	}
	return values
}
