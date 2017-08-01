package main

// Code from https://github.com/howeyc/fsnotify

import (
	"log"
	"github.com/howeyc/fsnotify"
)

func main() {

	WATCHED_DIR := "mounted"

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch(WATCHED_DIR)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting watching in: " + WATCHED_DIR);

	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()
}
