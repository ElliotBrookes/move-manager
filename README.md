# move-manager

## Resources
- Clean architechture Book ofc
- https://github.com/GSabadini/go-clean-architecture/blob/master/ -> Good example of how Application layer should be used by Ingress/ infrastructure layer

## Notes to remember/ keep in mind
- Architecture layer does the data retrieval, the adapters are doing data translation for the other layers
- Im going to make a config layer as well, this will then pass down all relevant obj to allow layers to invoke others indirectly via interfaces

## Notes on structure and structs :p
- Configuration layer is going to be responsible for having presets and stringing together the other layers
- it will also cause them to run ^

- The infrastructure layer defines different technologies used in different places e.g. the offlineDataset and the MoveData
- these two are used for different things within the application, however they are both KVStores as defined in the adapters and behave as such
- handlers defined in the adapter layer will define their uses and behaviour, though they could both be used with other adapters - their functionality would just be limited/ lacking

- Currently a type of technology is defined in the adapters, which is defined more literally in the infrastructure
- this technology type is then paired with an adapter in the adapter layer which is passed to the application layer to execute functionality

## Notes on execution responsibility
- The configuration will contain a starting point method which will need to be envoked with the various args required to start execution of the programs logic.
- More of the different appication logic is contained in the other areas of ther codebase and the handling of the various execution steps will be contained within this logic, meaning that the configuration object is only calling a single objects method
