package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func update() {
	running = !rl.WindowShouldClose()

	dt = rl.GetFrameTime()

	// player walking
	if playerDir == -1 || playerDir == 1 {
		playerDest.X += ((float32(playerDir) * float32(playerSpeed)) * playerRunning) * dt
		fmt.Println("moving")
	}

	playerDir = 0
	playerRunning = 1
	totalFrames++
}
