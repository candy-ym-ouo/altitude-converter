# Bug Reproduction

## Baseline

This reproduction is based on branch `base_bug_002`.

## Steps

1. Pass an otherwise valid altitude value with surrounding whitespace, such as `" \t1250\n"`, to the conversion path.
2. Use valid source and target units.

## Actual Result

The request fails with an invalid-altitude error. The numeric parser receives the original string, so surrounding whitespace prevents a valid finite number from being parsed.

## Expected Result

Numeric input with surrounding whitespace should be accepted and converted as the same numeric value. Blank, non-numeric, NaN, infinite, and overflowing values must remain invalid.

## Focused Check

```bash
go test ./internal/altitude -run TestParseValueAcceptsSurroundingWhitespace -count=1
```

On the `base_bug_002` baseline, the whitespace-surrounded numeric value is rejected.
