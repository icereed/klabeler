package app

import (
	"strings"

	"github.com/icereed/klabeler/internal/pkg/entities"
	"sigs.k8s.io/yaml"
)


func yamlToJSON(jsonOrYaml string) (string, error) {
	splittedContents := strings.Split(jsonOrYaml, "\n---")

	if len(splittedContents) < 2 {
		json, err := yaml.YAMLToJSON([]byte(jsonOrYaml))
		return string(json), err
	}

	var jsonParts []string
	for _, block := range splittedContents {
		json, err := yaml.YAMLToJSON([]byte(block))
		if err != nil {
			return "", err
		}
		jsonParts = append(jsonParts, string(json))
	}

	return "[" + strings.Join(jsonParts, ",") + "]", nil
}

type klabelerImpl struct {
	dataJSON        string
	prefix          string
	gitHashProvider GitHashProvider
}

func (labeler *klabelerImpl) GetJSON() string {
	return labeler.dataJSON
}

func (labeler *klabelerImpl) GetYAML() string {
	yaml, err := yaml.JSONToYAML([]byte(labeler.GetJSON()))

	if err != nil {
		panic(err)
	}

	return string(yaml)
}

func (labeler *klabelerImpl) SetLabelPrefix(prefix string) KLabeler {
	labeler.prefix = prefix
	return labeler
}

func (labeler *klabelerImpl) ApplyCurrentGitHash() KLabeler {
	return labeler.ApplyLabel("git-hash", labeler.gitHashProvider.getCurrentGitHash())
}

func (labeler *klabelerImpl) ApplyLabel(key string, value string) KLabeler {

	genericLabeler, err := entities.NewGenericLabeler(labeler.GetJSON())
	if err != nil {
		panic(err)
	}
	genericLabeler.SetLabelPrefix(labeler.prefix)
	genericLabeler.ApplyLabel(key, value)

	labeler.dataJSON = genericLabeler.GetJSON()
	return labeler
}
