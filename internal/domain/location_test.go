package domain

import (
	"testing"
)

// Test for Location Entity in domain layer
func TestLocation(t *testing.T) {
	// Setup for test location

	// Test for the NewLocation Function
	t.Run("Test Creation with test data", func(t *testing.T) {
		var tests = []struct {
			name           string
			coords         Coordinates
			postcode       string
			wantedCoords   Coordinates
			wantedPostCode string
		}{
			{"Lower cased postcode", Coordinates{5.0, 5.0}, "ab12cd", Coordinates{5.0, 5.0}, "AB12CD"},
			{"Postcode with spaces", Coordinates{5.0, 5.0}, "ab1 2cd", Coordinates{5.0, 5.0}, "AB12CD"},
			{"Postcode with symbols", Coordinates{5.0, 5.0}, "a%12cd", Coordinates{5.0, 5.0}, "A12CD"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				// Create Location from func
				createdLocation, err := NewLocation(test.coords, test.postcode)
				if err != nil {
					t.Errorf("%e", err)
				}

				// assert values are the wanted values
				if createdLocation.postCode != test.wantedPostCode {
					t.Errorf("error initialising postcode: got %s, want %s", createdLocation.postCode, test.wantedPostCode)
				}

				if *createdLocation.coords != test.wantedCoords {
					t.Errorf("error initialising Coords: got %g, want %g", *createdLocation.coords, test.wantedCoords)
				}

			})
		}

	})

	// Test ensuring the getter functions are working as intended
	t.Run("Test Getter functions with test function", func(t *testing.T) {
		// Create location struct
		testLocation, err := manualLocationCreation()
		if err != nil {
			t.Errorf("%e", err)
		}

		// obtain values from getter functions
		lat := testLocation.Latitude()
		long := testLocation.Longitude()
		pc := testLocation.PostCode()

		// Validate no mutations occured
		if lat != testLocation.coords.Lat ||
			long != testLocation.coords.Long ||
			pc != testLocation.postCode {
			t.Errorf("Getter functions should not mutate field values")
		}

	})

}

// Creates a Locataion struct manually without the use of the NewLocation function in the pkg
func manualLocationCreation() (Location, error) {
	// Create and return a location struct for testing purposes
	return Location{
		coords: &Coordinates{
			Lat:  10.0,
			Long: 10.0,
		},
		postCode: "AA11BB",
	}, nil
}
