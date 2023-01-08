package typeParser

import "strconv"

var FloatParser = &floatParser{}

type floatParser struct {
}

func (f *floatParser) Parse(input string) (*float64, error) {
	value, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return nil, generateStandardisedErrorMessage("float")
	}

	return &value, nil
}
