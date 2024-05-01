package albums

import "os"

func Get(path string) ([]string, error) {
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
