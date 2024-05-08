package main

import (
	"fmt"
	"github.com/alexdachin/google-photos-fix/utils/albums"
	"github.com/alexdachin/google-photos-fix/utils/extensions"
	"github.com/alexdachin/google-photos-fix/utils/media"
	"github.com/alexdachin/google-photos-fix/utils/metadata"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to the Google Photos export directory as an argument.")
	}

	path := os.Args[1]
	albumEntries, err := albums.Get(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range albumEntries {
		albumPath := filepath.Join(path, album)
		fmt.Println("📓 processing album:", album)
		processAlbum(albumPath)
	}
}

func processAlbum(albumPath string) {
	files, err := media.Get(albumPath)
	if err != nil {
		log.Fatal(err)
	}

	for mediaFile, jsonFile := range files {
		mediaPath := filepath.Join(albumPath, mediaFile)
		jsonPath := filepath.Join(albumPath, jsonFile)

		newMediaPath, err := extensions.Fix(mediaPath)
		if err != nil {
			fmt.Printf("🚨 extension for file %s could not be fixed: %s\n", mediaPath, err)
			continue
		}

		err = metadata.Apply(newMediaPath, jsonPath)
		if err != nil {
			fmt.Printf("🚨 metadata for file %s could not be applied: %s\n", newMediaPath, err)
			continue
		}
	}

	for _, jsonFile := range files {
		jsonPath := filepath.Join(albumPath, jsonFile)
		err := os.Remove(jsonPath)
		if err != nil {
			if os.IsNotExist(err) {
				// file was already removed
				continue
			}

			fmt.Printf("🚨 json file %s could not be removed: %s\n", jsonPath, err)
		}
	}
}
