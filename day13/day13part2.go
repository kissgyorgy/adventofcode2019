package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
	"github.com/kissgyorgy/adventofcode2019/intcode"
)

const (
	arcadeGameFile = "day13-input.txt"
	playForFree    = 2
	blackCell      = tcell.Style(tcell.ColorBlack)
	animationSpeed = 50 * time.Millisecond
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

func pollEvents(s tcell.Screen, keys chan<- tcell.Key, quit chan<- bool) {
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

func getDigits(num int) string {
	digits := ""
	for num > 0 {
		digit := num % 10
		char := strconv.Itoa(digit)
		digits = char + digits
		num /= 10
	}
	return digits
}

func runGame(s tcell.Screen, inputs chan<- int, outputs <-chan int, keyPress <-chan tcell.Key, quit <-chan bool) {
	var ok bool
	var x, y, obj int
	var paddleX int
	var tid tileId

	scoreEnd := len("Score: ")

	for x, c := range "Score: " {
		s.SetContent(x, 0, c, nil, blackCell)
	}

	for {
		select {
		case x, ok = <-outputs:
			if !ok {
				return
			}
			y = <-outputs
			obj = <-outputs

			tid = tileId(obj)
			s.SetContent(x, y+1, tileIdMap[tid], nil, blackCell)

			if x == -1 && y == 0 {
				for i, digit := range getDigits(obj) {
					s.SetContent(scoreEnd+i, 0, digit, nil, blackCell)
				}
			}
			s.Sync()

			switch tid {
			case paddle:
				paddleX = x
			case ball:
				if x < paddleX {
					inputs <- paddleLeft
				} else if x > paddleX {
					inputs <- paddleRight
				} else {
					inputs <- paddleStay
				}
				time.Sleep(animationSpeed)
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
	fmt.Println("Press any key to quit")
	s.PollEvent()
	s.Fini()
}
