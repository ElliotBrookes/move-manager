package output

import "ElliotBrookes/move-manager/internal/domain"

type HtmlEgressManager struct{}

func (r HtmlEgressManager) OutputAll(domain.MoveData) error {
	return nil
}

func (r HtmlEgressManager) OutputVehicle(domain.Vehicle) error {
	return nil
}

func (r HtmlEgressManager) OutputAllJourneys([]*domain.Journey) error {
	return nil
}

func (r HtmlEgressManager) OutputJourney(domain.Journey) error {
	return nil
}
