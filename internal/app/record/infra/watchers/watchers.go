package watchers

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type RealTimeStreamWatcher struct {
	Watcher *fsnotify.Watcher
}

func NewRealTimeStreamWatcher() (*RealTimeStreamWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &RealTimeStreamWatcher{
		Watcher: watcher,
	}, nil
}

func (r *RealTimeStreamWatcher) Watch(path string) error {

	err := r.Watcher.Add(path)
	if err != nil {
		return err
	}

	return nil
}

func (r *RealTimeStreamWatcher) Unwatch(path string) error {
	return r.Watcher.Remove(path)
}

// Start listening for events.
func (r *RealTimeStreamWatcher) Listen() {
	go func() {
		for {
			select {
			case event, ok := <-r.Watcher.Events:
				if !ok {
					return
				}

				if event.Op.Has(fsnotify.Create) {
					log.Println("created file:", event.Name)
					if strings.Contains(event.Name, "myvideo.mp4.part") {
						log.Println("created file:", event.Name)
					}
				}
			case err, ok := <-r.Watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
}
