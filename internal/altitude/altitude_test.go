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

func TestConvertAcceptsSpacedUnitAlias(t *testing.T) {
	for _, unit := range []string{"nautical mile", " nautical mile ", "nautical\tmile", "nautical\u00a0mile"} {
		got, from, to, err := Convert(1, unit, "m")
		if err != nil || from != "nm" || to != "m" || got != 1852 {
			t.Errorf("Convert(1, %q, %q) = %v, %q, %q, %v", unit, "m", got, from, to, err)
		}
	}
	got, from, to, err := Convert(1852, "m", " nautical mile ")
	if err != nil || from != "m" || to != "nm" || got != 1 {
		t.Fatalf("Convert() = %v, %q, %q, %v", got, from, to, err)
	}
}

func TestNormalizeUnitPreservesAliases(t *testing.T) {
	for alias, want := range aliases {
		got, err := NormalizeUnit(" \t" + alias + "\n")
		if err != nil || got != want {
			t.Errorf("NormalizeUnit(%q) = %q, %v; want %q, nil", alias, got, err, want)
		}
	}
	if _, err := NormalizeUnit(" \t\u00a0\n"); err == nil {
		t.Error("NormalizeUnit() unexpectedly accepted whitespace-only input")
	}
}
