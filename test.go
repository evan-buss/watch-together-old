package main

import (
	"fmt"
	"log"
	"os/exec"
)

func StartConvert() {

	args := []string{"-i", "/home/evan/Videos/treasure/treasure.mkv",
		"-profile:v", "high10",
		"-level", "3.0",
		"-start_number", "0",
		"-hls_time", "10",
		"-hls_list_size", "0",
		"-hls_playlist_type", "event",
		"-f", "hls",
		"/home/evan/Videos/treasure/index.m3u8",
	}

	transCmd := exec.Command("ffmpeg", args...)

	out, err := transCmd.CombinedOutput()
	if err != nil {
		log.Fatal(out)
	}

	fmt.Println("DONE CONVERSION")

	// test := []string{"~/Videos/legend/legend.mp4"}

	// out, err := exec.Command("ffmpeg", "-i", "~/Downloads/legend.mp4", "sample.mp4").CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("date failed: %v %v", err, string(out))
	// }
	// fmt.Printf("The date is %s\n", out)
	// bin, _ := exec.LookPath("echo")
	// fmt.Println(bin)

	// cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", "/home/evan/Downloads/legend.mp4")

	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Println("ERROR")
	// 	// log.Println(out)
	// 	// log.Fatal(err)
	// }

	// var encoded interface{}
	// fmt.Print(string(out))
	// if err := json.Unmarshal(out, &encoded); err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v", encoded)

	// TODO: This works but is hacky as all hell...
	// format := encoded.(map[string]interface{})["format"]
	// bitRate := format.(map[string]interface{})["bit_rate"]

	// fmt.Println("\n\n")
	// fmt.Println(bitRate)

	// out, _ := cmd.StdoutPipe()
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// scanner := bufio.NewScanner(out)

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// 	if scanner.Err() != nil {
	// 		log.Fatal(scanner.Err())
	// 	}
	// }

	// cmd.Wait()
}
