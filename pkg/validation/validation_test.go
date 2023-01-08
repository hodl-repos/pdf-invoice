package validation

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type testValidationStruct struct {
	Integer          *int       `json:"integer" validate:"required"`
	UniqueIdentifier *uuid.UUID `json:"uniqueIdentifier" validate:"required"`
	String           *string    `json:"string" validate:"required"`
	Float            *float64   `json:"float" validate:"required"`
}

func TestAllNilData(t *testing.T) {
	data := testValidationStruct{
		Integer:          nil,
		UniqueIdentifier: nil,
		String:           nil,
		Float:            nil,
	}

	err := ValidateStruct(&data)

	assert.NotNil(t, err)

	val, ok := err.(*ValidationError)
	assert.True(t, ok)

	assert.Equal(t, 4, len(val.InvalidParams))

	for i := 0; i < 4; i++ {
		assert.Equal(t, "is required", val.InvalidParams[i].Reason)
	}
}

func TestAllFilled(t *testing.T) {
	integer := 1
	unique := uuid.New()
	strings := "hi"
	float := 12.2

	data := testValidationStruct{
		Integer:          &integer,
		UniqueIdentifier: &unique,
		String:           &strings,
		Float:            &float,
	}

	err := ValidateStruct(&data)

	assert.Nil(t, err)
}
