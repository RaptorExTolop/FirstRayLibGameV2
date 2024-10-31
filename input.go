package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		fmt.Println("jump")
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		fmt.Println("right")
		playerDir += 1
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		fmt.Println("left")
		playerDir -= 1
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		playerRunning = 1.75
	}
}
