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
	for _, input := range []string{"", "nan", "+Inf", "abc"} {
		if _, err := ParseValue(input); err == nil {
			t.Errorf("ParseValue(%q) unexpectedly succeeded", input)
		}
	}
	if _, _, _, err := Convert(1, "yards", "m"); err == nil {
		t.Error("unknown unit unexpectedly succeeded")
	}
}

func TestConvertRejectsOverflowResult(t *testing.T) {
	if _, _, _, err := Convert(math.MaxFloat64, "km", "m"); err == nil {
		t.Fatal("overflowing conversion unexpectedly succeeded")
	}
	if got, _, _, err := Convert(math.MaxFloat64, "km", "km"); err != nil || got != math.MaxFloat64 {
		t.Fatalf("same-unit conversion = %v, %v", got, err)
	}
	for _, value := range []float64{math.NaN(), math.Inf(1)} {
		if _, _, _, err := Convert(value, "m", "km"); err == nil {
			t.Errorf("Convert(%v) unexpectedly succeeded", value)
		}
	}
}
