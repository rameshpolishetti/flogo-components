package gql

import (
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

func init() {
	_ = activity.Register(&Activity{})
}

// Input input meta data
type Input struct {
	Request string `md:"request"` // request string
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"request": i.Request,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Request, err = coerce.ToString(values["request"])
	if err != nil {
		return err
	}

	return nil
}

type Output struct {
	QueryDepth int `md:"queryDepth"` // query depth
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

	input := &Input{}
	ctx.GetInputObject(input)

	msg := fmt.Sprintf("GraphQL request: %s \n type: Query \n query depth: 10", input.Request)

	ctx.Logger().Info(msg)

	// set output
	err = ctx.SetOutput("queryDepth", int(10))
	if err != nil {
		return false, err
	}

	return true, nil
}
