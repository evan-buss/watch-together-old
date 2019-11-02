# FFMPEG Conversion Recipes
> Our goal is to transcode a normal video source file to HLS on-the-fly

## There are a few different complexities to overcome

1. Input files vary in many different ways. We care about video codec, language, and subtitles.
   - We need a system that can read and parse the important bits from `ffprobe` so we run the right ffmpeg command
2. Transcoding requires a speed of greater than 1x so the user does'nt have to wait for video to load. We should try to adjust quality levels to match their hardware and network speed.
3. HLS requires a constant feed of video segments to prevent buffering.
4. Our transcoder needs to start and stop on command. 
   - Do to the live nature, the user skipping ahead requires our transcoder to stop what it's doing and start transcoding at the requested time. We must invalidate the previously transcoded segments because segments are not deterministic. The content and length varies depending on the starting time. Splitting video is not clean.

<!-- TODO: Continue here. Go into detail about each of the methods and the ffmpeg queries -->
## Useful FFMPEG Flags

```
-v quiet
  Disable the verbose logging

-c:v h264_nvenc
  Use NVIDIA GPU to encode the video (3x faster)

-progress tcp://127.0.0.1:[port] 
  Send ffmpeg progress updates to the specified tcp server (See progress.go)

```

## VIDEO

### Platforms
  - Linux:
    - Hardware acceleration isn't shipped with the binary. You need to compile it manually from what I understand
  - Windows:
    - Hardware acceleration works out of the box. 

### Convert a standard inpu file to HLS segments

`ffmpeg -i input.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8`

### 10 Bit x265 to 8 Bit x264

> If the video file uses 10 bit encoding, it doesn't work on Firefox but works on Chrome. We need to re-encode 8 bit x264 for HLS

`ffmpeg -i her.mkv -c:v libx264 -crf 18 -vf format=yuv420p -c:a copy -start_number 0 -hls_time 5 -hls_list_size 0 -hls_playlist_type event -f hls index.m3u8`

## AUDIO

## SUBTITLES

### PGS Subtitles
> PGS Subtitles are image files, as such they cannot be converted to the text based VTT format and served in a live playlist like HLS. We therefore encode the subtitle directly into the video stream.

`ffmpeg -i treasure.mkv -c:v h264_nvenc -filter_complex "[0:v][0:s]overlay[v]" -map "[v]" -map 0:a:0 output.mkv`

### PGS Subtitles with HLS x265 Base 10 to x264 Base 8
> Putting that together with our base 10 to base 8 encoder command
 
ffmpeg -progress tcp://127.0.0.1:8082/ -i treasure.mkv -c:v h264_nvenc -filter_complex "[0:v][0:s]overlay,format=yuv420p[v]" -map "[v]" -map 0:a:0 -crf 18 -c:a copy -start_number 0 -hls_time 5 -hls_list_size 0 -hls_playlist_type event -f hls index.m3u8

TODO:
  - Restructure so only th relevant bits are shown under each heading
  Limit bitrate
  Add progress watching with -progress
  Start working on seeking logic
    Ideas:
      User clicks seek location -> wait 3 seconds -> send request
      Serverdelete all old .ts files -> stop running command -> execute ffmpeg at new time -> *need to somehow reflect changes in the .m3u8 file...
