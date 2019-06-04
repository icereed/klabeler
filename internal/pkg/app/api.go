package app

// KLabeler represents the main use-cases for the app
type KLabeler interface {
	ApplyCurrentGitHash() KLabeler
	ApplyLabel(key string, value string) KLabeler
	SetLabelPrefix(prefix string) KLabeler
	GetJSON() string
	GetYAML() string
}
