# Autobuild

Simple automatic build tool, which detects file changes by fsnotify.

## Installation

```bash
$ go get github.com/k0kubun/autobuild
```

## Usage

```bash
$ nohup autobuild "go build ~/something.go -o something" -f ~/something.go &
```

Then autobuild daemon executes `go build ~/something.go` whenever ~/something.go changes.
