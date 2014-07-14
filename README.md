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

## Example

This is example to automatically apply [dotremap](https://github.com/k0kubun/dotremap) configuration.

```bash
if ! pgrep -f autobuild_dotremap > /dev/null; then
  echo "Launched dotremap autobuild daemon"
  nohup autobuild "dotremap" -f ~/.remap -i autobuild_dotremap > /dev/null &
fi
```

You can use `-i` option for just assigning identifier to process.
