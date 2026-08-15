# Bug Reproduction

## Baseline

This reproduction is based on branch `base_bug_005`.

## Steps

1. Create a CSV input containing only a valid header, for example `value,unit`.
2. Read that file with a valid target unit.

## Actual Result

The reader accepts the file and returns an empty result set. The CSV contains no conversion data, but is treated as a successful batch operation.

## Expected Result

A CSV input must contain a header and at least one data row. Header-only input should return an error, while valid header and row parsing behavior remains unchanged.

## Focused Check

```bash
go test ./cmd/altitude-converter -run TestReadFileRejectsCSVWithoutDataRows -count=1
```

On the `base_bug_005` baseline, a header-only CSV unexpectedly succeeds.
