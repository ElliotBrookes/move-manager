package domain

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
