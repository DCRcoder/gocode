package gocode

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func WatchAnyEvent(repo Repository) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	eventChan := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file: ", event.Name)
				}
				repo.Sync()
			case err := <-watcher.Errors:
				log.Println("error: ", err)
			}
		}
	}()

	err = watcher.Add(repo.path)
	if err != nil {
		log.Fatal(err)
	}
	<-eventChan
}
