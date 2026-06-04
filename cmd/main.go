package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("assets/problems.csv")

	if err != nil {
		fmt.Println("Error Occured ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err != nil {
			break
		}

		fmt.Println(record)
	}

}
