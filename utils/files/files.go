package files

import (
	"os"
)

// GetAlbums returns a list of directories in the provided path.
// It is used to get the list of albums in the Google Photos export directory.
func GetAlbums(path string) ([]string, error) {
	var albums []string
	files, err := os.ReadDir(path)
	if err != nil {
		return albums, err
	}

	for _, file := range files {
		if file.IsDir() {
			albums = append(albums, file.Name())
		}
	}

	return albums, nil
}

// GetMedia returns a map of media files and their corresponding JSON files.
// It is used to get the list of media files and their corresponding JSON files in an album directory.
func GetMedia(albumPath string) (map[string]string, error) {
	result := make(map[string]string)

	entries, err := os.ReadDir(albumPath)
	if err != nil {
		return result, err
	}

	jsonFiles, mediaFiles := classifyFiles(entries)
	usedJsonFiles := make(map[string]struct{})

	for mediaFile := range mediaFiles {
		jsonFile, err := getJsonFile(mediaFile, jsonFiles)
		if err != nil {
			return nil, err
		}

		usedJsonFiles[jsonFile] = struct{}{}
		result[mediaFile] = jsonFile
	}

	err = checkUnusedJson(jsonFiles, usedJsonFiles)
	if err != nil {
		return nil, err
	}

	return result, nil
}
