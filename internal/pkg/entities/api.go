package entities

// SingleObjectLabeler represents the basic entity for any logic regarding labeling
type SingleObjectLabeler interface {
	SetLabelPrefix(prefix string)
	ApplyLabel(key string, value string) error
	GetJSON() string
}

// NewSingleObjectLabeler creates a new instance of a SingleObjectLabeler
func NewSingleObjectLabeler(json string) (SingleObjectLabeler, error) {
	labeler, err := newGabsObjectLabeler(json)
	return labeler, err
}
