package main

import rl "github.com/gen2brain/raylib-go/raylib"

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgcolour)

	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(0, 0), 0, rl.RayWhite)

	rl.EndDrawing()
}
