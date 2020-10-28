package file

import (
	"fmt"
	"github.com/alancesar/music-tidy/metadata"
	"github.com/alancesar/music-tidy/sanitize"
	"os"
	"strings"
)

func Check(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", path)
	}

	return nil
}

func BuildFilename(metadata metadata.Metadata, extension string) string {
	title := sanitize.Sanitize(metadata.Title)
	filename := fmt.Sprintf("%02d - %s%s", metadata.Track, title, extension)
	sanitized := strings.TrimSpace(filename)
	return sanitized
}
