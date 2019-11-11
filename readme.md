# Watch Together

A lightweight video streaming platform. Watch movies with your friends.


## Goals
- [X] Stream HLS encoding videos
- [X] Live chat
- [ ] Live transode videos on the fly to HLS
- [ ] Open anywhere on your file system and serve the videos found there

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