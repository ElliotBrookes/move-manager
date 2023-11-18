package domain

type Journey struct {
	//
	// Information regarding a portion of driving in a move
	//

	// Static size data
	From string
	To   string

	// Dynamically Sized data
	v Vehicle
	p Payload
}
