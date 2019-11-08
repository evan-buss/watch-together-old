package output

import (
	"reflect"
	"strings"

	"github.com/evan-buss/watch-together/info/scraper/data"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	db       *sqlx.DB
	DataType data.Parser
	Create   string
	Insert   string
}

// Init initializes the database
// Include your database creation statement as well as your insert statement
func (s *SQLite) Init() error {
	// We get the actual data type of the given data.Parser to use as DB File name
	name := reflect.TypeOf(s.DataType).String()
	name = name[strings.Index(name, ".")+1:]

	db, err := sqlx.Open("sqlite3", name+".db")
	if err != nil {
		return err
	}

	s.db = db

	s.db.MustExec(s.Create)
	return nil
}

func (s *SQLite) WriteSingle(obj data.Parser) error {
	stmt, err := s.db.Preparex(s.Insert)
	if err != nil {
		return err
	}

	_, err := stmt.Exec(structToArray(obj)...)
	if err != nil {
		return err
	}
	return nil
}

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

func (s *SQLite) WriteFull(objs []data.Parser) error { return nil }
func (s *SQLite) Close() {
	s.db.Close()
}
