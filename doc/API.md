# REST API

## /

Base http path, this serves the front-end Svelte application. All routes are then handled server-side in the browser

## /ws

Client sends a request to `/ws` to upgrade the connection to a websocket. Websocket is used for real-time 
communication between connected clients. Namely:
  - Chat
  - Video Events
    - Pause / Play
    - Current Time
    - Seeking

## /media

The video stream media itself is handled through the `/media` path. The server is responsibile for serving files from 
the correct path. The client makes a request for the `index.m3u8` file and reads the segments from there as well.


## /library

`GET` the movies from the user's library. 

```json
{
  // The ID of the movie in the users's local metadata database (primary key)
  id: 12
  // The location of the movie on the user's system
  location: "/home/evan/Videos/...",
  // The id of the movie in the metadata database, -1 if not found
  metadata: 2312 
}
```

`UPDATE` a specific local database item to use new metadata

```json
{
  // The ID of the item to be updated
  id: 12
  // The new metadata database id
  metadata: 2313
}
```



# Websocket Messages

All messages take the form 
```json
{
  "event": "event_name",
  "data": {}
}
```
The client-side message handler registers callback functions for specific events


## Chat
```json
{
  "event": "message",
  "data": {
    "name": "username",
    "message": "message string"
  }
}
```