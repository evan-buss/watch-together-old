## VIDEO

### Platforms
  - Linux:
    - Hardware acceleration isn't shipped with the binary. You need to compile it manually from what I understand
  - Windows:
    - Hardware acceleration works out of the box. 

### Convert a standard input file to HLS segments

`ffmpeg -i input.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8`

### 10 Bit x265 to 8 Bit x264

> If the video file uses 10 bit encoding, it doesn't work on Firefox but works on Chrome. We need to re-encode 8 bit x264 for HLS

`ffmpeg -i her.mkv -c:v libx264 -crf 18 -vf format=yuv420p -c:a copy -start_number 0 -hls_time 5 -hls_list_size 0 -hls_playlist_type event -f hls index.m3u8`


