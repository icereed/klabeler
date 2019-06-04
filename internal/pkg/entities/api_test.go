package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SingleObjectLabeler_GetJSON(t *testing.T) {
	input := `{"simpleObject":1234}`

	labeler, err := NewSingleObjectLabeler(input)

	output := labeler.GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, input, output)
}

func Test_SingleObjectLabeler_NewSingleObjectLabeler_InvalidJSON(t *testing.T) {
	input := `{nope:blahblah}`

	labeler, err := NewSingleObjectLabeler(input)

	assert.Error(t, err)
	assert.Nil(t, labeler)
}

func Test_SingleObjectLabeler_NewSingleObjectLabeler_ErrorOnArray(t *testing.T) {
	input := `[{"simpleObject":1234},{"simpleObject":2345}]`

	labeler, err := NewSingleObjectLabeler(input)

	assert.Error(t, err)
	assert.Nil(t, labeler)
}

func Test_SingleObjectLabeler_AddLabel_ToEmpty(t *testing.T) {
	input := `{}`

	labeler, err := NewSingleObjectLabeler(input)

	labeler.ApplyLabel("testKey", "testValue")

	output := labeler.GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, `{"metadata":{"labels":{"testKey":"testValue"}}}`, output)
}

func Test_SingleObjectLabeler_ApplyLabel_OverrideLabel(t *testing.T) {
	input := `{"metadata":{"labels":{"testKey":"oldValue"}}}`

	labeler, err := NewSingleObjectLabeler(input)

	labeler.ApplyLabel("testKey", "newValue")

	output := labeler.GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, `{"metadata":{"labels":{"testKey":"newValue"}}}`, output)
}

func Test_SingleObjectLabeler_ApplyLabel_AddLabel(t *testing.T) {
	input := `{"metadata":{"labels":{"test1":"value1"}}}`

	labeler, err := NewSingleObjectLabeler(input)

	labeler.ApplyLabel("test2", "value2")

	output := labeler.GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, `{"metadata":{"labels":{"test1":"value1","test2":"value2"}}}`, output)
}

func Test_SingleObjectLabeler_SetPrefix_AddLabel(t *testing.T) {
	input := `{}`

	labeler, err := NewSingleObjectLabeler(input)

	labeler.SetLabelPrefix("some.prefix.io/")
	labeler.ApplyLabel("testKey", "testValue")

	output := labeler.GetJSON()

	assert.NoError(t, err)
	assert.Equal(t, `{"metadata":{"labels":{"some.prefix.io/testKey":"testValue"}}}`, output)
}
