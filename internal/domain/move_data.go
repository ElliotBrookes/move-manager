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

func (m MoveData) Distance() int32 {
	var totalDistance int32

	for _, journey := range m.Journeys {
		totalDistance += journey.Distance
	}

	return totalDistance
}
