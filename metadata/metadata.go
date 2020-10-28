package metadata

import (
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

	return Metadata{
		Track:  track,
		Title:  metadata.Title(),
		Artist: metadata.AlbumArtist(),
		Album:  metadata.Album(),
		Year:   metadata.Year(),
	}, nil
}
