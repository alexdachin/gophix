package main

import (
	"fmt"
	"github.com/alexdachin/gophix/utils/extensions"
	"github.com/alexdachin/gophix/utils/files"
	"github.com/alexdachin/gophix/utils/metadata"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	OperationFix       = "fix"
	OperationCleanJson = "clean-json"
)

func init() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

func main() {
	operation, path := getArgs()
	albumEntries, err := files.GetAlbums(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range albumEntries {
		albumPath := filepath.Join(path, album)
		fmt.Println("📓 processing album:", album)

		switch operation {
		case OperationFix:
			fixAlbum(albumPath)
		case OperationCleanJson:
			cleanAlbumJson(albumPath)
		}
	}
}

func getArgs() (string, string) {
	help := `Usage: gophix <operation> <path>

Supported operations:
$ gophix fix /path/to/takeout
$ gophix clean-json /path/to/takeout

`

	if len(os.Args) < 3 {
		log.Fatal(help)
	}

	if os.Args[1] != OperationFix && os.Args[1] != OperationCleanJson {
		log.Fatalf("Operation %s is not supported.\n\n%s", os.Args[1], help)
	}

	return os.Args[1], os.Args[2]
}

func fixAlbum(albumPath string) {
	filesMap, err := files.GetMedia(albumPath)
	if err != nil {
		log.Fatal(err)
	}

	for mediaFile, jsonFile := range filesMap {
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
}

func cleanAlbumJson(albumPath string) {
	entries, err := os.ReadDir(albumPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if strings.EqualFold(filepath.Ext(entry.Name()), ".json") {
			err := os.Remove(filepath.Join(albumPath, entry.Name()))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
