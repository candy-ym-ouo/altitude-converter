# Bug Reproduction

## Baseline

This reproduction is based on branch `base_bug_004`.

## Steps

1. Format a value with more than six meaningful fractional digits, such as `1234.567890123`.
2. Parse the formatted text back as a `float64`.

## Actual Result

The formatter always emits six digits after the decimal point, truncating significant precision. Parsing that output does not recover the original value.

## Expected Result

Formatting must preserve enough significant precision for a `float64` to round-trip exactly, without adding unnecessary trailing zeroes to integers or exact decimals.

## Focused Check

```bash
go test ./internal/altitude -run TestFormatPreservesSignificantPrecision -count=1
```

On the `base_bug_004` baseline, high-precision values are truncated to six decimal places.
