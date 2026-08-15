# Bug Reproduction

## Baseline

This reproduction is based on branch `base_bug_003`.

## Steps

1. Call `Convert` with a finite value near the maximum float, such as `math.MaxFloat64`.
2. Convert from kilometres to metres.

## Actual Result

The multiplication overflows and the conversion returns an infinite result with no error. Callers can receive a successful record whose converted altitude is not finite.

## Expected Result

Conversions must return only finite results. A calculation that overflows or produces NaN must return an error, while valid finite conversions, including same-unit conversion, must retain their normal result.

## Focused Check

```bash
go test ./internal/altitude -run TestConvertOverflow -count=1
```

On the `base_bug_003` baseline, an overflowing conversion unexpectedly succeeds.
