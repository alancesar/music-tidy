package main

import (
	"github.com/alancesar/tidy-music/command"
	"github.com/alancesar/tidy-music/metadata"
	"github.com/alancesar/tidy-music/path"
	"path/filepath"
)

func Process(sourcePath, rootDestinationPath, pattern string, commands ...command.Command) (string, error) {
	m, err := metadata.NewExtractor(sourcePath).Extract()
	if err != nil {
		return "", err
	}

	destinationPath, err := path.BuildFromPattern(pattern, m)
	if err != nil {
		return "", err
	}

	destinationPath = destinationPath + filepath.Ext(sourcePath)
	destinationPath = filepath.Join(rootDestinationPath, destinationPath)
	destinationPath = filepath.Clean(destinationPath)

	if err := command.NewExecutor(sourcePath, destinationPath).Execute(commands...); err != nil {
		return destinationPath, err
	}
	return destinationPath, nil
}
