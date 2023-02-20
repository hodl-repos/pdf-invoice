package delimitor

type Delimitor int

const (
	NewLine Delimitor = iota + 1
	Tab               //uses 6 spaces instead of tab -> pdf wants it that way
	Space
)

func (s Delimitor) String() string {
	return [...]string{"\n", "      ", " "}[s-1]
}
