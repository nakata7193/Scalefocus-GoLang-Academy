package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	ApiKey = "3d4e293aa6fe4c9fc0c4318f73ea299d"
)

type Description struct {
	Description string `json:"description"`
}

type Temperature struct {
	Temp float64 `json:"temp"`
}

type WeatherDataPayload struct {
	Weather []Description `json:"weather"`
	City    string        `json:"name"`
	Temp    Temperature   `json:"main"`
}

type WeatherService struct {
	Lat string
	Lon string
}

func GetWeather(Lat string, Lon string) (map[string]string, error) {
	url := "http://api.openweathermap.org/data/2.5/weather?lat=" + Lat + "&lon=" + Lon + "&APPID=" + ApiKey
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	defer resp.Body.Close()

	payload := WeatherDataPayload{}
	err = json.NewDecoder(resp.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	payloadMap := map[string]string{}
	payloadMap["city"] = payload.City
	payloadMap["description"] = payload.Weather[0].Description
	payloadMap["temp"] = strconv.FormatFloat(payload.Temp.Temp, 'f', 2, 64)
	return payloadMap, nil
}
