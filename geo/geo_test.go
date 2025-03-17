package geo_test

import (
	"demo-weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}

	got, err := geo.GetMyLocation(city)

	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получение %v", expected, got)
	}
}

func TestGetMyLocationNocity(t *testing.T) {
	city := "Londongdgdhh"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получение %v", geo.ErrNoCity, err)
	}
}
