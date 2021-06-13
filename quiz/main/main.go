package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func loadProblems() ([][]string, error) {
	file, err := os.Open("./problems.csv")
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	problems, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func askForProblem(ch chan string) {
	var input string;
	fmt.Scanf("%s\n", &input)
	ch <- strings.TrimSuffix(input, "\n")
}

func displayScore(score int, total int) {
	fmt.Printf("You scored %d out of %d", score, total)
}

func main() {
	score := 0
	problems, err := loadProblems()
	if err != nil {
		log.Fatal(err)
	}
	loop:
		for i, problemAndAnswer := range problems {
			problem, answer := problemAndAnswer[0], problemAndAnswer[1]
			fmt.Printf("Problem #%d: %v = ", i+1, problem)

			ch := make(chan string)
			go askForProblem(ch)

			select {
			case usersAnswer := <- ch:
				if usersAnswer == string(answer) {
					score = score + 1
				}
			case <- time.After(2 * time.Second):
				println("")
				break loop
			}
		}
	displayScore(score, len(problems))
}
