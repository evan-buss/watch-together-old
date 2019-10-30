# ffmpeg -hide_banner -hwaccel cuda -y -i "Treasure Planet (2002).mkv" \
#   -vf "pad=ceil(iw/2)*2:ceil(ih/2)*2" -c:a aac -ar 48000 -c:v h264 -profile:v high10 -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod  -b:v 800k -maxrate 856k -bufsize 1200k -b:a 96k -hls_segment_filename planet\/360p_%03d.ts planet\/360p.m3u8 \
#   -vf "pad=ceil(iw/2)*2:ceil(ih/2)*2" -c:a aac -ar 48000 -c:v h264 -profile:v high10 -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 1400k -maxrate 1498k -bufsize 2100k -b:a 128k -hls_segment_filename planet\/480p_%03d.ts planet\/480p.m3u8 \
#   -vf "pad=ceil(iw/2)*2:ceil(ih/2)*2" -c:a aac -ar 48000 -c:v h264 -profile:v high10 -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 2800k -maxrate 2996k -bufsize 4200k -b:a 128k -hls_segment_filename planet\/720p_%03d.ts planet\/720p.m3u8 \
#   -vf "pad=ceil(iw/2)*2:ceil(ih/2)*2" -c:a aac -ar 48000 -c:v h264 -profile:v high10 -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 5000k -maxrate 5350k -bufsize 7500k -b:a 192k -hls_segment_filename planet\/1080p_%03d.ts planet\/1080p.m3u8

ffmpeg -hwaccel cuvid  -i "Treasure Planet (2002).mkv" -vf "scale=w=1280:h=720:force_original_aspect_ratio=decrease,pad=ceil(iw/2)*2:ceil(ih/2)*2" -c:a aac -ar 48000 -b:a 128k -c:v h264 -profile:v high10 -crf 20 -g 48 -keyint_min 48 -sc_threshold 0 -b:v 2500k -maxrate 2675k -bufsize 3750k -hls_time 4 -hls_playlist_type vod -hls_segment_filename beach\\720p_%03d.ts beach\\720p.m3u8

# FFMPEG Conversion Recipes

## Convert an input file to HLS segments
ffmpeg -i input.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8

## 10 Bit x265 to 8 Bit x264
ffmpeg -i her.mkv -c:v libx264 -crf 18 -vf format=yuv420p -c:a copy -start_number 0 -hls_time 5 -hls_list_size 0 -hls_playlist_type event -f hls index.m3u8

TODO:
  Limit bitrate
  Add progress watching with -progress
  Start working on seeking logic
    Ideas:
      User clicks seek location -> wait 3 seconds -> send request
      Serverdelete all old .ts files -> stop running command -> execute ffmpeg at new time -> *need to somehow reflect changes in the .m3u8 file...
