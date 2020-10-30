package main

import (
	"flag"
	"fmt"
	"github.com/alancesar/tidy-music/command"
	"github.com/alancesar/tidy-music/mime"
	"github.com/alancesar/tidy-music/path"
	"github.com/alancesar/tidy-music/processor"
	"os"
)

const (
	defaultPattern = "{{.Artist}}/[{{.Year}}] {{.Album}}/{{printf \"%02d\" .Track}} - {{.Title}}"
)

func main() {
	rootSourcePath := flag.String("s", "./", "source directory")
	rootDestinationPath := flag.String("o", "./", "output directory")
	pattern := flag.String("p", defaultPattern, "output pattern")
	sandbox := flag.Bool("t", false, "run in test mode")
	flag.Parse()

	fmt.Println("Reading source directory...")
	paths := path.LookFor(*rootSourcePath, mime.AudioType)
	total := len(paths)

	var commands []command.Command

	if !*sandbox {
		commands = []command.Command{command.MkDirCommand, os.Rename}
	}

	for index, p := range paths {
		destination, err := processor.Process(p, *rootDestinationPath, *pattern, commands...)
		if err != nil {
			fmt.Printf("(%d/%d) [failed ] %s\n", index+1, total, destination)
		} else {
			fmt.Printf("(%d/%d) [success] %s\n", index+1, total, destination)
		}
	}
}
