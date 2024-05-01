package media

import (
	"fmt"
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

	if strings.HasSuffix(mediaFile, ".MP4") {
		extensions := [...]string{".jpg", ".jpeg", ".heic", ".JPG", ".JPEG", ".HEIC"}
		for _, ext := range extensions {
			jsonFile := mediaFile[:len(mediaFile)-4] + ext + ".json"
			if _, ok := jsonFiles[jsonFile]; ok {
				return jsonFile, nil
			}
		}
	}

	return "", fmt.Errorf("json file not found for %s", mediaFile)
}
