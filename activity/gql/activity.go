package gql

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

func init() {
	_ = activity.Register(&Activity{})
}

// Input input meta data
type Input struct {
	Message string `md:"message"` // The message to log
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message": i.Message,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Message, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}

	return nil
}

type Output struct {
	AvailableLimit int `md:"availableLimit"` // available limit
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

	msg := input.Message

	ctx.Logger().Info(msg)

	// set output
	err = ctx.SetOutput("availableLimit", int(10))
	if err != nil {
		return false, err
	}

	return true, nil
}
