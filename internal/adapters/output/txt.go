package output

import "ElliotBrookes/move-manager/internal/domain"

type TxtFileEgressManager struct{}

func (r TxtFileEgressManager) OutputAll(domain.MoveData) error {
	return nil
}

func (r TxtFileEgressManager) OutputVehicle(domain.Vehicle) error {
	return nil
}

func (r TxtFileEgressManager) OutputAllJourneys([]*domain.Journey) error {
	return nil
}

func (r TxtFileEgressManager) OutputJourney(domain.Journey) error {
	return nil
}
