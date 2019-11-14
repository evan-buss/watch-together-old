We need to use `ffprobe` to detmine what settings are required for the `ffmpeg` transcode


## Output file info in JSON format
`ffprobe -v quiet -print_format json -show_format -show_streams test.txt`