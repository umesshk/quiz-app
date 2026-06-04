package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Problem struct {
	question string
	answer   string
}

func main() {

	csv_file_name := flag.String("csv", "problems.csv", "Enter the name of a CSV file contaning [questions,answers]")

	flag.Parse()
	file, err := os.Open("assests/" + *csv_file_name)

	if err != nil {
		fmt.Println("Error Occured ", err)

		exit(fmt.Sprintf("Error Openeing file %v", err))
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Error Reading Records %v", err))
	}

	problem_set := ParseProblems(records)

}

func ParseProblems(records [][]string) []Problem {

	problem_set := make([]Problem, len(records))

	for i, rec := range records {

		problem_set[i] = Problem{
			question: rec[0],
			answer:   rec[1],
		}
	}
	return problem_set

}

func exit(msg string) {
	log.Fatal(msg)
}
