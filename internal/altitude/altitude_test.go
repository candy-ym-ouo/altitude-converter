package altitude

import (
	"math"
	"testing"
)

func TestConvert(t *testing.T) {
	got, from, to, err := Convert(1000, "m", "ft")
	if err != nil || from != "m" || to != "ft" || math.Abs(got-3280.839895013123) > 1e-9 {
		t.Fatalf("Convert() = %v, %q, %q, %v", got, from, to, err)
	}
}

func TestInvalidInput(t *testing.T) {
	for _, input := range []string{"", " \t\r\n", "\u00a0", "nan", "+Inf", "-Inf", "1e309", "abc"} {
		if _, err := ParseValue(input); err == nil {
			t.Errorf("ParseValue(%q) unexpectedly succeeded", input)
		}
	}
	if _, _, _, err := Convert(1, "yards", "m"); err == nil {
		t.Error("unknown unit unexpectedly succeeded")
	}
}

func TestParseValueAcceptsSurroundingWhitespace(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  float64
	}{
		{" \t1250\n", 1250},
		{"\u00a0-1.25\u00a0", -1.25},
		{" 1.7976931348623157e+308 ", math.MaxFloat64},
	} {
		got, err := ParseValue(tc.input)
		if err != nil || got != tc.want {
			t.Errorf("ParseValue(%q) = %v, %v; want %v, nil", tc.input, got, err, tc.want)
		}
	}
}
