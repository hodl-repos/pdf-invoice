package utils

import (
	"fmt"
	"strings"
)

func SsnFormatter(ssn string) string {
	part1 := ssn[0:4]
	part2 := ssn[4:]
	return fmt.Sprintf("%s %s", part1, part2)
}

func IsoDateFormatter(isoDate string, formatStr string) string {
	if isoDate == "" {
		return ""
	}

	parts := strings.Split(isoDate, "-")
	y := parts[0]
	m := parts[1]
	d := parts[2]
	switch formatStr {
	case "dd.mm.yyyy":
		return fmt.Sprintf("%s.%s.%s", d, m, y)
	default:
		return fmt.Sprintf("%s-%s-%s", y, m, d)
	}
}

func CentToString(cent int) string {
	f := float64(cent) / 100.0
	return strings.ReplaceAll(fmt.Sprintf("%.2f", f), ".", ",")
}
