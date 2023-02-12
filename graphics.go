package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

type Game struct {
	runes     []rune
	text      string
	counter   int
	questions []string
}

var flag_start bool
var flag_main bool
var flag_num bool = true
var counter int = 0
var point int = 0
var input_string string
var time_count int = 0

func (g *Game) Update() error {
	// Add runes that are input by the user by AppendInputChars.
	// Note that AppendInputChars result changes every frame, so you need to call this
	// every frame.
	g.runes = ebiten.AppendInputChars(g.runes[:0])
	g.text += string(g.runes)
	input_string += string(g.runes)

	// If the enter key is pressed, add a line break.

	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		flag_start = true
	}
	if flag_start {
		if time_count == 0 {
			time_count = g.counter
		}
		switch {
		case time_count == g.counter:
			g.text = "3"
		case time_count+60 == g.counter:
			g.text = "2"
		case time_count+120 == g.counter:
			g.text = "1"
		case time_count+180 == g.counter:
			flag_main = true

		}

	}

	if flag_main {
		if flag_num {
			flag_num = g.Draw_question(counter)
		}

		if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) {
			counter, point, flag_num = g.Check_input_text(counter, point)
		}
		if len(g.questions) < counter {
			flag_main = false
			g.text = "finish"

		}
		if repeatingKeyPressed(ebiten.KeyBackspace) && flag_main == true {
			if len(input_string) >= 1 {
				g.text = g.text[:len(g.text)-1]
				input_string = input_string[:len(input_string)-1]
			}
		}

	}

	g.counter++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Blink the cursor.
	t := g.text
	if g.counter%60 < 30 {
		t += "_"
	}
	ebitenutil.DebugPrint(screen, t)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func create_g() *Game {
	g := &Game{
		text:    "Press the Enter key\n",
		counter: 0,
	}
	return g
}
func (g *Game) Check_input_text(counter int, point int) (int, int, bool) {
	var flag bool = false
	if input_string == g.questions[counter] {
		point += 1
		counter += 1
		flag = true
		fmt.Print(counter)
		input_string = ""
	}
	return counter, point, flag
}
func (g *Game) Draw_question(i int) bool {
	g.text = string(g.questions[i]) + "\n"
	return false
}
