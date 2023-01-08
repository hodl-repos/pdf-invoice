package typeParser

import (
	"github.com/google/uuid"
)

var UUIDParser = &uuidParser{}

type uuidParser struct {
}

func (u *uuidParser) Parse(input string) (*uuid.UUID, error) {
	value, err := uuid.Parse(input)

	if err != nil {
		return nil, generateStandardisedErrorMessage("uuid")
	}

	return &value, nil
}
