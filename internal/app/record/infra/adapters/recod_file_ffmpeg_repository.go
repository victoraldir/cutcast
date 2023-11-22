package adapters

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/victoraldir/cutcast/internal/app/record/domain"
	"github.com/victoraldir/cutcast/pkg/command"
)

const (
	VideoFileName = "myvideo.mp4"
)

type RecordFileFFMPEGRepository struct {
	// commandExecutor command.CommandExecutor
	commandBuilder command.CommandBuilder
}

func NewRecordFileFFMPEGRepository(commandBuilder command.CommandBuilder) RecordFileFFMPEGRepository {
	return RecordFileFFMPEGRepository{
		commandBuilder: commandBuilder,
	}
}

func (r RecordFileFFMPEGRepository) Create(done chan struct{}, record domain.Record, mediaPath string) <-chan error {

	filePath := fmt.Sprintf("%s/%s", mediaPath, VideoFileName)

	command := r.commandBuilder.Build("yt-dlp",
		//"-q",
		"--write-info-json",
		"--hls-use-mpegts",
		"--hls-prefer-ffmpeg",
		"-v",
		"-o", filePath, record.Url)

	errCh := make(chan error, 1)

	// Pass command to closure
	startCmd := func() error {
		if err := command.Run(); err != nil {
			return err
		}
		return nil
	}

	stopCmd := func() error {
		fmt.Printf("Stopping to record\n")

		if err := command.Signal(); err != nil {
			return err
		}

		return nil
	}

	go func() {
		defer close(done)
	out:
		for {
			select {
			case <-done:
				return
			default:
				if err := startCmd(); err != nil {
					errCh <- err
				}
				break out
			}
		}
	}()

	go func() {
		fmt.Println("Waiting for record...")
		<-done
		if err := stopCmd(); err != nil {
			errCh <- err
		}
	}()

	return errCh
}

func (r RecordFileFFMPEGRepository) Trim(id string, trim domain.Trim, mediaDir string) (trimmedPath *string, err error) {

	// Create folder to save the trimmed video
	trimmedVideoPath := fmt.Sprintf("%s/%s", mediaDir, trim.GetStartEndTimeFormatted())

	if _, err := os.Stat(trimmedVideoPath); os.IsNotExist(err) {
		if err := os.MkdirAll(trimmedVideoPath, 0755); err != nil {
			return nil, err
		}
	}

	var command command.CommandExecutor

	// Trim video
	command = r.commandBuilder.Build(
		"ffmpeg",
		"-i", fmt.Sprintf("%s/myvideo.mp4.part", mediaDir),
		"-ss", trim.StartTime,
		"-to", trim.EndTime,
		"-c", "copy", fmt.Sprintf("%s/%s", trimmedVideoPath, VideoFileName))

	if err := command.Run(); err != nil {
		return nil, err
	}

	// Create thumbnail
	command = r.commandBuilder.Build(
		"ffmpeg",
		"-i", fmt.Sprintf("%s/myvideo.mp4.part", mediaDir),
		"-ss", "00:00:01.000",
		"-vframes", "1",
		fmt.Sprintf("%s/thumbnail.jpg", trimmedVideoPath))

	if err := command.Run(); err != nil {
		return nil, err
	}

	return &trimmedVideoPath, nil
}

func (r RecordFileFFMPEGRepository) CreateHLS(inputFile string, segmentDuration int) error {

	pathSplit := strings.Split(inputFile, "/")
	mediaPath := strings.Join(pathSplit[0:len(pathSplit)-1], "/")

	createHLSCmd := func() error {
		// Create the HLS playlist and segment the video using ffmpeg
		command := r.commandBuilder.Build(
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

		if err := command.Run(); err != nil {
			return err
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

func (r RecordFileFFMPEGRepository) CreateDir(mediaDir string) error {
	return os.MkdirAll(mediaDir, 0755)
}
