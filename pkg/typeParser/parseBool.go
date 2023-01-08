package typeParser

import "strconv"

var BoolParser = &boolParser{}

type boolParser struct {
}

func (f *boolParser) Parse(input string) (*bool, error) {
	value, err := strconv.ParseBool(input)

	if err != nil {
		return nil, generateStandardisedErrorMessage("bool")
	}

	return &value, nil
}
