package output

import (
	"log"
	"reflect"
	"strings"

	"github.com/evan-buss/watch-together/info/data"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// SQLite defines a generic structure for persisting data to an sqlite3 database
type SQLite struct {
	db         *sqlx.DB
	DataType   data.Parser //The datatype the database will expect, the *.db file is named after this
	Create     string      // The table generation query
	Insert     string      // The single row insertion query
	insertStmt *sqlx.Stmt
}

// Init initializes the database
// Include your database creation statement as well as your insert statement
func (s *SQLite) Init() error {
	// We get the actual data type of the given data.Parser to use as DB file name
	name := reflect.TypeOf(s.DataType).String()
	name = name[strings.Index(name, ".")+1:]
	var err error
	s.db, err = sqlx.Open("sqlite3", name+".db")
	if err != nil {
		return err
	}

	s.db.MustExec(s.Create)

	s.insertStmt, err = s.db.Preparex(s.Insert)
	if err != nil {
		return err
	}

	return nil
}

// WriteSingle inserts a single row into the database table
func (s *SQLite) WriteSingle(obj data.Parser) error {
	_, err := s.insertStmt.Exec(structToArray(obj)...)
	if err != nil {
		return err
	}
	return nil
}

// Close cleans up and closes the database
func (s *SQLite) Close() {
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
