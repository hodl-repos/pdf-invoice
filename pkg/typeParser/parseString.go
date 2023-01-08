package typeParser

var StringParser = &stringParser{}

type stringParser struct {
}

func (f *stringParser) Parse(input string) (*string, error) {
	return &input, nil
}
