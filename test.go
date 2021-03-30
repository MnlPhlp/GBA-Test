package main

import (
	"GBA-Test/gbaLib"
	"image/color"
	"machine"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

var (
	w, h, x, y, r, oldX, oldY, frameCount int16
	dsp                                   = machine.Display
	background                            = color.RGBA{G: 125}
	foreground                            = color.RGBA{B: 255}
	xSpeed, ySpeed, groundLevel           int16
)

func init() {
	dsp.Configure()
	w, h = dsp.Size()
	x = w / 2
	y = h / 2
	r = 10
	tinydraw.FilledRectangle(dsp, 0, 0, w, h, background)
	oldX = x
	oldY = y
	xSpeed = 0
	ySpeed = 0
	groundLevel = h - 11
}

func onKeyPress() {
	if gbaLib.Buttons.Right.IsPressed() {
		xSpeed = 2
	}
	if gbaLib.Buttons.Left.IsPressed() {
		xSpeed = -2
	}
	if gbaLib.Buttons.A.IsPressed() && y == groundLevel {
		ySpeed = -10
	}
}

func update() {
	// if frameCount < 5 {
	// 	frameCount++
	// 	return
	// }
	// frameCount = 0
	// clear last frame
	tinydraw.Circle(
		dsp,
		oldX,
		oldY,
		10,
		background,
	)
	oldX = x
	oldY = y
	// draw new
	tinydraw.Circle(
		dsp,
		x,
		y,
		10,
		foreground,
	)
	dsp.Display()
	move()
}

func move() {
	newX := x + xSpeed
	if newX < w-11 && newX > 11 {
		x = newX
	}
	newY := y + ySpeed
	if newY < h-11 && newY > 11 {
		y = newY
	} else {
		y = groundLevel
	}
	if xSpeed > 0 {
		xSpeed--
	}
	if xSpeed < 0 {
		xSpeed++
	}
	if ySpeed > 5 {
		ySpeed--
	}
	if ySpeed < 5 {
		ySpeed++
	}
}

func main() {
	gbaLib.SetKeypadInterrupt(onKeyPress)
	gbaLib.SetVBlankInterrupt(update)
	for !gbaLib.Buttons.Start.IsPressed() {
	}
	tinyfont.WriteLine(dsp, &tinyfont.Picopixel, 20, h>>1, "Bye Bye !!", foreground)
}
