# FFMPEG Conversion Recipes
> Our goal is to transcode a normal video source file to HLS on-the-fly

## There are a few different complexities to overcome

1. Input files vary in many different ways. We care about video codec, language, and subtitles.
2. Transcoding requires a speed of greater than 1x so the user does'nt have to wait for video to load. We should try to adjust quality levels to match their hardware and network speed.
3. HLS requires a constant feed of video segments to prevent buffering.
4. Our transcoder needs to start and stop on command. 
   - Do to the live nature, the user skipping ahead requires our transcoder to stop what it's doing and start transcoding at the requested time. We must invalidate the previously transcoded segments because segments are not deterministic. The content and length varies depending on the starting time. Splitting video is not clean.

<!-- TODO: Continue here. Go into detail about each of the methods and the ffmpeg queries -->
## VIDEO

## Convert an input file to HLS segments
ffmpeg -i input.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8

## 10 Bit x265 to 8 Bit x264
ffmpeg -i her.mkv -c:v libx264 -crf 18 -vf format=yuv420p -c:a copy -start_number 0 -hls_time 5 -hls_list_size 0 -hls_playlist_type event -f hls index.m3u8

With Subtitles

## AUDIO

## SUBTITLES

```
Notes: SRT is the preferred subtitle format for source files. It is text based an can easily be streamed via VTT
       PGS is an image based format. It has to be burned into the HLS stream *.ts files
```

### PGS Subtitles
ffmpeg -i treasure.mkv -c:v h264_nvenc -filter_complex "[0:v][0:s]overlay[v]" -map "[v]" -map 0:a:0 output.mkv
### PGS Subtitles with HLS x265 Base 10 to x264 Base 8
 ffmpeg -progress tcp://127.0.0.1:8082/ -i treasure.mkv -c:v h264_nvenc -filter_complex "[0:v][0:s]overlay,format
=yuv420p[v]" -map "[v]" -map 0:a:0 -crf 18 -c:a copy -start_number 0 -hls_time 5 -hls_list_size 0 -hls_playlist_ty
pe event -f hls index.m3u8

<!-- Working btw -->
ffmpeg -progress tcp://127.0.0.1:8082/ -i treasure.mkv -c:v h264_nvenc -filter_complex "[0:v][0:s]overlay,format =yuv420p[v]" -map "[v]" -map 0:a:0 -crf 18 -c:a copy -start_number 0 -hls_time 5 -hls_list_size 0 -hls_playlist_ty pe event -f hls index.m3u8

TODO:
  Limit bitrate
  Add progress watching with -progress
  Start working on seeking logic
    Ideas:
      User clicks seek location -> wait 3 seconds -> send request
      Serverdelete all old .ts files -> stop running command -> execute ffmpeg at new time -> *need to somehow reflect changes in the .m3u8 file...
