package main

import (
	"fmt"
	term "github.com/nsf/termbox-go" // 使用 termbox-go 处理终端界面
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	width  = 20
	height = 10
)

type point struct {
	x, y int
}

var (
	snake []point
	food  point
	dir   = "right"
	score = 0
)

func clearScreen() {
	cmd := exec.Command("clear") // for Linux and MacOS
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func setup() {
	snake = []point{{3, 1}, {2, 1}, {1, 1}}
	food = point{5, 5}
	dir = "right"
	score = 0
	rand.Seed(time.Now().Unix())
	term.Init()
}

func draw() {
	clearScreen()
	fmt.Printf("Score: %d\n", score)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			isSnake := false
			for _, s := range snake {
				if s.x == x && s.y == y {
					fmt.Print("■")
					isSnake = true
					break
				}
			}
			if !isSnake {
				if x == food.x && y == food.y {
					fmt.Print("★")
				} else {
					fmt.Print("  ")
				}
			}
		}
		fmt.Println()
	}
}

func getInput() {
	term.SetInputMode(term.InputEsc)
	ev := term.PollEvent()
	if ev.Type == term.EventKey {
		switch ev.Key {
		case term.KeyArrowUp:
			if dir != "down" {
				dir = "up"
			}
		case term.KeyArrowDown:
			if dir != "up" {
				dir = "down"
			}
		case term.KeyArrowLeft:
			if dir != "right" {
				dir = "left"
			}
		case term.KeyArrowRight:
			if dir != "left" {
				dir = "right"
			}
		case term.KeyEsc:
			term.Close()
			os.Exit(0)
		}
	}
}

func update() bool {
	head := snake[0]
	var newHead point

	switch dir {
	case "up":
		newHead = point{head.x, head.y - 1}
	case "down":
		newHead = point{head.x, head.y + 1}
	case "left":
		newHead = point{head.x - 1, head.y}
	case "right":
		newHead = point{head.x + 1, head.y}
	}

	if newHead.x < 0 || newHead.x >= width || newHead.y < 0 || newHead.y >= height {
		return false // 游戏结束
	}

	for _, s := range snake {
		if newHead.x == s.x && newHead.y == s.y {
			return false // 游戏结束
		}
	}

	snake = append([]point{newHead}, snake...)
	if newHead.x == food.x && newHead.y == food.y {
		score++
		placeFood()
	} else {
		snake = snake[:len(snake)-1]
	}

	return true
}

func placeFood() {
	food = point{rand.Intn(width), rand.Intn(height)}
}

func main() {
	setup()
	defer term.Close()

	for {
		draw()
		getInput()
		if !update() {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Game Over. Your score:", score)
}
