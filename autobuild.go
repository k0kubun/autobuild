package main

import (
	"log"
	"os/exec"

	"github.com/howeyc/fsnotify"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Files      []string `short:"f" long:"files" description:"files whose changes are detected"`
	Identifier string   `short:"i" long:"identifier" description:"You can pass this option to distinguish processes but it does nothing"`
}

func autobuild(command string, paths []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	for _, path := range paths {
		err = watcher.Watch(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	for {
		select {
		case event := <-watcher.Event:
			if event.IsModify() {
				exec.Command(command).Run()
			}
		}
	}
}

func main() {
	options := new(Options)
	args, err := flags.Parse(options)
	if err != nil {
		return
	}

	if len(options.Files) < 1 {
		println("Given no files to watch. You have to select them by -f option.")
		return
	}

	if len(args) < 1 {
		println("No command is given. You have to specify build command like:\n  $ autobuild 'go build foo.go' -f foo.go")
		return
	}

	autobuild(args[0], options.Files)
}
