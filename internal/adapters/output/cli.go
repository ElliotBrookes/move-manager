package output

import "ElliotBrookes/move-manager/internal/domain"

type CliEgressManager struct {
	// Private fields used for storage
}

func RenderTemplate() error {
	// Choose template, load and execute
	return nil
}

func (r CliEgressManager) OutputAll(domain.MoveData) error {
	return nil
}

func (r CliEgressManager) OutputVehicle(domain.Vehicle) error {
	return nil
}

func (r CliEgressManager) OutputAllJourneys([]*domain.Journey) error {
	return nil
}

func (r CliEgressManager) OutputJourney(domain.Journey) error {
	return nil
}
