package domain

type Vehicle struct {
	//
	// Fields holding vehicle information
	//

	// Static size fields
	reg                  string
	make                 string
	fuelType             string
	realDrivingEmissions string
	engCap               int32
	emissionsgkm         int32
	mWeight              int32
	yom                  int32

	// Dynamic size fields
}
