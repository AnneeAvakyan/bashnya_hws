package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Options struct {
	CountMode    bool // -c подсчитать количество встречаний строки во входных данных
	RepeatedMode bool // -d вывести только те строки, которые повторились во входных данных.
	UniqueMode   bool // -u вывести только те строки, которые не повторились во входных данных.
	IgnoreCase   bool // -i не учитывать регистр букв.
	NumFields    int  // -f не учитывать первые num_fields полей в строке.
	NumChars     int  // -s не учитывать первые num_chars символов в строке
}

func processLine(line string, options Options) string {
	if options.NumFields > 0 {
		fields := strings.Fields(line)
		if options.NumFields < len(fields) {
			line = strings.Join(fields[options.NumFields:], " ")
		} else {
			line = ""
		}
	}

	if options.NumChars > 0 && len(line) > options.NumChars {
		line = line[options.NumChars:]
	}

	if options.IgnoreCase {
		line = strings.ToLower(line)
	}

	return line
}

func Uniq(lines []string, options Options) []string {
	if len(lines) == 0 {
		return []string{}
	}

	counts := make(map[string]int)
	originalLines := make(map[string]string)
	order := []string{}

	for _, line := range lines {
		processed := processLine(line, options)

		if _, exists := originalLines[processed]; !exists {
			originalLines[processed] = line
			order = append(order, processed)
		}

		counts[processed]++
	}

	var result []string

	for _, processed := range order {
		count := counts[processed]
		original := originalLines[processed]

		switch {
		case options.CountMode:
			result = append(result, fmt.Sprintf("%d %s", count, original))

		case options.RepeatedMode:
			if count > 1 {
				result = append(result, original)
			}

		case options.UniqueMode:
			if count == 1 {
				result = append(result, original)
			}

		default:
			result = append(result, original)
		}
	}

	return result
}

func readLines(reader io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func writeLines(writer io.Writer, lines []string) error {
	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateOptions(options Options) error {
	modes := 0
	if options.CountMode {
		modes++
	}
	if options.RepeatedMode {
		modes++
	}
	if options.UniqueMode {
		modes++
	}

	if modes > 1 {
		return fmt.Errorf("flags -c, -d, -u are mutually exclusive")
	}

	if options.NumFields < 0 {
		return fmt.Errorf("number of fields cannot be negative")
	}
	if options.NumChars < 0 {
		return fmt.Errorf("number of chars cannot be negative")
	}

	return nil
}

func main() {
	var (
		countMode    = flag.Bool("c", false, "count occurrences")
		repeatedMode = flag.Bool("d", false, "only repeated lines")
		uniqueMode   = flag.Bool("u", false, "only unique lines")
		ignoreCase   = flag.Bool("i", false, "ignore case")
		numFields    = flag.Int("f", 0, "number of fields to skip")
		numChars     = flag.Int("s", 0, "number of chars to skip")
	)

	flag.Parse()
	options := Options{
		CountMode:    *countMode,
		RepeatedMode: *repeatedMode,
		UniqueMode:   *uniqueMode,
		IgnoreCase:   *ignoreCase,
		NumFields:    *numFields,
		NumChars:     *numChars,
	}

	if err := validateOptions(options); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var input io.Reader = os.Stdin
	inputFile := flag.Arg(0)
	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	lines, err := readLines(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	result := Uniq(lines, options)

	var output io.Writer = os.Stdout
	outputFile := flag.Arg(1)
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	if err := writeLines(output, result); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}
}
