package localize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberFormat(t *testing.T) {
	printer := createPrinter("en", "de")

	output := printer.Sprintf("%.2f", 123456789.92234)

	assert.Equal(t, "123.456.789,92", output)
}
