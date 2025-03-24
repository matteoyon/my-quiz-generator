package main

import (
	"fmt"
	"log"
	"ocp-quiz-maker/quiz"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <questions.json> <number_of_questions>")
		return
	}

	filename := os.Args[1]
	questions, err := quiz.LoadQuestions(filename)
	if err != nil {
		log.Fatalf("Failed to load questions: %v", err)
	}

	quizLength, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Failed to parse quiz length: %v", err)
	}

	quiz.RunQuiz(questions, quizLength)
}
