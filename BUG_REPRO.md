# Bug Reproduction

## Baseline

This reproduction is based on branch `base_bug_001`.

## Steps

1. Convert one nautical mile to metres with the source unit `nautical mile`.
2. Repeat with a spaced alias containing tabs or other Unicode whitespace, such as `nautical\tmile`.

## Actual Result

The conversion returns an unsupported-unit error. The normalization path removes only leading and trailing whitespace, leaving whitespace inside an otherwise supported alias.

## Expected Result

Supported aliases remain valid when their words are separated by ordinary, tab, or Unicode whitespace. Whitespace-only input must remain unsupported.

## Focused Check

```bash
go test ./internal/altitude -run TestConvertAcceptsSpacedUnitAlias -count=1
```

On the `base_bug_001` baseline, the spaced `nautical mile` alias fails to normalize.
