# GQL

The `gql` service type accepts GraphQL request and applies policies and validates against the schema.

The service `settings` and available `input` for the request are as follows:

| Name   |  Type   | Description   |
|:-----------|:--------|:--------------|
| query | input | GraphQL request string |
| schemaFile | input | GraphQL schema file path |
| maxQueryDepth | input | Maximum allowed GraphQL query depth |

The available response outputs are as follows:

| Name   |  Type   | Description   |
|:-----------|:--------|:--------------|
| valid | boolean | `true` if the GraphQL query is valid |
| error | boolean | `true` if any error occured while inspecting the GraphQL query  |
| errorMessage | string | The error message |

A sample `service` definition is:

```json
{
    "name": "GQL",
    "description": "GraphQL policies service",
    "ref": "github.com/rameshpolishetti/flogo-components/activity/gql"
}
```

An example `step` that invokes `JQL` service using a `GraphQL request` from a HTTP trigger is:

```json
{
    "service": "GQL",
    "input": {
        "query": "=$.payload.content",
        "schemaFile": "schema.graphql",
        "maxQueryDepth": 2
    }
}
```

Utilizing and extracting the response values can be seen in a conditional evaluation:

```json
{
    "if": "$.GQL.outputs.error == true",
    "error": true,
    "output": {
        "code": 200,
        "data": {
            "error": "=$.GQL.outputs.errorMessage"
        }
    }
}
```
## Maximum Query Depth
This policy allows to prevent clients from abusing deep query depth, Knowing your schema might give you an idea of how deep a legitimate query can go.
example bad query:
```sh
query badquery {            #depth 0
  author() {                #depth 1
    posts {                 #depth 2
      author {              #depth 3
        posts {             #depth 4
          author {          #depth 5
          }
        }
      }
    }
  }
}
```
gateway configured with `maxQueryDepth` to 3 would consider above query too deep and the query is invalid.


## TODO
* Policy based on GraphQL query complexity
* Throttling Based on Server Time
* Throttling Based on Query Complexity
