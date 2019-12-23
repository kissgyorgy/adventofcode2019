package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/kissgyorgy/adventofcode2019/intcode"
)

const (
	arcadeGameFile = "day13-input.txt"
	playForFree    = 2
	blackCell      = tcell.Style(tcell.ColorBlack)
)

type tileId int

const (
	empty  tileId = 0
	wall   tileId = 1
	block  tileId = 2
	paddle tileId = 3
	ball   tileId = 4
)

const (
	paddleStay  = 0
	paddleLeft  = -1
	paddleRight = 1
)

var tileIdMap = map[tileId]rune{
	empty:  ' ',
	wall:   '#',
	block:  '▧',
	paddle: '🤽',
	ball:   '🏐',
}

func initScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		msg := fmt.Sprintln("Cannot initialize screen:", err)
		panic(msg)
	}
	if err = s.Init(); err != nil {
		msg := fmt.Sprintln("Cannot initialize screen:", err)
		panic(msg)
	}
	return s
}

func pollEvents(s tcell.Screen, keys chan<- rune, quit chan<- bool) {
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				close(quit)
				return
			case tcell.KeyCtrlL:
				s.Sync()
			case tcell.KeyLeft, tcell.KeyRight, tcell.KeyUp:
				keys <- ev.Key()
			}
		case *tcell.EventResize:
			s.Sync()
		}
	}
}

func runGame(s tcell.Screen, inputs chan<- int, outputs <-chan int, keys <-chan rune, quit <-chan bool) {
	ok := true
	var x, y int
	var tid tileId
	var k tcell.Key

	for {
		select {
		case x, ok = <-outputs:
			if !ok {
				return
			}
			y = <-outputs
			tid = tileId(<-outputs)
			s.SetContent(x, y+1, tileIdMap[tid], nil, blackCell)

			if x == -1 && y == 0 {
				s.SetContent(0, 0, '1', nil, blackCell)
			}
			s.Sync()

		case k = <-keyPress:
			switch k {
			case tcell.KeyLeft:
				inputs <- paddleLeft
			case tcell.KeyRight:
				inputs <- paddleRight
			case tcell.KeyUp:
				inputs <- paddleStay
			}

		case <-quit:
			return
		}
	}
}

func main() {
	program := intcode.Load(arcadeGameFile)
	inputs, outputs := make(chan int, 1), make(chan int)
	program[0] = playForFree
	go intcode.Run("Arcade", program, inputs, outputs)

	s := initScreen()
	quit := make(chan bool)
	keyPress := make(chan tcell.Key)

	go pollEvents(s, keyPress, quit)

	runGame(s, inputs, outputs, keyPress, quit)
	s.PollEvent()
}
