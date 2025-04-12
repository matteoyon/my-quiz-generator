package quiz

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

func RunQuiz(questions []Question, quizLength int) {
	reader := bufio.NewReader(os.Stdin)
	score := 0

	instanceQuestions := []Question{}

	for range quizLength {
		questionIndex := rand.Intn(len(questions))
		instanceQuestions = append(instanceQuestions, questions[questionIndex])
		questions = append(questions[:questionIndex], questions[questionIndex+1:]...)
	}

	for i, q := range instanceQuestions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)

		if q.Code != "" {
			fmt.Println("\nCode:")
			fmt.Println(q.Code)
		}

		fmt.Println("\nOptions:")
		for _, opt := range q.Options {
			fmt.Println(opt)
		}

		fmt.Print("Your answer (split multiple choices with comma ','): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		myAnswers := strings.Split(input, ",")
		sort.Slice(myAnswers, func(x, y int) bool {
			return myAnswers[x] < myAnswers[y]
		})

		questionAnswers := q.Answer
		sort.Slice(questionAnswers, func(x, y int) bool {
			return questionAnswers[x] < questionAnswers[y]
		})

		for _, q := range questionAnswers {
			q = strings.TrimSpace(q)
		}

		correct := true
		if len(myAnswers) == len(questionAnswers) {
			for j, ans := range questionAnswers {
				if !strings.EqualFold(myAnswers[j], ans) {
					correct = false
					break
				}
			}
		} else {
			correct = false
		}

		if correct {
			fmt.Println("✅ Correct!")
			score++
		} else {
			fmt.Println("❌ Wrong! The right answer is:", questionAnswers)
		}

		fmt.Println("Explanation:", q.Explanation)
	}

	fmt.Printf("\nYour final score: %d/%d\n", score, quizLength)
}
