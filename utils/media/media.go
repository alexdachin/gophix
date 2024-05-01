package media

import (
	"os"
)

func Get(albumPath string) (map[string]string, error) {
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
