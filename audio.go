package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

var (
	audioContext *audio.Context
	typing       *audio.Player
	countdown    *audio.Player
	finish       *audio.Player
	boo          *audio.Player
)

func init() {
	var err error
	audioContext = audio.NewContext(44100)

	f, err := os.Open("SE/typing1.mp3")
	if err != nil {
		log.Fatal(err)
	}

	d, err := mp3.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}

	typing, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	typing.SetVolume(0.10)
	f, err = os.Open("SE/countdown.mp3")
	if err != nil {
		log.Fatal(err)
	}

	d, err = mp3.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}

	countdown, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	countdown.SetVolume(0.10)

	f, err = os.Open("SE/finish.mp3")
	if err != nil {
		log.Fatal(err)
	}

	d, err = mp3.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}

	finish, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	finish.SetVolume(0.50)

	f, err = os.Open("SE/boo.mp3")
	if err != nil {
		log.Fatal(err)
	}

	d, err = mp3.Decode(audioContext, f)
	if err != nil {
		log.Fatal(err)
	}

	boo, err = audio.NewPlayer(audioContext, d)
	if err != nil {
		log.Fatal(err)
	}

	boo.SetVolume(0.50)
}
func countdown_audio() {
	countdown.Rewind()
	countdown.Play()
}
func typing_audio() {
	typing.Rewind()
	typing.Play()
}
func finish_audio() {
	finish.Rewind()
	finish.Play()
}
func boo_audio() {
	boo.Rewind()
	boo.Play()
}
