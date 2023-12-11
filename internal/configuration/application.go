package configuration

import (
	"ElliotBrookes/move-manager/internal/adapters/store"
	"ElliotBrookes/move-manager/internal/application"
	"ElliotBrookes/move-manager/internal/ports"
)

// Plan is for this to be a composable application based on what adapters and ingress methods are available
// Using this to show how the Clean Architechture attempt has allowed them to be changeable without critical
// failures in the application
type MoveManagerApp struct {
	name           string
	portManager    ports.PortManager
	outputManager  application.OutputManager // change this to adapter dependancy
	kvstore        store.KVStore
	routeFetcher   application.RouteRetriever
	vehicleFetcher application.VehicleRetriever
}

func New() *MoveManagerApp {
	return &MoveManagerApp{}
}

func (a MoveManagerApp) Start() {
	a.portManager.Handle()
}

func (a MoveManagerApp) Name(n string) {
	a.name = n
}

func (a MoveManagerApp) PortManager(pm ports.PortManager) {
	a.portManager = pm
}

func (a MoveManagerApp) OutputManager(om application.OutputManager) {
	a.outputManager = om
}

func (a MoveManagerApp) RetrievalManager(rm application.RetrievalManager) {
	a.retrievalManager = rm
}
