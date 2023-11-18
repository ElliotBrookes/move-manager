package retrieval

import "ElliotBrookes/move-manager/internal/domain"

type ExternalRetriever struct {
	//
	// Manager that implements all functionality for retrieving external data/information
	//

}

func (r ExternalRetriever) GetRoute(from domain.Address, to domain.Address) (domain.Journey, error) {
	// data race between multiple different data sources if possible
	getMapBoxRoute()

	return domain.Journey{}, nil
}
