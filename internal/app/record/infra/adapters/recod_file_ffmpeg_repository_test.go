package adapters

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

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

	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())

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

		commandExecutorMock.EXPECT().Run().Return(nil).Times(1)

		commandExecutorMock.EXPECT().Signal().Return(nil).Do(func() {
			wg.Done()
		}).AnyTimes()

		done := make(chan struct{})

		mediaPath := "/tmp"

		recordFileFFMPEGRepository := NewRecordFileFFMPEGRepository(commandBuilderMock)

		// Act
		err := recordFileFFMPEGRepository.Create(done, domain.Record{
			Url: "https://www.youtube.com/watch?v=6g4dkBF5anU",
		}, mediaPath)

		wg.Wait()

		time.Sleep(1 * time.Second)

		// Assert
		assert.Empty(t, err)
		assert.Equal(t, 3, runtime.NumGoroutine())
	})

	t.Run("should not create a record. Error when executing command", func(t *testing.T) {

		setup(t)

		// Arrange
		wg := sync.WaitGroup{}

		expectedError := errors.New("error when executing command")

		wg.Add(1)

		commandExecutorMock.EXPECT().Run().Return(expectedError).Times(1)

		commandBuilderMock.EXPECT().Build(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).Return(commandExecutorMock)

		commandExecutorMock.EXPECT().Signal().Return(nil).Do(func() {
			wg.Done()
		}).AnyTimes()

		recordFileFFMPEGRepository := NewRecordFileFFMPEGRepository(commandBuilderMock)

		done := make(chan struct{})
		mediaPath := "/tmp"

		// Act
		errCh := recordFileFFMPEGRepository.Create(done, domain.Record{
			Url: "https://www.youtube.com/watch?v=6g4dkBF5anU",
		}, mediaPath)

		wg.Wait()

		err := <-errCh

		// Assert
		assert.Equal(t, expectedError, err)
		assert.Equal(t, 3, runtime.NumGoroutine())
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

		mediaPath := "/tmp"

		// Act
		errCh := recordFileFFMPEGRepository.Create(done, domain.Record{
			Url: "https://www.youtube.com/watch?v=6g4dkBF5anU",
		}, mediaPath)

		done <- struct{}{} // Stop command

		time.Sleep(5 * time.Second)

		fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())

		wg.Wait()

		fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())

		// Assert
		assert.Empty(t, errCh)
	})

	t.Run("should error when stopping command", func(t *testing.T) {

		setup(t)

		// Arrange
		wg := sync.WaitGroup{}

		wg.Add(1)

		commandExecutorMock.EXPECT().Run().DoAndReturn(func() error {
			time.Sleep(10 * time.Second)
			return nil
		}).AnyTimes()

		expectedError := errors.New("error when executing command")

		commandBuilderMock.EXPECT().Build(gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any()).Return(commandExecutorMock)

		commandExecutorMock.EXPECT().Signal().Return(expectedError).Do(func() {
			wg.Done()
		}).AnyTimes()

		recordFileFFMPEGRepository := NewRecordFileFFMPEGRepository(commandBuilderMock)

		done := make(chan struct{})
		recordCh := make(chan domain.Record, 1)
		mediaPath := "/tmp"

		recordCh <- domain.Record{
			Url: "https://www.youtube.com/watch?v=6g4dkBF5anU",
		}

		// Act
		errCh := recordFileFFMPEGRepository.Create(done, domain.Record{
			Url: "https://www.youtube.com/watch?v=6g4dkBF5anU",
		}, mediaPath)

		done <- struct{}{} // Stop command

		wg.Wait()

		err := <-errCh

		// time.Sleep(2 * time.Second)

		// Assert
		assert.Equal(t, expectedError, err)
		// assert.Equal(t, 3, runtime.NumGoroutine())
	})

}
