package entities

import (
	"fmt"
	"strings"

	"github.com/Jeffail/gabs"
)

func newGabsObjectLabeler(json string) (SingleObjectLabeler, error) {
	jsonParsed, err := gabs.ParseJSON([]byte(json))
	if err != nil {
		return nil, err
	}

	if checkIfJSONIsArray(jsonParsed) {
		return nil, fmt.Errorf("cannot deal with an array of objects. Must be single object")
	}

	return &gabsLabeler{
		dataContainer: jsonParsed,
	}, nil
}

func checkIfJSONIsArray(dataContainer *gabs.Container) bool {
	parsedString := dataContainer.String()
	return strings.HasPrefix(parsedString, "[") && strings.HasSuffix(parsedString, "]")
}

// gabsLabeler is an implementation of SingleObjectLabeler using the gabs framework
type gabsLabeler struct {
	dataContainer *gabs.Container
	prefix        string
}

func (labeler *gabsLabeler) GetJSON() string {
	return labeler.dataContainer.String()
}

func (labeler *gabsLabeler) ApplyLabel(key string, value string) error {
	_, err := labeler.dataContainer.Set(value, "metadata", "labels", labeler.prefix+key)
	return err
}

func (labeler *gabsLabeler) SetLabelPrefix(prefix string) {
	labeler.prefix = prefix
}
