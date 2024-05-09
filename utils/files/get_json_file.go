package files

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var numberSuffixRe = regexp.MustCompile(`\(\d+\)`)

func getJsonFile(mediaFile string, jsonFiles map[string]struct{}) (string, error) {
	if _, ok := jsonFiles[mediaFile+".json"]; ok {
		return mediaFile + ".json", nil
	}

	if strings.Contains(mediaFile, "-edited") {
		return getJsonFile(strings.Replace(mediaFile, "-edited", "", 1), jsonFiles)
	}

	if numberSuffixRe.MatchString(mediaFile) {
		match := numberSuffixRe.FindString(mediaFile)
		jsonFile := strings.Replace(mediaFile, match, "", 1) + match + ".json"
		if _, ok := jsonFiles[jsonFile]; ok {
			return jsonFile, nil
		}
	}

	if len(mediaFile) > 46 && numberSuffixRe.MatchString(mediaFile) {
		match := numberSuffixRe.FindString(mediaFile)
		jsonFile := mediaFile[:46] + match + ".json"
		if _, ok := jsonFiles[jsonFile]; ok {
			return jsonFile, nil
		}
	}

	if len(mediaFile) > 46 {
		jsonFile := mediaFile[:46] + ".json"
		if _, ok := jsonFiles[jsonFile]; ok {
			return jsonFile, nil
		}
	}

	mediaExt := filepath.Ext(mediaFile)
	if strings.EqualFold(mediaExt, ".mp4") {
		baseMediaFile := strings.TrimSuffix(mediaFile, mediaExt)
		extensions := [...]string{".jpg", ".jpeg", ".heic"}
		for _, ext := range extensions {
			jsonFile := baseMediaFile + ext + ".json"
			jsonFileUpper := baseMediaFile + strings.ToUpper(ext) + ".json"
			if _, ok := jsonFiles[jsonFile]; ok {
				return jsonFile, nil
			}
			if _, ok := jsonFiles[jsonFileUpper]; ok {
				return jsonFileUpper, nil
			}
		}
	}

	return "", fmt.Errorf("json file not found for %s", mediaFile)
}
