package domain

type Coordinates struct {
	Lat  float64 `json:"long"`
	Long float64 `json:"lat"`
}

type Address struct {
	//
	// Struct for holding address information
	//

	// Static
	PostCode string
	Coordinates

	// Dynamic
}
