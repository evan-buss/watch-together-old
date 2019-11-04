package video

import (
	"fmt"
	"log"
	"os/exec"
)

// FFProbe runs the FFProbe utility on the given file
func FFProbe(path string) {
	if path == "" {
		path = "/home/evan/Downloads/legend.mp4"
	}
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", path)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
