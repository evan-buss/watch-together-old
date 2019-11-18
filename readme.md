# Watch Together :camera:

> A lightweight video streaming platform. Watch movies with your friends.

Share your media library with a single command. Friends from all over the world can watch with you.

## Usage

### Initialize settings and scan library for movie files

`watch-together init`

This creates a new config directory at ~/.config/watch-together

Inside the directory a `.watch-together.toml` file will be created that persists config settings between runs

TOML Configuration Options (defaults)
```
# The port that the server will be hosted on. Forward this port for public access.
port = "8080"

# The directory to look for movies in
video-dir = "/home/[username]/Videos

# The name of the metadata database to create in your config directory
database = "library.db"
```

## Architecture

### Frontend

The frontend is written in SvelteJS with Tailwind CSS. The UI is bundled inside the binary.
The server host simply launches the binary, opens the port on their router, and sends other watchers their IP address.

### Server Package

The server is responsible for serving HTTP request from all clients. It is run with the binary. It serves the static assets as well as the HLS segments.

###  Video Video Packge

Video is responsible for all video data transcoding and library processing. Most of its functions are called by the host
via the website interface. 

### Movies Package

Metadata consists of an IMDB scraper and a simple REST microservice to serve query requests. When a user first launches
Watch Together, their local harddisk will be scanned for video files. Watch Together will try to contact the metadata 
API to retrieve imdb info including the poster, summary, rating, etc... If the server could not automatically find an appropriate match
users will be able to manually set the metdata via the web interface.

The movie database currently has ~200,000 movies within it.

## Goals
- [X] Stream HLS encoding videos
- [X] Live chat
- [X] Scan directory for movies and download metadata
- [ ] Live transode videos on the fly to HLS

## Technology

- Svelte
  - svelte-spa-router
- Tailwind CSS
- clappr
- Golang
  - cobra / viper (CLI and config)
  - gorilla/websocket
  - chi/router
- FFMPEG