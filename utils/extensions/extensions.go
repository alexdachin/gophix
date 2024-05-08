package extensions

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Fix(mediaPath string) (string, error) {
	currentExt := filepath.Ext(mediaPath)
	newExt, err := getNewExtension(mediaPath)
	if err != nil {
		return mediaPath, fmt.Errorf("could not get the proper extensions for %s: %w", mediaPath, err)
	}

	if areExtensionsCompatible(currentExt, newExt) {
		return mediaPath, nil
	}

	baseFileName := strings.TrimSuffix(mediaPath, currentExt)
	newMediaPath, err := getNewFileName(baseFileName, newExt)
	if err != nil {
		return mediaPath, fmt.Errorf("could not generate a new file name for %s with %s extensions: %w", mediaPath, newExt, err)
	}

	err = os.Rename(mediaPath, newMediaPath)
	if err != nil {
		return mediaPath, err
	}

	fmt.Printf("🔄 renamed %s -> %s\n", mediaPath, newMediaPath)
	return newMediaPath, nil
}

func getNewExtension(mediaPath string) (string, error) {
	cmd := exec.Command("exiftool", "-p", ".$FileTypeExtension", mediaPath)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func areExtensionsCompatible(ext1 string, ext2 string) bool {
	if strings.EqualFold(ext1, ext2) {
		return true
	}

	compatibleExtensions := map[string]string{
		".jpg": ".jpeg",
		".tif": ".tiff",
		".mov": ".mp4",
	}

	for k, v := range compatibleExtensions {
		if (strings.EqualFold(ext1, k) && strings.EqualFold(ext2, v)) ||
			(strings.EqualFold(ext1, v) && strings.EqualFold(ext2, k)) {
			return true
		}
	}

	return false
}

func getNewFileName(baseFileName string, newExtension string) (string, error) {
	if !doesFileExist(baseFileName + newExtension) {
		return baseFileName + newExtension, nil
	}

	const maxAttempts = 10
	for i := 0; i < maxAttempts; i++ {
		newFileName := baseFileName + "-" + generateRandomSuffix() + newExtension
		if !doesFileExist(newFileName) {
			return newFileName, nil
		}
	}

	return "", fmt.Errorf("could not generate a unique file name after %d attempts", maxAttempts)
}

func generateRandomSuffix() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func doesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
