package adapters

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victoraldir/cutcast/internal/app/record/domain"
	command_mock "github.com/victoraldir/cutcast/pkg/command/mocks"
	"go.uber.org/mock/gomock"
)

var commandBuilderMock *command_mock.MockCommandBuilder
var commandExecutorMock *command_mock.MockCommandExecutor

func setup(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commandBuilderMock = command_mock.NewMockCommandBuilder(ctrl)
	commandExecutorMock = command_mock.NewMockCommandExecutor(ctrl)
}

func TestRecordFileFFMPEGRepository(t *testing.T) {

	t.Run("should create a record", func(t *testing.T) {

		setup(t)

		// Arrange
		wg := sync.WaitGroup{}

		wg.Add(1)

		commandBuilderMock.EXPECT().Build(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).Return(commandExecutorMock).Times(1)

		commandExecutorMock.EXPECT().Run().Return(nil).Do(func() {
			wg.Done()
		}).Times(1)

		done := make(chan struct{})
		recordCh := make(chan domain.Record, 1)
		mediaPath := "/tmp"

		recordCh <- domain.Record{
			Url: "https://www.youtube.com/watch?v=6g4dkBF5anU",
		}

		recordFileFFMPEGRepository := NewRecordFileFFMPEGRepository(commandBuilderMock)

		// Act
		errCh := recordFileFFMPEGRepository.Create(done, recordCh, mediaPath)

		wg.Wait()

		// Assert
		assert.Empty(t, errCh)
	})

	t.Run("should not create a record. Error when executing command", func(t *testing.T) {

		setup(t)

		// Arrange
		wg := sync.WaitGroup{}

		expectedError := errors.New("error when executing command")

		wg.Add(1)

		commandExecutorMock.EXPECT().Run().Return(expectedError).Do(func() {
			wg.Done()
		}).Times(1)

		commandBuilderMock.EXPECT().Build(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).Return(commandExecutorMock)

		recordFileFFMPEGRepository := NewRecordFileFFMPEGRepository(commandBuilderMock)

		done := make(chan struct{})
		recordCh := make(chan domain.Record, 1)
		mediaPath := "/tmp"

		recordCh <- domain.Record{
			Url: "https://www.youtube.com/watch?v=6g4dkBF5anU",
		}

		// Act
		errCh := recordFileFFMPEGRepository.Create(done, recordCh, mediaPath)

		wg.Wait()

		err := <-errCh

		// Assert
		assert.Equal(t, expectedError, err)

	})

	t.Run("should stop command", func(t *testing.T) {

		setup(t)

		// Arrange
		wg := sync.WaitGroup{}

		wg.Add(1)

		commandExecutorMock.EXPECT().Run().Return(nil).AnyTimes()

		// expectedError := errors.New("error when executing command")

		commandExecutorMock.EXPECT().Signal().Return(nil).Do(func() {
			wg.Done()
		}).AnyTimes()

		commandBuilderMock.EXPECT().Build(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).Return(commandExecutorMock)

		recordFileFFMPEGRepository := NewRecordFileFFMPEGRepository(commandBuilderMock)

		done := make(chan struct{})
		recordCh := make(chan domain.Record, 1)
		mediaPath := "/tmp"

		recordCh <- domain.Record{
			Url: "https://www.youtube.com/watch?v=6g4dkBF5anU",
		}

		// Act
		errCh := recordFileFFMPEGRepository.Create(done, recordCh, mediaPath)

		done <- struct{}{} // Stop command

		wg.Wait()

		// Assert
		assert.Empty(t, errCh)
	})

	t.Run("should error when stopping command", func(t *testing.T) {

		setup(t)

		// Arrange
		wg := sync.WaitGroup{}

		wg.Add(1)

		commandExecutorMock.EXPECT().Run().Return(nil).Do(func() { wg.Done() }).AnyTimes()

		expectedError := errors.New("error when executing command")

		commandBuilderMock.EXPECT().Build(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).Return(commandExecutorMock)

		commandExecutorMock.EXPECT().Signal().Return(expectedError).AnyTimes()

		recordFileFFMPEGRepository := NewRecordFileFFMPEGRepository(commandBuilderMock)

		done := make(chan struct{})
		recordCh := make(chan domain.Record, 1)
		mediaPath := "/tmp"

		recordCh <- domain.Record{
			Url: "https://www.youtube.com/watch?v=6g4dkBF5anU",
		}

		// Act
		errCh := recordFileFFMPEGRepository.Create(done, recordCh, mediaPath)

		wg.Wait()

		done <- struct{}{} // Stop command

		err := <-errCh

		// Assert
		assert.Equal(t, expectedError, err)
	})

	t.Run("should error when stopping command. No command", func(t *testing.T) {
		setup(t)

		// Arrange

		commandExecutorMock.EXPECT().Run().Return(nil).AnyTimes()

		commandBuilderMock.EXPECT().Build(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).Return(commandExecutorMock)

		recordFileFFMPEGRepository := NewRecordFileFFMPEGRepository(commandBuilderMock)

		done := make(chan struct{})
		recordCh := make(chan domain.Record, 1)
		mediaPath := "/tmp"

		// Act
		errCh := recordFileFFMPEGRepository.Create(done, recordCh, mediaPath)

		done <- struct{}{} // Stop command

		err := <-errCh

		// Assert
		assert.Equal(t, "command is nil", err.Error())
	})

}
