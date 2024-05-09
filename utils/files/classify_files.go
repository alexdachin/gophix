package files

import (
	"os"
	"strings"
)

func classifyFiles(entries []os.DirEntry) (map[string]struct{}, map[string]struct{}) {
	jsonFiles := make(map[string]struct{})
	mediaFiles := make(map[string]struct{})

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if strings.HasSuffix(entry.Name(), ".json") {
			jsonFiles[entry.Name()] = struct{}{}
		} else {
			mediaFiles[entry.Name()] = struct{}{}
		}
	}

	return jsonFiles, mediaFiles
}
