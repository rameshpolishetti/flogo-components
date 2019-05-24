# Gateway for graphql endpoints

This is a gateway application demonstrates how following polices can be applied for a graphql endpoint:
* validate GraphQL request against provided schema
* configure Maximum Query Depth

## Installation
* Install [Go](https://golang.org/)
* Install the flogo [cli](https://github.com/project-flogo/cli)

## Setup
```
git clone https://github.com/rameshpolishetti/flogo-components
cd flogo-components/activity/gql/examples/json
```

## Testing
Create the gateway:
```
flogo create -f flogo.json
cd MyProxy
flogo build
```

Start the gateway:
```
bin/MyProxy
```
and test below scenarios:

### Validated graphql request against schema

### Validated graphql request against maxQueryDepth