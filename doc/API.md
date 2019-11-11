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

TODO: This should show the user's library. The library is determined by scanning the base directory and all underlying 
directories for movie files. If a file is encountered, the data type is parsed and saved to some sort of persistance.


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