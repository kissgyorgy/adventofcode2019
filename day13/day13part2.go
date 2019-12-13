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
	defaultSpeed   = 50 * time.Millisecond
	width          = 44
	speedEnd       = width - 3
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
			case tcell.KeyRune:
				keys <- ev.Rune()
			}
		case *tcell.EventResize:
			s.Sync()
		}
	}
}

func getDigits(num int) string {
	if num == 0 {
		return "0"
	}
	digits := ""
	for num > 0 {
		digit := num % 10
		char := strconv.Itoa(digit)
		digits = char + digits
		num /= 10
	}
	return digits
}

func drawSleep(s tcell.Screen, speed time.Duration) {
	// clear it first
	for i := speedEnd; i < width; i++ {
		s.SetContent(i, 0, ' ', nil, blackCell)
	}
	for i, digit := range getDigits(int(speed / time.Millisecond)) {
		s.SetContent(speedEnd+i, 0, digit, nil, blackCell)
	}
}

func runGame(s tcell.Screen, inputs chan<- int, outputs <-chan int, keys <-chan rune, quit <-chan bool) {
	var ok bool
	var x, y, obj int
	var paddleX int
	var tid tileId
	var k rune

	scoreEnd := len("Score: ")
	speed := defaultSpeed

	for x, c := range "Score: " {
		s.SetContent(x, 0, c, nil, blackCell)
	}

	for x, c := range "Sleep: " {
		s.SetContent(speedEnd-(len("Sleep: ")-x), 0, c, nil, blackCell)
	}

	drawSleep(s, speed)

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
			}

			time.Sleep(speed)

		case k = <-keys:
			switch k {
			case '+':
				speed += 10 * time.Millisecond
			case '-':
				speed -= 10 * time.Millisecond
			}
			if speed < 0*time.Millisecond {
				speed = 0 * time.Millisecond
			}
			drawSleep(s, speed)

		case <-quit:
			return
		}
	}
}

func main() {
	program := intcode.Load(arcadeGameFile)
	inputs, outputs := make(chan int, 1), make(chan int)
	program[0] = playForFree
	go intcode.Run("Arcade", program, inputs, outputs, intcode.Silent)

	s := initScreen()
	quit := make(chan bool)
	keys := make(chan rune)

	go pollEvents(s, keys, quit)

	runGame(s, inputs, outputs, keys, quit)
	fmt.Println("Press any key to quit")
	s.PollEvent()
	s.Fini()
}
