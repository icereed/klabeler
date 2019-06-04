package entities

// GenericLabeler represents the basic entity for any logic regarding labeling
type GenericLabeler interface {
	SetLabelPrefix(prefix string)
	ApplyLabel(key string, value string)
	GetJSON() string
}

// NewGenericLabeler creates a new instance of a GenericLabeler capable of dealing with single objects or arrays
func NewGenericLabeler(json string) (GenericLabeler, error) {
	isArray := checkIfJSONIsArray(json)

	if isArray {
		return NewMultiObjectLabeler(json)
	}

	return NewSingleObjectLabeler(json)
}
