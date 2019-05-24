package gql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{})
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is an GQLActivity
// inputs : {message}
// outputs: none
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - TBD
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	fmt.Println("Evaluate graphQL policies")
	// get inputs
	input := &Input{}
	ctx.GetInputObject(input)
	fmt.Println("query: ", input.Query)
	fmt.Println("schema file: ", input.SchemaFile)

	// check schema file is provided or not
	if input.SchemaFile == "" {
		// set error flag & error message in the output
		err = ctx.SetOutput("error", true)
		if err != nil {
			return false, err
		}
		errMsg := "Schema file is required"
		err = ctx.SetOutput("errorMessage", errMsg)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	// load schema
	schemaStr, err := ioutil.ReadFile(input.SchemaFile)
	if err != nil {
		fmt.Printf("Not able to read the schema file[%s] with the error - %s \n", input.SchemaFile, err)
		// set error flag & error message in the output
		err = ctx.SetOutput("error", true)
		if err != nil {
			return false, err
		}
		errMsg := fmt.Sprintf("Not able to read the schema file[%s] with the error - %s \n", input.SchemaFile, err)
		err = ctx.SetOutput("errorMessage", errMsg)
		if err != nil {
			return false, err
		}

		return true, nil
	}
	schema, err := graphql.ParseSchema(string(schemaStr), nil)
	if err != nil {
		fmt.Println("Error while parsing GQL schema: ", err)
		// set error flag & error message in the output
		err = ctx.SetOutput("error", true)
		if err != nil {
			return false, err
		}
		errMsg := fmt.Sprintf("Error while parsing GQL schema: %s", err)
		err = ctx.SetOutput("errorMessage", errMsg)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	// parse request
	var gqlQuery struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	err = json.Unmarshal([]byte(input.Query), &gqlQuery)
	if err != nil {
		fmt.Println("Error while parsing GQL query: ", err)
		// set error flag & error message in the output
		err = ctx.SetOutput("error", true)
		if err != nil {
			return false, err
		}
		errMsg := fmt.Sprintf("Not a valid graphQL request. Details: %s", err)
		err = ctx.SetOutput("errorMessage", errMsg)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	// check query depth
	depth := calculateQueryDepth(gqlQuery.Query)
	if depth > input.MaxQueryDepth {
		// set error flag & error message in the output
		err = ctx.SetOutput("error", true)
		if err != nil {
			return false, err
		}
		errMsg := fmt.Sprintf("graphQL request query depth[%v] is exceeded allowed maxQueryDepth[%v]", depth, input.MaxQueryDepth)
		fmt.Println(errMsg)
		err = ctx.SetOutput("errorMessage", errMsg)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	// validate request
	validationErrors := schema.Validate(gqlQuery.Query)
	if validationErrors != nil {
		fmt.Printf("Invalid GQL request: %s \n", validationErrors)

		// set error flag & error message in the output
		err = ctx.SetOutput("error", true)
		if err != nil {
			return false, err
		}
		errMsg := fmt.Sprintf("Not a valid graphQL request. Details: %s", validationErrors)
		err = ctx.SetOutput("errorMessage", errMsg)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	// set output
	err = ctx.SetOutput("valid", true)
	if err != nil {
		return false, err
	}
	validationMsg := fmt.Sprintf("Valid graphQL query. query = %s\n type = Query \n queryDepth = %v", input.Query, depth)
	err = ctx.SetOutput("validationMessage", validationMsg)
	if err != nil {
		return false, err
	}

	return true, nil
}
