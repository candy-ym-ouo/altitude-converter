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
	for _, input := range []string{"", " \t\n", "nan", " \tNaN\n", " \t+Inf\n", "abc"} {
		if _, err := ParseValue(input); err == nil {
			t.Errorf("ParseValue(%q) unexpectedly succeeded", input)
		}
	}
	if _, _, _, err := Convert(1, "yards", "m"); err == nil {
		t.Error("unknown unit unexpectedly succeeded")
	}
}

func TestParseValueAcceptsSurroundingWhitespace(t *testing.T) {
	got, err := ParseValue("\u00a0 \t1250\n\u00a0")
	if err != nil || got != 1250 {
		t.Fatalf("ParseValue() = %v, %v", got, err)
	}
}

func TestParseValueNumericBounds(t *testing.T) {
	if got, err := ParseValue(strconv.FormatFloat(math.MaxFloat64, 'g', -1, 64)); err != nil || got != math.MaxFloat64 {
		t.Fatalf("ParseValue(max finite value) = %v, %v", got, err)
	}
	if _, err := ParseValue("1.7976931348623159e308"); err == nil {
		t.Error("ParseValue(overflow) unexpectedly succeeded")
	}
}
