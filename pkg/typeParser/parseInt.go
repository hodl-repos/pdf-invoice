package typeParser

import "strconv"

var IntParser = &intParser{}

type intParser struct {
}

func (i *intParser) Parse(input string) (*int64, error) {
	value, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		return nil, generateStandardisedErrorMessage("int")
	}

	return &value, nil
}
