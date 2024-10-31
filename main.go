package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func quit() {
	rl.CloseWindow()
}

func main() {
	for running {
		input()
		update()
		draw()
	}
	quit()
}
