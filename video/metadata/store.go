package metadata

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	// Use sqlite database driver
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

// create or open the local metadata database
func createStore() {
	path := filepath.Join(filepath.Dir(viper.ConfigFileUsed()), viper.GetString("database"))

	fmt.Println("Saving results to database at", path)

	var err error
	db, err = sqlx.Open("sqlite3", path)
	if err != nil {
		log.Fatal("could not open your movie library metdata database")
	}

	// Create a table that maps a movie on disk with a movie metdata results in our api
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS movies(	
		location TEXT,
		metadata INT
	)`)
	if err != nil {
		log.Fatal("could not create table")
	}
}
