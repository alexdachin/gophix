package main

import (
	"github.com/alexdachin/google-photos-import/utils/albums"
	"github.com/alexdachin/google-photos-import/utils/media"
	"github.com/schollz/progressbar/v3"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	path := "/Users/alexdachin/Downloads/Takeout/Google Photos/"
	newPath := "/Users/alexdachin/Downloads/Takeout/Google Photos Fixed/"
	albumEntries, err := albums.Get(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range albumEntries {
		albumPath := filepath.Join(path, album)
		newAlbumPath := filepath.Join(newPath, album)
		err := os.MkdirAll(newAlbumPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("📓 Processing album:", album)
		processAlbum(albumPath, newAlbumPath)
	}
}

func processAlbum(albumPath string, newAlbumPath string) {
	files, err := media.Get(albumPath)
	if err != nil {
		log.Fatal(err)
	}

	bar := progressbar.Default(int64(len(files)))
	for mediaFile, jsonFile := range files {
		mediaPath := filepath.Join(albumPath, mediaFile)
		jsonPath := filepath.Join(albumPath, jsonFile)
		newMediaPath := filepath.Join(newAlbumPath, mediaFile)

		cmd := exec.Command(
			"exiftool",
			"-d",
			"%s",
			"-TagsFromFile",
			jsonPath,
			"-Title<Title",
			"-Description<Description",
			"-ImageDescription<Description",
			"-Caption-Abstract<Description",
			"-AllDates<PhotoTakenTimeTimestamp",
			"-GPSAltitude<GeoDataAltitude",
			"-GPSLatitude<GeoDataLatitude",
			"-GPSLatitudeRef<GeoDataLatitude",
			"-GPSLongitude<GeoDataLongitude",
			"-GPSLongitudeRef<GeoDataLongitude",
			"-o",
			newMediaPath,
			mediaPath,
		)

		data, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("🚨 Error processing file:", mediaFile)
			log.Println("➡️", cmd.String())
			log.Println("ℹ️", string(data))
		}

		_ = bar.Add(1)
	}
}
