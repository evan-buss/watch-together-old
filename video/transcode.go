package video

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Transcode begins FFMpeg HLS transcoding on the given movie file path
func Transcode(path string) {
	if path == "" {
		path = "/home/evan/Videos/treasure/treasure.mkv"
	}
	args := []string{"-i", path,
		"-profile:v", "high10",
		"-level", "3.0",
		"-start_number", "0",
		"-hls_time", "10",
		"-hls_list_size", "0",
		"-hls_playlist_type", "event",
		"-f", "hls",
		filepath.Join(os.Getenv("VIDEO_DIR"), "index.m3u8"),
	}

	transCmd := exec.Command("ffmpeg", args...)

	out, err := transCmd.CombinedOutput()
	if err != nil {
		log.Fatal(string(out))
	}
}
