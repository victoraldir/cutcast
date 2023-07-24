package adapters

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/victoraldir/cutcast/internal/app/record/domain"
)

const (
	VideoFileName = "myvideo.mp4"
)

type RecordFileFFMPEGRepository struct {
	cmd *exec.Cmd
	mu  *sync.Mutex
}

func NewRecordFileFFMPEGRepository() RecordFileFFMPEGRepository {
	return RecordFileFFMPEGRepository{
		cmd: nil,
		mu:  &sync.Mutex{},
	}
}

func (r RecordFileFFMPEGRepository) Create(done <-chan struct{}, recordCh <-chan domain.Record, mediaPath string) error {

	startCmd := func(url string, videoPath string) error {
		r.mu.Lock()
		defer r.mu.Unlock()

		if r.cmd != nil {
			return fmt.Errorf("download already in progress")
		}

		r.cmd = exec.Command(
			"youtube-dl",
			"-q",
			"--write-info-json",
			"--hls-use-mpegts",
			"--hls-prefer-ffmpeg",
			"-o", videoPath, url)
		r.cmd.Stdout = os.Stdout
		r.cmd.Stderr = os.Stderr

		return r.cmd.Start()
	}

	stopCmd := func() error {
		r.mu.Lock()
		defer r.mu.Unlock()

		if r.cmd == nil {
			return fmt.Errorf("no download in progress")
		}

		if err := r.cmd.Process.Signal(os.Interrupt); err != nil {
			return err
		}

		r.cmd = nil

		return nil
	}

	go func() {
		for {
			select {
			case <-done:
				if err := stopCmd(); err != nil {
					fmt.Println(err)
				}
			case record := <-recordCh:

				url := record.Url
				videoPath := fmt.Sprintf("%s/%s", mediaPath, VideoFileName)

				if err := startCmd(url, videoPath); err != nil {
					fmt.Println(err)
				}

			}
		}
	}()

	return nil
}

func (r RecordFileFFMPEGRepository) Trim(id string, trim domain.Trim, mediaDir string) (trimmedPath *string, err error) {

	// Create folder to save the trimmed video
	trimmedVideoPath := fmt.Sprintf("%s/%s", mediaDir, trim.GetStartEndTimeFormatted())

	if _, err := os.Stat(trimmedVideoPath); os.IsNotExist(err) {
		if err := os.MkdirAll(trimmedVideoPath, 0755); err != nil {
			return nil, err
		}
	}

	// Trim video
	cmd := exec.Command(
		"ffmpeg",
		// "-i", fmt.Sprintf("/tmp/%s/myvideo.mp4.part", id),
		"-i", fmt.Sprintf("%s/myvideo.mp4.part", mediaDir),
		"-ss", trim.StartTime,
		"-to", trim.EndTime,
		"-c", "copy", fmt.Sprintf("%s/%s", trimmedVideoPath, VideoFileName))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	// Create thumbnail
	cmd = exec.Command(
		"ffmpeg",
		// "-i", fmt.Sprintf("%s/myvideo.mp4", trimmedVideoPath),
		"-i", fmt.Sprintf("%s/myvideo.mp4.part", mediaDir),
		"-ss", "00:00:01.000",
		"-vframes", "1",
		fmt.Sprintf("%s/thumbnail.jpg", trimmedVideoPath))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return &trimmedVideoPath, nil
}

func (r RecordFileFFMPEGRepository) CreateHLS(mediaPath string, segmentDuration int) error {
	// Create the output directory if it does not exist
	if err := os.MkdirAll(mediaPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	inputFile := fmt.Sprintf("%s/%s.part", mediaPath, VideoFileName)

	createHLSCmd := func() error {
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
			fmt.Sprintf("%s/playlist.m3u8", mediaPath),
		)

		output, err := ffmpegCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to create HLS: %v\nOutput: %s", err, string(output))
		}

		return nil
	}

	// Execute go routine
	go func() {

		// Wait a few seconds to start the .part file creation
		time.Sleep(10 * time.Second)

		if err := createHLSCmd(); err != nil {
			fmt.Println(err)
		}
	}()

	return nil
}
