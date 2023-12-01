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

## 11/23
- Introduced some test files, thinking on how the application will really work, wether it is cmart to have the application layer be multiple structs or if having a single struct would be better
- Thoughts on this are to do with single structs allow for individual testing of a specific struct function, though these are effectively integration tests due to them only calling functions from
- the domain and Other layers

- Another note is that I am trying more to stick to the TDD principles and have introduced test files and will be writing tests before any meaningful implementation.
- This means laying out roughly a testing process for each layer as mentioned in the Clean Architechture book
