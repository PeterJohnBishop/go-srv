package watcher

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func watch() {
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
					return
				}

				if event.Has(fsnotify.Write) {
					log.Println("File modified:", event.Name)
				} else if event.Has(fsnotify.Remove) {
					log.Println("File removed:", event.Name)
				} else if event.Has(fsnotify.Rename) {
					log.Println("File renamed:", event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	// Add the file (or directory) to watch
	filename := "example.txt"
	err = watcher.Add(filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Watching %s for changes...\n", filename)

	<-done
}
