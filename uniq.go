package main

import (
	"fmt"

	"github.com/DeveloperMan313/bashnya-go-hw-2/unilines"
)

func main() {
	options := unilines.Options{Count: true, Duplicate: false, Unique: false, Num_fields: 0, Num_chars: 0, IgnoreRegister: false}
	lines, err := unilines.UniqueLines("file.txt", &options)
	if err != nil {
		if err.Error() == "of Count, Duplicate, Unique only one can be true" {
			panic("\nusage\nuniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		}
		panic(err)
	}
	for _, line := range *lines {
		if options.Count {
			fmt.Printf("%d ", line.Count)
		}
		fmt.Println(line.Str)
	}
}
