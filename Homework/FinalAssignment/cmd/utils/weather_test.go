package utils

import "testing"

func TestGetWeather(t *testing.T) {
	latitude, longitude := "0", "0"
	weather, err := GetWeather(latitude, longitude)
	if err != nil {
		t.Error(err)
	}

	if weather["city"] == "" || weather["description"] == "" || weather["temp"] == "" {
		t.Errorf("No weather data returned")
	}
	t.Log(weather)
}
