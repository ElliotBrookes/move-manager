# move-manager

## Resources
- Clean architechture Book ofc
- https://github.com/GSabadini/go-clean-architecture/blob/master/ -> Good example of how Application layer should be used by Ingress/ infrastructure layer

## Notes to remember/ keep in mind
- Architecture layer does the data retrieval, the adapters are doing data translation for the other layers
- Im going to make a config layer as well, this will then pass down all relevant obj to allow layers to invoke others indirectly via interfaces
