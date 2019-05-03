package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename) //flag package requires passing a pointer to a string

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s \n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll() // reads all of the lines of the csv file
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)
	correct := 0 // initializing the correct answer score
	// Now that we have all the problems, need to iterate over them and ask for answers

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer) // careful using this when there is a possibility of multi word string answers
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("you scored %d out of %d.\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem { // takes a 2d string slice, returns a single problem slice
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
