# Movie Metadata 

Package consists of an IMDB scraper and a REST server.

## Scraper

- The scraper programmatically traverse and scrapes movie data from IMDB
- All data is saved to an SQLite3 databse file. (This might change soon)


## API

- The API connects to the database and queries data in response to specific HTTP requests
- See `design.md` in the `api` package for API handler documentation


