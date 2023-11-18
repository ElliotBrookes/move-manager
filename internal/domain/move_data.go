package domain

import "time"

type MoveData struct {
	//
	// Struct holding data for all move information
	//

	// Static
	cost int

	// Dynamic
	driveTime time.Duration

	Journeys []*Journey
	Vehicle  []*Vehicle
}
