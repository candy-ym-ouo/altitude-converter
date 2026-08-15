package altitude

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Units supported by the converter. Values are meters per unit.
var units = map[string]float64{
	"m":  1,
	"ft": 0.3048,
	"km": 1000,
	"nm": 1852,
}

var aliases = map[string]string{
	"m": "m", "meter": "m", "meters": "m", "metre": "m", "metres": "m",
	"ft": "ft", "foot": "ft", "feet": "ft",
	"km": "km", "kilometer": "km", "kilometers": "km", "kilometre": "km", "kilometres": "km",
	"nm": "nm", "nmi": "nm", "nauticalmile": "nm", "nauticalmiles": "nm",
}

// NormalizeUnit returns the canonical abbreviation for a supported unit.
func NormalizeUnit(unit string) (string, error) {
	key := strings.ToLower(strings.Join(strings.Fields(unit), ""))
	canonical, ok := aliases[key]
	if !ok {
		return "", fmt.Errorf("unsupported unit %q (use m, ft, km, or nm)", unit)
	}
	return canonical, nil
}

// ParseValue rejects blank, non-numeric, infinite, and NaN input.
func ParseValue(text string) (float64, error) {
	value, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
	if err != nil || math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, fmt.Errorf("invalid altitude value %q", text)
	}
	return value, nil
}

// Convert converts a value between supported altitude units.
func Convert(value float64, from, to string) (float64, string, string, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, "", "", fmt.Errorf("altitude must be finite")
	}
	from, err := NormalizeUnit(from)
	if err != nil {
		return 0, "", "", err
	}
	to, err = NormalizeUnit(to)
	if err != nil {
		return 0, "", "", err
	}
	result := value * units[from] / units[to]
	if math.IsInf(result, 0) {
		return 0, "", "", fmt.Errorf("conversion result is outside the supported numeric range")
	}
	return result, from, to, nil
}

// Format gives a stable, readable representation without gratuitous trailing zeroes.
func Format(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
