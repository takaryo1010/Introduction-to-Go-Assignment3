package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func (g *Game) readSentence() {
	var questions []string
	f, err := os.Open("Sentences.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		ss := strings.Split(s.Text(), ",")

		questions = append(questions, ss[0])
	}
	g.questions = g.shuffle(questions)

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
}
func (g *Game) shuffle(questions []string) []string {
	t := time.Now().UnixNano()
	rand.Seed(t)
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
	return questions
}

func (g *Game) AddResult(new float64) []float64 {
	f, err := os.Open("Results.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
	defer f.Close()
	var results []float64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err == io.EOF {
			break
		}
		if scanner.Text() != "\n" {
			value, err := strconv.ParseFloat(strings.TrimRight(scanner.Text(), "\n"), 64)
			results = append(results, value)
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー：", err)
				os.Exit(1)
			}

		}

	}
	results = append(results, new)
	sort.Float64s(results)

	file, err := os.Create("Results.csv")
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
	for i := 0; i < 10; i++ {

		file.WriteString(strconv.FormatFloat(results[i], 'f', 2, 64) + "\n")
	}
	return results
}
