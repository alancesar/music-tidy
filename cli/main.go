package main

import (
	"flag"
	"fmt"
	"github.com/alancesar/tidy-file/command"
	"github.com/alancesar/tidy-file/mime"
	"github.com/alancesar/tidy-file/path"
	"github.com/alancesar/tidy-music/metadata"
	"os"
	"path/filepath"
)

const (
	defaultPattern = "{{.Artist}}/[{{.Year}}] {{.Album}}/{{printf \"%02d\" .Track}} - {{.Title}}"
)

func main() {
	rootSourcePath := flag.String("s", "", "source directory")
	rootDestinationPath := flag.String("o", "", "output directory")
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

	for index, sourcePath := range paths {
		destination, err := process(sourcePath, *rootDestinationPath, *pattern, commands...)
		if err != nil {
			fmt.Printf("(%d/%d) [failed ] %s\n", index+1, total, destination)
		} else {
			fmt.Printf("(%d/%d) [success] %s\n", index+1, total, destination)
		}
	}
}

func process(sourcePath, rootDestinationPath, pattern string, commands ...command.Command) (string, error) {
	m, err := metadata.NewExtractor(sourcePath).Extract()
	if err != nil {
		return "", err
	}

	parsed, err := path.NewBuilder().FromPattern(pattern, m)
	if err != nil {
		return "", err
	}

	destinationPath := parsed + filepath.Ext(sourcePath)
	destinationPath = filepath.Join(rootDestinationPath, destinationPath)
	destinationPath = filepath.Clean(destinationPath)

	err = command.NewExecutor(sourcePath, destinationPath).Execute(commands...)
	return destinationPath, nil
}
