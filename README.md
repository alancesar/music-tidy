# tidy-music

Organize your music library according its tags.

## Usage
```sh
go build
./tidy-music -s /origin/path -o /destination/path
```
This command will move all music files to the specified output directory, following this structure:
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
