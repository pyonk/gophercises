package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	score := 0
	file, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	problems, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, problem_and_answer := range problems {
		problem, answer := problem_and_answer[0], problem_and_answer[1]
		fmt.Printf("Problem #%d: %v = ", i+1, problem)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSuffix(input, "\n")
		if input == string(answer) {
			score = score + 1
		}
	}
	fmt.Printf("You scored %d out of %d", score, len(problems))
}
