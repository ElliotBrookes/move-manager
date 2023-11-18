package application

import "ElliotBrookes/move-manager/internal/domain"

type OutputManager interface {
	//
	// Functionality for outputting data from application layer
	//

	OutputAll(domain.MoveData) error
	OutputVehicle(domain.Vehicle) error
	OutputAllJourneys([]*domain.Journey) error
	OutputJourney(domain.Journey) error
}

type RetrievalManager interface {
	//
	// Functionality for getting data from external resources
	//

	RetrieveRoute(domain.Address, domain.Address) (domain.Journey, error)
}
