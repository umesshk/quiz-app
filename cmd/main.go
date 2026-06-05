package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	question string
	answer   string
}

var timer *time.Timer

func main() {

	csv_file_name := flag.String("csv", "problems.csv", "Enter the name of a CSV file contaning [questions,answers]")
	time_limit := flag.Int("limit", 30, "Enter the time limit for each queston(seconds) ")

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

	DisplayQuestios(problem_set, time_limit)

}

func DisplayQuestios(problem_set []Problem, time_limit *int) {

	n := len(problem_set)
	i := 0

	for i < n {
		timer = time.NewTimer(time.Duration(*time_limit) * time.Second)
		fmt.Printf("%s : ", problem_set[i].question)
		inputChannel := make(chan string)

		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			inputChannel <- ans

		}()

		select {
		case <-timer.C:

			exit(fmt.Sprintf("\nTime Over!! \nScore %v/%v", i, n))

		case ans := <-inputChannel:
			if ans != problem_set[i].answer {
				exit(fmt.Sprintf("Wrong Answer \nScore %v/%v", i+1, n))
			}
		}
		i++
	}

	exit(fmt.Sprintf("Total Score :  %v/%v", n, n))

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
	fmt.Println(msg)
	os.Exit(1)
}
