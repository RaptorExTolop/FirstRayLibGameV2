package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func input() {
	if rl.IsKeyDown(rl.KeyW) {
		fmt.Println("jump")
	}
	if rl.IsKeyDown(rl.KeyD) {
		fmt.Println("right")
		playerDir += 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		fmt.Println("left")
		playerDir -= 1
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		playerRunning = 1.5
	}
}
