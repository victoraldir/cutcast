package adapters

import (
	"github.com/victoraldir/cutcast/internal/app/record/infra/watchers"
)

type WatcherDbMemoryRepository struct {
	Watchers        map[string]watchers.RealTimeStreamWatcher
	WatchCallback   func(path string) error
	UnwatchCallback func(path string) error
}

func NewWatcherDbMemoryRepository(watchCallback func(string) error, unwatchCallback func(path string) error) WatcherDbMemoryRepository {
	return WatcherDbMemoryRepository{
		Watchers:        make(map[string]watchers.RealTimeStreamWatcher, 100),
		WatchCallback:   watchCallback,
		UnwatchCallback: unwatchCallback,
	}
}

func (r WatcherDbMemoryRepository) Watch(path string) error {

	// realTimeStreamWatcher, err := watchers.NewRealTimeStreamWatcher(path, createHLSRecordGroupUseCase)

	// if err != nil {
	// 	return err
	// }

	// r.Watchers[path] = *realTimeStreamWatcher

	err := r.WatchCallback(path)
	if err != nil {
		return err
	}

	return nil
}

func (r WatcherDbMemoryRepository) Unwatch(path string) error {
	return r.Watchers[path].Watcher.Close()
}
