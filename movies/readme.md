# Movie Metadata API

The metadata api scrapes data from imdb and exposes the data via a REST API

Goals
- URL lists grow exponentially. We need some way to save the current url and not parse it if the queue is full.

How Do I Manage the Massive List of URLS without using too much memory?

TODO:
  - Might use the db only instead of the map. This would ensure no duplicates and reduce the memory overhead a bit
  - Implement output mode flag
  - Might implement some of the things I've seen other packages used (more robust queue system)

