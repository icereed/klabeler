package entities

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

// NewSingleObjectLabeler creates a new instance of a GenericLabeler capable of dealing with one single object
func NewSingleObjectLabeler(json string) (GenericLabeler, error) {
	labeler, err := newGabsObjectLabeler(json)
	return labeler, err
}

func newGabsObjectLabeler(json string) (GenericLabeler, error) {
	if checkIfJSONIsArray(json) {
		return nil, fmt.Errorf("cannot deal with an array of objects. Must be single object")
	}

	jsonParsed, err := gabs.ParseJSON([]byte(json))
	if err != nil {
		return nil, err
	}

	return &gabsLabeler{
		dataContainer: jsonParsed,
	}, nil
}

// gabsLabeler is an implementation of GenericLabeler using the gabs framework
type gabsLabeler struct {
	dataContainer *gabs.Container
	prefix        string
}

func (labeler *gabsLabeler) GetJSON() string {
	return labeler.dataContainer.String()
}

func (labeler *gabsLabeler) ApplyLabel(key string, value string) {
	labeler.dataContainer.Set(value, "metadata", "labels", labeler.prefix+key)

}

func (labeler *gabsLabeler) SetLabelPrefix(prefix string) {
	labeler.prefix = prefix
}
