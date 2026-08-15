package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"altitude-converter/internal/altitude"
)

type record struct{ Value, From, To, Result string }

func main() {
	var value, from, to, input, output, fields string
	flag.StringVar(&value, "value", "", "altitude value for a single conversion")
	flag.StringVar(&from, "from", "", "source unit for a single conversion")
	flag.StringVar(&to, "to", "", "target unit (required for all modes)")
	flag.StringVar(&input, "input", "", "input .txt or .csv file")
	flag.StringVar(&output, "output", "", "output CSV file")
	flag.StringVar(&fields, "fields", "value,from,to,result", "CSV fields: value,from,to,result")
	flag.Parse()

	if to == "" {
		fail(errors.New("-to is required"))
	}
	if value != "" && input != "" {
		fail(errors.New("use either -value/-from or -input, not both"))
	}
	var rows []record
	var err error
	if input != "" {
		rows, err = readFile(input, to)
	} else {
		rows, err = single(value, from, to)
	}
	if err != nil {
		fail(err)
	}
	for _, row := range rows {
		fmt.Printf("%s %s = %s %s\n", row.Value, row.From, row.Result, row.To)
	}
	if output != "" {
		if err := writeCSV(output, fields, rows); err != nil {
			fail(err)
		}
	}
}

func single(value, from, to string) ([]record, error) {
	if value == "" || from == "" {
		return nil, errors.New("single conversion requires -value and -from")
	}
	return convert(value, from, to)
}

func readFile(path, to string) ([]record, error) {
	if _, err := altitude.NormalizeUnit(to); err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input: %w", err)
	}
	defer f.Close()
	if strings.EqualFold(filepath.Ext(path), ".csv") {
		return readCSV(f, to)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}
	var rows []record
	for lineNo, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.FieldsFunc(line, func(r rune) bool { return r == ',' || r == ';' || r == '\t' || r == ' ' })
		if len(parts) != 2 {
			return nil, fmt.Errorf("line %d: expected value and unit", lineNo+1)
		}
		converted, err := convert(parts[0], parts[1], to)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNo+1, err)
		}
		rows = append(rows, converted...)
	}
	return rows, nil
}

func readCSV(r io.Reader, to string) ([]record, error) {
	rows, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read CSV: %w", err)
	}
	if len(rows) < 1 {
		return nil, errors.New("CSV must contain a header and at least one data row")
	}
	valueIndex, unitIndex := -1, -1
	for i, h := range rows[0] {
		switch strings.ToLower(strings.TrimSpace(h)) {
		case "value", "altitude":
			valueIndex = i
		case "unit", "from":
			unitIndex = i
		}
	}
	if valueIndex < 0 || unitIndex < 0 {
		return nil, errors.New("CSV header must include value (or altitude) and unit (or from)")
	}
	var out []record
	for i, row := range rows[1:] {
		if valueIndex >= len(row) || unitIndex >= len(row) {
			return nil, fmt.Errorf("CSV row %d: missing value or unit", i+2)
		}
		converted, err := convert(row[valueIndex], row[unitIndex], to)
		if err != nil {
			return nil, fmt.Errorf("CSV row %d: %w", i+2, err)
		}
		out = append(out, converted...)
	}
	return out, nil
}

func convert(value, from, to string) ([]record, error) {
	v, err := altitude.ParseValue(value)
	if err != nil {
		return nil, err
	}
	r, canonicalFrom, canonicalTo, err := altitude.Convert(v, from, to)
	if err != nil {
		return nil, err
	}
	return []record{{altitude.Format(v), canonicalFrom, canonicalTo, altitude.Format(r)}}, nil
}

func writeCSV(path, requested string, rows []record) error {
	valid := map[string]bool{"value": true, "from": true, "to": true, "result": true}
	fields := strings.Split(requested, ",")
	for i := range fields {
		fields[i] = strings.ToLower(strings.TrimSpace(fields[i]))
		if !valid[fields[i]] {
			return fmt.Errorf("unsupported CSV field %q", fields[i])
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output: %w", err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	if err := w.Write(fields); err != nil {
		return err
	}
	for _, row := range rows {
		values := map[string]string{"value": row.Value, "from": row.From, "to": row.To, "result": row.Result}
		record := make([]string, len(fields))
		for i, field := range fields {
			record[i] = values[field]
		}
		if err := w.Write(record); err != nil {
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return fmt.Errorf("write output CSV: %w", err)
	}
	return nil
}

func fail(err error) { fmt.Fprintln(os.Stderr, "error:", err); flag.Usage(); os.Exit(1) }
