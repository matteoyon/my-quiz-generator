package main

import (
	"fmt"
	"log"
	"ocp-quiz-maker/quiz"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <number_of_questions> <questions1.json> [<questions2.json> ...]")
		return
	}

	quizLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to parse quiz length: %v", err)
	}

	var allQuestions []quiz.Question

	if os.Args[2] == "-a" {
		if len(os.Args) < 4 {
			log.Fatalf("Please specify a folder after the -a flag")
		}
		allQuestions = loadAllQuestionsFromFolder(os.Args[3])
	} else {
		allQuestions = loadSpecificQuestions(os.Args[2:])
	}

	quiz.RunQuiz(allQuestions, quizLength)
}

func loadSpecificQuestions(files []string) []quiz.Question {
	var allQuestions []quiz.Question
	for _, filename := range files {

		questions, err := quiz.LoadQuestions(filename)
		if err != nil {
			log.Fatalf("Failed to load questions: %s: %v", filename, err)
		}
		allQuestions = append(allQuestions, questions...)
	}
	return allQuestions
}

func loadAllQuestionsFromFolder(folder string) []quiz.Question {
	var allQuestions []quiz.Question

	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatalf("Failed to read folder: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(folder, file.Name())
			questions, err := quiz.LoadQuestions(filePath)
			if err != nil {
				log.Fatalf("Failed to load questions from %s: %v", filePath, err)
			}
			allQuestions = append(allQuestions, questions...)
		}
	}

	return allQuestions
}
