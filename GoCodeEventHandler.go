package gocode

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
)

var FilterName = []string{".git", "eggs", "develop-eggs", "gen-go", "gen-py", ".idea"}

func filterDir(fileName string) bool {
	for _, name := range FilterName {
		if name == fileName {
			return true
		}
	}
	return false
}

func findAllDir(path string) []string {
	dirs := make([]string, 0, 0)
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if filterDir(f.Name()) {
			continue
		}
		subDirName := path + "/" + f.Name()
		if f.IsDir() {
			dirs = append(dirs, subDirName)
			subDirs := findAllDir(subDirName)
			dirs = append(dirs, subDirs...)
		}
	}
	return dirs
}

func WatchAnyEvent(repo Repository) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					fmt.Println("cache error event")
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
				repo.mut.Lock()
				go repo.Sync()
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	dirs := findAllDir(repo.path)
	err = watcher.Add(repo.path)
	for _, dir := range dirs {
		watcher.Add(dir)
	}

	if err != nil {
		log.Fatal(err)
	}
	<-done
}
