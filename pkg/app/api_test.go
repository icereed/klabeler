package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewKLabeler_NotNil(t *testing.T) {
	input := `{"simpleObject":1234}`

	labeler, err := NewKLabeler(input)

	assert.NoError(t, err)
	assert.NotNil(t, labeler)
}

func Test_NewKLabeler_GetJSON(t *testing.T) {
	input := `{"simpleObject":1234}`

	labeler, err := NewKLabeler(input)

	output := labeler.GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, input, output)
}

func Test_NewKLabeler_YAML_Input(t *testing.T) {
	input := `simpleObject: 1234`

	labeler, err := NewKLabeler(input)
	output := labeler.GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, `{"simpleObject":1234}`, output)
}

func Test_NewKLabeler_Array_Input(t *testing.T) {
	input := `[{"simpleObject":1234},{"simpleObject":2345}]`

	labeler, err := NewKLabeler(input)
	output := labeler.GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, input, output)
}

func Test_KLabeler_ApplyLabel(t *testing.T) {
	input := `{}`

	labeler, err := NewKLabeler(input)

	output := labeler.ApplyLabel("testKey", "testValue").GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, `{"metadata":{"labels":{"testKey":"testValue"}}}`, output)
}

func Test_KLabeler_ApplyLabel_YAML(t *testing.T) {
	input := `{}`

	labeler, err := NewKLabeler(input)

	output := labeler.ApplyLabel("testKey", "testValue").GetYAML()

	assert.NoError(t, err)
	assert.Equal(t, "metadata:\n  labels:\n    testKey: testValue\n", output)
}

func Test_KLabeler_ApplyLabelToArray(t *testing.T) {
	input := `[{},{}]`

	labeler, err := NewKLabeler(input)

	output := labeler.ApplyLabel("testKey", "testValue").GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, `[{"metadata":{"labels":{"testKey":"testValue"}}},{"metadata":{"labels":{"testKey":"testValue"}}}]`, output)
}

func Test_KLabeler_ApplyLabelToArray_PrefixSet(t *testing.T) {
	input := `[{},{}]`

	labeler, err := NewKLabeler(input)

	output := labeler.SetLabelPrefix("some.io/").ApplyLabel("testKey", "testValue").GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, `[{"metadata":{"labels":{"some.io/testKey":"testValue"}}},{"metadata":{"labels":{"some.io/testKey":"testValue"}}}]`, output)
}

func Test_KLabeler_ApplyCurrentGitHash(t *testing.T) {
	input := `[{},{}]`

	labeler, err := NewKLabelerWithGitHashProvider(input, &testGitHashProvider{
		echoHash: "c0ffee",
	})
	output := labeler.ApplyCurrentGitHash().GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, `[{"metadata":{"labels":{"git-hash":"c0ffee"}}},{"metadata":{"labels":{"git-hash":"c0ffee"}}}]`, output)
}

func Test_KLabeler_MultiYAMLInput(t *testing.T) {
	input := `number: 1
---
number: 2`

	labeler, err := NewKLabeler(input)
	output := labeler.GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, `[{"number":1},{"number":2}]`, output)
}

type testGitHashProvider struct {
	echoHash string
}

func (gitProvider *testGitHashProvider) getCurrentGitHash() string {
	return gitProvider.echoHash
}
