package metadata

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func Apply(mediaPath string, jsonPath string) error {
	var outMediaPath string
	if hasSupportedExtension(mediaPath) {
		outMediaPath = mediaPath
	} else {
		fmt.Printf("🚗 writing metadata in xmp sidecar for %s\n", mediaPath)
		outMediaPath = mediaPath + ".xmp"
	}

	cmd := exec.Command(
		"exiftool",
		"-d", "%s",
		"-m",
		"-TagsFromFile", jsonPath,
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
		"-overwrite_original",
		outMediaPath,
	)

	data, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("could not fix metadata for %s\nerror: %w\noutput: %s", mediaPath, err, string(data))
	}

	return nil
}

func hasSupportedExtension(path string) bool {
	ext := filepath.Ext(path)

	var exifSupportExtensions = []string{
		".3gp",
		".dng",
		".gif",
		".heic",
		".jpeg",
		".jpg",
		".m4v",
		".mov",
		".mp4",
		".png",
		".tif",
	}

	for _, supportedExt := range exifSupportExtensions {
		if strings.EqualFold(ext, supportedExt) {
			return true
		}
	}

	return false
}
