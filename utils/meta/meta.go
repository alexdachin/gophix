package meta

import "encoding/json"

type GeoData struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Altitude      float64 `json:"altitude"`
	LatitudeSpan  float64 `json:"latitudeSpan"`
	LongitudeSpan float64 `json:"longitudeSpan"`
}

type Time struct {
	Timestamp string `json:"timestamp"`
}

type Data struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	CreationTime   Time    `json:"creationTime"`
	PhotoTakenTime Time    `json:"photoTakenTime"`
	GeoData        GeoData `json:"geoData"`
}

func Parse(fileContent []byte) (Data, error) {
	var data Data
	err := json.Unmarshal(fileContent, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
