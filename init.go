package main

import rl "github.com/gen2brain/raylib-go/raylib"

func init() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	playerSprite = rl.LoadTexture("res/mainCharacter.png")
	playerDest = rl.NewRectangle(0, 0, 156, 156)
	playerSrc = rl.NewRectangle(0, 0, 52, 52)
}
