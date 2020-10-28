package main

import (
	"flag"
	"fmt"
	"github.com/alancesar/music-tidy/file"
	"github.com/alancesar/music-tidy/processor"
	"os"
	"path/filepath"
)

func main() {
	rootSourcePath := flag.String("s", "./", "source directory")
	rootDestinationPath := flag.String("o", "./", "output directory")
	flag.Parse()

	paths := make([]string, 0)
	_ = filepath.Walk(*rootSourcePath, func(path string, info os.FileInfo, err error) error {
		if err := file.CheckFile(path); err != nil {
			return nil
		}

		paths = append(paths, path)
		return nil
	})

	total := len(paths)

	for index, path := range paths {
		destination, bytes, err := processor.Process(path, *rootDestinationPath)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s [%d] (%d/%d)\n", destination, bytes, index, total)
	}
}
