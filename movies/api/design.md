# API Design

## Get Specific Movie

`/id/{moviID}`

- Get the metadata of a specific movie by its ID

## Search for movies
`/?title=` 

- Search for movies by title
- Example:
  - `/?title=Interstellar`

`/?title=(?)&year=(?)`

- Search for movies by title and year
- _Note_: You cannot search by year only, as it is too big of a request to return anything meaningful.
- Example:
  - `/?title=Matrix&year=1999`

### Paginate Results
- All responses limit to movies 10 per request. You can get the next 10 using an offset query
- Example:
  - `/?title=Matrix&offset=1`

### Response Format
```json
{
  // Total number of database items that matched the query. We only get 10 at a time
  "total": 100,
  // 10 movie objects
  "movies:" []
}
```

