# FFMPEG Conversion Recipes
> Our goal is to transcode a normal video source file to HLS on-the-fly

## There are a few different complexities to overcome

1. Input files vary in many different ways. We care about video codec, language, and subtitles.
   - We need a system that can read and parse the important bits from `ffprobe` so we run the right ffmpeg command
2. Transcoding requires a speed of greater than 1x so the user does'nt have to wait for video to load. We should try to adjust quality levels to match their hardware and network speed.
3. HLS requires a constant feed of video segments to prevent buffering.
4. Our transcoder needs to start and stop on command. 
   - Do to the live nature, the user skipping ahead requires our transcoder to stop what it's doing and start transcoding at the requested time. We must invalidate the previously transcoded segments because segments are not deterministic. The content and length varies depending on the starting time. Splitting video is not clean.


## Useful FFMPEG Flags

```
-v quiet
  Disable the verbose logging

-c:v h264_nvenc
  Use NVIDIA GPU to encode the video (3x faster)

-progress tcp://127.0.0.1:[port] 
  Send ffmpeg progress updates to the specified tcp server (See progress.go)

```

# Screenshots

For now I am using a simple video screenshot to show the video in the library selection screen

`ffmpeg -ss 01:23:45 -i [input] -vframes 1 -q:v 2 output.jpg`