package file

import (
	"fmt"
	"github.com/alancesar/tidy-music/metadata"
	"github.com/alancesar/tidy-music/sanitize"
	"os"
	"strings"
)

func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	if stat.Mode().IsDir() {
		return false
	}

	if !stat.Mode().IsRegular() {
		return false
	}

	return true
}

func BuildFilename(metadata metadata.Metadata, extension string) string {
	title := sanitize.Sanitize(metadata.Title)
	filename := fmt.Sprintf("%02d - %s%s", metadata.Track, title, extension)
	sanitized := strings.TrimSpace(filename)
	return sanitized
}
