package retrieval

import (
	"errors"
	"net/http"
	"os"
)

func getMapBoxRoute() (http.Response, error) {
	// get the api key from env
	key, ok := os.LookupEnv("MAPBOX_API_TOKEN")
	if !ok {
		return http.Response{}, errors.New("No MapBox API token set")
	}

	// Using passed data and api key, make api request
	http.Get()

	// Handle and return response data

	return

}
