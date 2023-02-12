package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func (g *Game) readSentence() {
	f, err := os.Open("Sentences.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		ss := strings.Split(s.Text(), ",")

		g.questions = append(g.questions, ss[0])
	}
	g.shuffle()
	fmt.Println(g.questions)

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
}
func (g *Game) shuffle() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	rand.Shuffle(len(g.questions), func(i, j int) {
		g.questions[i], g.questions[j] = g.questions[j], g.questions[i]
	})

}
