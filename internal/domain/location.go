package domain

import (
	"strings"
	"unicode"
)

type Coordinates struct {
	Lat  float64 `json:"long"`
	Long float64 `json:"lat"`
}

type Location struct {
	//
	// Struct for holding address information
	//

	// Static
	postCode string

	// Dynamic
	coords *Coordinates
}

func (l Location) Latitude() float64 {
	return l.coords.Lat
}

func (l Location) Longitude() float64 {
	return l.coords.Long
}

func (l Location) PostCode() string {
	return l.postCode
}

func NewLocation(cds Coordinates, pc string) (Location, error) {
	// Uppers and strips spacing around postcode
	upc := strings.ToUpper(pc)
	upc = spaceStringsBuilder(upc)

	return Location{
		coords:   &cds,
		postCode: upc,
	}, nil
}

// Small function to quickly remove space chars from postcodes
func spaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) && (unicode.IsNumber(ch) || unicode.IsLetter(ch)) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
