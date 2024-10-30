package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)
}

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
