package app

// KLabeler represents the main use-cases for the app
type KLabeler interface {
	ApplyCurrentGitHash() KLabeler
	ApplyLabel(key string, value string) KLabeler
	SetLabelPrefix(prefix string) KLabeler
	GetJSON() string
	GetYAML() string
}

// NewKLabeler creates a new instance from JSON or YAML input
func NewKLabeler(jsonOrYaml string) (KLabeler, error) {
	return NewKLabelerWithGitHashProvider(jsonOrYaml, &goGitHashProvider{})
}

// NewKLabelerWithGitHashProvider creates a new instance from JSON or YAML input with the given GitHashProvider
func NewKLabelerWithGitHashProvider(jsonOrYaml string, gitHashProvider GitHashProvider) (KLabeler, error) {

	json, err := yamlToJSON(jsonOrYaml)

	if err != nil {
		return nil, err
	}

	return &klabelerImpl{
		dataJSON:        string(json),
		gitHashProvider: gitHashProvider,
	}, nil
}
