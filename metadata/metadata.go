package metadata

import (
	"github.com/alancesar/tidy-music/sanitize"
	"github.com/dhowden/tag"
	"os"
)

type Metadata struct {
	Track  int
	Title  string
	Artist string
	Album  string
	Year   int
}

type Extractor struct {
	path string
}

func NewExtractor(path string) *Extractor {
	return &Extractor{
		path: path,
	}
}

func (e *Extractor) Extract() (Metadata, error) {
	source, err := os.Open(e.path)
	if err != nil {
		return Metadata{}, err
	}

	defer func() {
		_ = source.Close()
	}()

	metadata, err := tag.ReadFrom(source)
	if err != nil {
		return Metadata{}, err
	}

	return buildOutput(metadata), nil
}

func extractTrackAndArtist(metadata tag.Metadata) (int, string) {
	track, _ := metadata.Track()
	artist := metadata.AlbumArtist()
	if artist == "" {
		artist = metadata.Artist()
	}

	return track, artist
}

func buildOutput(metadata tag.Metadata) Metadata {
	track, artist := extractTrackAndArtist(metadata)

	return Metadata{
		Track:  track,
		Title:  sanitize.Sanitize(metadata.Title()),
		Artist: sanitize.Sanitize(artist),
		Album:  sanitize.Sanitize(metadata.Album()),
		Year:   metadata.Year(),
	}
}
