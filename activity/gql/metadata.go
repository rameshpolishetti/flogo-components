package gql

import "github.com/project-flogo/core/data/coerce"

// Input input meta data
type Input struct {
	Query         string `md:"query"`
	SchemaFile    string `md:"schemaFile"`
	MaxQueryDepth int    `md:"maxQueryDepth"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"query":         i.Query,
		"schemaFile":    i.SchemaFile,
		"maxQueryDepth": i.MaxQueryDepth,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.Query, err = coerce.ToString(values["query"])
	if err != nil {
		return err
	}
	i.SchemaFile, err = coerce.ToString(values["schemaFile"])
	if err != nil {
		return err
	}
	i.MaxQueryDepth, err = coerce.ToInt(values["maxQueryDepth"])
	if err != nil {
		return err
	}

	return nil
}

type Output struct {
	Valid             bool   `md:"valid"`
	ValidationMessage string `md:"validationMessage"`
	Error             bool   `md:"error"`
	ErrorMessage      string `md:"errorMessage"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	valid, err := coerce.ToBool(values["valid"])
	if err != nil {
		return err
	}
	o.Valid = valid
	o.ValidationMessage = values["validationMessage"].(string)
	o.Error = values["error"].(bool)
	o.ErrorMessage = values["errorMessage"].(string)
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"valid":             o.Valid,
		"validationMessage": o.ValidationMessage,
		"error":             o.Error,
		"errorMessage":      o.ErrorMessage,
	}
}
