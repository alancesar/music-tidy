package main

import (
	"flag"
	"fmt"
	"github.com/alancesar/tidy-music/command"
	"github.com/alancesar/tidy-music/file"
	"github.com/alancesar/tidy-music/processor"
	"os"
	"path/filepath"
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
	paths := make([]string, 0)
	_ = filepath.Walk(*rootSourcePath, func(path string, info os.FileInfo, err error) error {
		if !file.IsFile(path) {
			return nil
		}

		paths = append(paths, path)
		return nil
	})

	total := len(paths)
	var commands []command.Command

	if !*sandbox {
		commands = []command.Command{command.MkDirCommand, os.Rename}
	}

	for index, path := range paths {
		destination, err := processor.Process(path, *rootDestinationPath, *pattern, commands...)
		if err != nil && err != processor.MetadataErr {
			panic(err)
		}

		if err == nil {
			fmt.Printf("(%d/%d) %s\n", index+1, total, destination)
		}
	}
}
