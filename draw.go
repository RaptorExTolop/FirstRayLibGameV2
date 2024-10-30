package main

import rl "github.com/gen2brain/raylib-go/raylib"

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgcolour)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

	rl.EndDrawing()
}
