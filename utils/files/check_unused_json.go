package files

import "fmt"

func checkUnusedJson(jsonFiles map[string]struct{}, usedJsonFiles map[string]struct{}) error {
	for jsonFile := range jsonFiles {
		if _, ok := usedJsonFiles[jsonFile]; !ok {
			return fmt.Errorf("json file %s not used", jsonFile)
		}
	}

	return nil
}
