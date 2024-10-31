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
	playerSpeed                   = 300
	playerRunning         float32 = 1
	dt                    float32
	totalFrames           = 0
)
