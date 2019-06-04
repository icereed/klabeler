package entities

import (
	"fmt"
	"strings"

	"github.com/Jeffail/gabs"
)

// NewMultiObjectLabeler creates a new instance of a GenericLabeler capable of dealing with object arrays
func NewMultiObjectLabeler(json string) (GenericLabeler, error) {

	labeler, err := newCompositeLabeler(json)
	return labeler, err
}

func newCompositeLabeler(json string) (GenericLabeler, error) {
	if !checkIfJSONIsArray(json) {
		return nil, fmt.Errorf("compositeLabeler can only deal with an input of array")
	}

	jsonParsed, err := gabs.ParseJSON([]byte(json))
	if err != nil {
		return nil, err
	}

	children, err := jsonParsed.Children()
	if err != nil {
		return nil, err
	}

	var labelers []GenericLabeler

	for _, child := range children {
		childLabeler, err := NewSingleObjectLabeler(child.String())
		if err != nil {
			return nil, err
		}
		labelers = append(labelers, childLabeler)
	}

	return &compositeLabeler{
		labelers: labelers,
	}, nil
}

type compositeLabeler struct {
	labelers []GenericLabeler
	prefix   string
}

func (labeler *compositeLabeler) GetJSON() string {
	var jsonOutputs []string

	for _, singleLabeler := range labeler.labelers {
		jsonOutputs = append(jsonOutputs, singleLabeler.GetJSON())
	}

	return "[" + strings.Join(jsonOutputs, ",") + "]"
}

func (labeler *compositeLabeler) ApplyLabel(key string, value string) {
	for _, singleLabeler := range labeler.labelers {
		singleLabeler.ApplyLabel(key, value)
	}
}

func (labeler *compositeLabeler) SetLabelPrefix(prefix string) {
	for _, singleLabeler := range labeler.labelers {
		singleLabeler.SetLabelPrefix(prefix)
	}
}
