## SUBTITLES

### PGS Subtitles
> PGS Subtitles are image files, as such they cannot be converted to the text based VTT format and served in a live playlist like HLS. We therefore encode the subtitle directly into the video stream.

`ffmpeg -i treasure.mkv -c:v h264_nvenc -filter_complex "[0:v][0:s]overlay[v]" -map "[v]" -map 0:a:0 output.mkv`

### PGS Subtitles with HLS x265 Base 10 to x264 Base 8
> Putting that together with our base 10 to base 8 encoder command
 
ffmpeg -progress tcp://127.0.0.1:8082/ -i treasure.mkv -c:v h264_nvenc -filter_complex "[0:v][0:s]overlay,format=yuv420p[v]" -map "[v]" -map 0:a:0 -crf 18 -c:a copy -start_number 0 -hls_time 5 -hls_list_size 0 -hls_playlist_type event -f hls index.m3u8
