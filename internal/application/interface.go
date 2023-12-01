package application

//
// Main interface for application/ use case structures will implement to be called by other services
//

type appInstance interface {
	Execute(OutputManager, RetrievalManager, StorageManager) error
}
