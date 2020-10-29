package metadata

import (
	"github.com/alancesar/tidy-music/sanitize"
	"github.com/dhowden/tag"
	"io"
)

type Metadata struct {
	Track  int
	Title  string
	Artist string
	Album  string
	Year   int
}

func ExtractMetadata(r io.ReadSeeker) (Metadata, error) {
	metadata, err := tag.ReadFrom(r)
	if err != nil {
		return Metadata{}, err
	}

	track, _ := metadata.Track()
	artist := metadata.AlbumArtist()
	if artist == "" {
		artist = metadata.Artist()
	}

	return Metadata{
		Track:  track,
		Title:  sanitize.Sanitize(metadata.Title()),
		Artist: sanitize.Sanitize(artist),
		Album:  sanitize.Sanitize(metadata.Album()),
		Year:   metadata.Year(),
	}, nil
}
