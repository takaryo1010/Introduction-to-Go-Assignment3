package main

import (
	"image/color"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    30,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
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

var (
	flag_start      bool
	flag_main       bool
	flag_num        bool = true
	counter         int  = 0
	point           int  = 0
	input_string    string
	time_count      int = 0
	point_clear     int = 10
	start_time      time.Time
	mplusNormalFont font.Face
	frame           int = 0
)

func (g *Game) Update() error {
	g.runes = ebiten.AppendInputChars(g.runes[:0])
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) {
		flag_start = true
	}
	if repeatingKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if flag_start {
		if time_count == 0 {
			time_count = g.counter
		}
		switch {
		case time_count == g.counter:
			g.text = "3"
			countdown_audio()
		case time_count+60 == g.counter:
			g.text = "2"
		case time_count+120 == g.counter:
			g.text = "1"
		case time_count+180 == g.counter:
			start_time = time.Now()
			flag_main = true
			g.readSentence()

		}

	}
	if flag_main {
		if len(g.runes) != 0 && frame > 2 {
			typing_audio()
			frame = 0
		}
		g.text += string(g.runes)
		input_string += string(g.runes)
		if flag_num && len(g.questions) != 0 {
			flag_num = g.Draw_question(counter)
		}

		if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter) {
			counter, point, flag_num = g.Check_input_text(counter, point)
		}
		if point == point_clear {
			flag_main = false
			g.Result()
			finish_audio()
			flag_start = false
			time_count = 0
			point_clear = 10
			point = 0
			counter = 0

		}
		if repeatingKeyPressed(ebiten.KeyBackspace) && flag_main == true {
			if len(input_string) >= 1 {
				g.text = g.text[:len(g.text)-1]
				input_string = input_string[:len(input_string)-1]
			}
		}

	}

	g.counter++
	frame++
	return nil
}

func (g *Game) Result() {
	result := time.Since(start_time).Seconds()
	results := g.AddResult(result)
	g.text = "Finish!" + "\n" + "time:" + strconv.FormatFloat(result, 'f', 2, 64) + "\n" + "\n" + "Ranking!" + "\n"
	for i := 0; i < 9; i++ {
		if result == results[i] {
			g.text = g.text + strconv.Itoa(i+1) + " : " + strconv.FormatFloat(results[i], 'f', 2, 64) + "<-your time" + "\n"
		}
		if result != results[i] {
			g.text = g.text + strconv.Itoa(i+1) + " : " + strconv.FormatFloat(results[i], 'f', 2, 64) + "\n"
		}

	}
	g.text += "\n" + "Enter : play again  Esc : quit game"
}

func (g *Game) Draw(screen *ebiten.Image) {
	t := g.text
	if g.counter%60 < 30 && flag_main {
		t += "_"
	}
	text.Draw(screen, t, mplusNormalFont, 20, 30, color.White)
	if flag_main {
		text.Draw(screen, "TIME : "+strconv.FormatFloat(time.Since(start_time).Seconds(), 'f', 2, 64), mplusNormalFont, 400, 30, color.White)
	}
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
		input_string = ""
	}
	if flag != true {
		boo_audio()
	}
	return counter, point, flag
}
func (g *Game) Draw_question(i int) bool {
	if len(g.questions) >= i {
		s := ""
		switch i {
		case 0:
			s = "[▪---------]"
		case 1:
			s = "[▪▪--------]"
		case 2:
			s = "[▪▪▪-------]"
		case 3:
			s = "[▪▪▪▪------]"
		case 4:
			s = "[▪▪▪▪▪-----]"
		case 5:
			s = "[▪▪▪▪▪▪----]"
		case 6:
			s = "[▪▪▪▪▪▪▪--=]"
		case 7:
			s = "[▪▪▪▪▪▪▪▪--]"
		case 8:
			s = "[▪▪▪▪▪▪▪▪▪-]"
		case 9:
			s = "[▪▪▪▪▪▪▪▪▪▪]"
		}
		g.text = s + "(" + strconv.Itoa(i+1) + "/10)" + "\n" + string(g.questions[i]) + "\n"
	}
	return false
}
