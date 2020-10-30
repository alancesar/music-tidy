# tidy-music

Organize your music library according its tags.

## Building
```sh
cd cli
go build -o tidy-music
```

## Usage
```sh
./tidy-music [-s] [-o] [-t] [-p]
```

### Parameters
- **s**: The source path. Its default value is `./`.
- **o**: The output path. Its default value also is `./`.
- **t**: Test mode. If true only show the expected output. Its default is `false`.
- **p**: The output directory structure pattern. It follows the
Golang's [text/template](https://golang.org/pkg/text/template/) guide.
Available fields:
  - Artist
  - Album
  - Year
  - Track number
  - Title

## Output
The default `p` value is `{{.Artist}}/[{{.Year}}] {{.Album}}/{{printf "%02d" .Track}} - {{.Title}}` so the generated
output will be this following structure:
```
Artist
└── [Year] Album
    └── # - Title.ext
```
Example:
```
Green Day
├── [1994] Dookie
│   ├── 01 - Burnout.mp3
│   └── 02 - Having a Blast.mp3
└── [1995] Insomniac
    └── 01 - Armatage Shanks.wma
```
All the `/` and `:` in track and album names will be replaced by `-`.
