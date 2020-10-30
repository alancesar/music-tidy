package path

import (
	"github.com/alancesar/tidy-music/metadata"
	"os"
	"strings"
	"testing"
)

func TestBuildFromPattern(t *testing.T) {
	m := metadata.Metadata{
		Track:  7,
		Title:  "Basket Case",
		Artist: "Green Day",
		Album:  "Dookie",
		Year:   1994,
	}

	path, err := BuildFromPattern("{{.Artist}}/[{{.Year}}] {{.Album}}/{{printf \"%02d\" .Track}} - {{.Title}}", m)
	if err != nil {
		t.Error(err)
	}

	expected := strings.ReplaceAll("Green Day/[1994] Dookie/07 - Basket Case", "/", string(os.PathSeparator))
	if expected != path {
		t.Error("unexpected path")
	}
}
