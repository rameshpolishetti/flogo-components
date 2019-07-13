package spectrum

import (
	"context"
	"fmt"

	"github.com/project-flogo/core/action"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

func init() {
	fmt.Println("INIT")
	_ = action.Register(&Action{}, &ActionFactory{})
}

var actionMd = action.ToMetadata(&Settings{}, &Input{}, &Output{})

type ActionFactory struct {
}

func (f *ActionFactory) Initialize(ctx action.InitContext) error {
	fmt.Println("ActionFactory: Initialize()")
	return nil
}

func (f *ActionFactory) New(config *action.Config) (action.Action, error) {
	fmt.Println("ActionFactory: New()")

	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	fmt.Printf("SSETTINGS: %s", s.ASetting)
	log.RootLogger().Debugf("Setting: %s", s.ASetting)

	act := &Action{settings: s}

	return act, nil

}

type Action struct {
	settings *Settings
}

// Metadata implements action.Action.Metadata
func (a *Action) Metadata() *action.Metadata {
	fmt.Println("Action: Metadata()")
	return actionMd
}

// IOMetadata implements action.Action.IOMetadata
func (a *Action) IOMetadata() *metadata.IOMetadata {
	fmt.Println("Action: IOMetadata()")
	return nil
}

// Run implements action.SyncAction.Run
func (a *Action) Run(ctx context.Context, inputValues map[string]interface{}) (map[string]interface{}, error) {
	fmt.Println("Action: Run()")

	input := &Input{}
	err := input.FromMap(inputValues)
	if err != nil {
		return nil, err
	}

	log.RootLogger().Infof("Input: %s", input.AnInput)

	output := &Output{AnOutput: input.AnInput}

	return output.ToMap(), nil
}
