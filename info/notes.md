Design

Main Thread


Workers
- Process a single url
- Precipitate the data to a channel
- Need some way to signal done and some way to precipitate URLS


Goals
- URL lists grow exponentially. We need some way to save the current url and not parse it if the queue is full.


How Do I Manage the Massive List of URLS without using too much memory?
