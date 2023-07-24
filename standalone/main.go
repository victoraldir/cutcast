package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func createHLS(inputFile string, outputDir string, segmentDuration int) error {
	// Create the output directory if it does not exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Create the HLS playlist and segment the video using ffmpeg
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i", inputFile,
		"-profile:v", "baseline", // baseline profile is compatible with most devices
		"-level", "3.0",
		"-start_number", "0", // start numbering segments from 0
		"-hls_time", strconv.Itoa(segmentDuration), // duration of each segment in seconds
		"-hls_list_size", "0", // keep all segments in the playlist
		"-f", "hls",
		fmt.Sprintf("%s/playlist.m3u8", outputDir),
	)

	output, err := ffmpegCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create HLS: %v\nOutput: %s", err, string(output))
	}

	return nil
}

func main() {
	inputFile := "/tmp/c74f5760-5678-4ddc-b976-0c8a8b50af6c/myvideo.mp4.part"
	outputDir := "/tmp/c74f5760-5678-4ddc-b976-0c8a8b50af6c/output"
	segmentDuration := 10 // duration of each segment in seconds

	if err := createHLS(inputFile, outputDir, segmentDuration); err != nil {
		log.Fatalf("Error creating HLS: %v", err)
	}

	log.Println("HLS created successfully")
}
