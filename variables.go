package main

import rl "github.com/gen2brain/raylib-go/raylib"

const ()

var (
	running      bool  = true
	screenHeight int32 = 720
	screenWidth  int32 = 1280
	bkgcolour          = rl.SkyBlue

	// player variables
	playerSprite          rl.Texture2D
	playerSrc, playerDest rl.Rectangle
	playerDir                     = 0
	playerSpeed                   = 7
	playerRunning         float32 = 1
)
