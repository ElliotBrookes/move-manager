package domain

type Journey struct {
	//
	// Information regarding a portion of driving in a move
	//

	// Static size data
	from     *Location
	to       *Location
	distance int32

	v *Vehicle
	p *Payload
	// Dynamically Sized data
}

func (j Journey) SetDistance(d int32) {
	j.distance = d
}
