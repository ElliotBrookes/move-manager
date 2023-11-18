package configuration

// Plan is for this to be a composable application based on what adapters and ingress methods are available
// Using this to show how the Clean Architechture attempt has allowed them to be changeable without critical
// failures in the application
type MoveManagerApp struct{}

func (a MoveManagerApp) NewApp() {}
