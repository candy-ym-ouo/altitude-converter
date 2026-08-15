package altitude

import (
	"math"
	"strconv"
	"testing"
)

func TestConvert(t *testing.T) {
	got, from, to, err := Convert(1000, "m", "ft")
	if err != nil || from != "m" || to != "ft" || math.Abs(got-3280.839895013123) > 1e-9 {
		t.Fatalf("Convert() = %v, %q, %q, %v", got, from, to, err)
	}
}

func TestInvalidInput(t *testing.T) {
	for _, input := range []string{"", "nan", "+Inf", "abc"} {
		if _, err := ParseValue(input); err == nil {
			t.Errorf("ParseValue(%q) unexpectedly succeeded", input)
		}
	}
	if _, _, _, err := Convert(1, "yards", "m"); err == nil {
		t.Error("unknown unit unexpectedly succeeded")
	}
}

func TestFormatPreservesSignificantPrecision(t *testing.T) {
	tests := []struct {
		value float64
		want  string
	}{
		{1000, "1000"},
		{12.5, "12.5"},
		{1234.567890123, "1234.567890123"},
		{math.Nextafter(1, 2), "1.0000000000000002"},
		{0.0000001, "0.0000001"},
	}
	for _, test := range tests {
		got := Format(test.value)
		if got != test.want {
			t.Errorf("Format(%v) = %q, want %q", test.value, got, test.want)
		}
		parsed, err := strconv.ParseFloat(got, 64)
		if err != nil || parsed != test.value {
			t.Errorf("Format(%v) = %q does not round-trip: %v", test.value, got, err)
		}
	}
}
