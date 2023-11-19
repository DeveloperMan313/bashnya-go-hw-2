package unilines

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Options struct {
	Count, Duplicate, Unique bool
	Num_fields, Num_chars    int
	IgnoreRegister           bool
}

type Line struct {
	Count int
	Str   string
}

func UniqueLines(filename string, options *Options) (*[]Line, error) {
	cdu := 0
	for _, flag := range []*bool{&options.Count, &options.Duplicate, &options.Unique} {
		if *flag {
			cdu++
		}
	}
	if cdu > 1 {
		return nil, errors.New("of Count, Duplicate, Unique only one can be true")
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	linesMap := map[string]Line{}
	duplicateCnt := 0
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")[options.Num_fields:]
		line := strings.Join(fields, " ")[options.Num_chars:]
		key := line
		if options.IgnoreRegister {
			key = strings.ToLower(line)
		}
		mapVal, valExists := linesMap[key]
		if !valExists {
			mapVal = Line{0, line}
		}
		if mapVal.Count == 1 {
			duplicateCnt++
		}
		mapVal.Count++
		linesMap[key] = mapVal
	}
	arrLen := len(linesMap)
	if options.Duplicate {
		arrLen = duplicateCnt
	} else if options.Unique {
		arrLen -= duplicateCnt
	}
	lines := make([]Line, 0, arrLen)
	for _, line := range linesMap {
		if !(options.Duplicate || options.Unique) || (options.Duplicate && line.Count > 1) || (options.Unique && line.Count == 1) {
			lines = append(lines, line)
		}
	}
	return &lines, nil
}
