package main

import (
	"github.com/MnlPhlp/gbaLib"
	"image/color"
	"machine"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

const (
	w      = 240
	h      = 160
	r      = 10
	bottom = h - (r + 3)
)

var (
	x, y, oldX, oldY, frameCount int16
	dsp                          = machine.Display
	background                   = color.RGBA{G: 125}
	foreground                   = color.RGBA{B: 255}
	xSpeed, ySpeed               int16
	jumping                      bool
)

func init() {
	dsp.Configure()
	x = w / 2
	y = h / 2
	oldX = x
	oldY = y
	xSpeed = 0
	ySpeed = 0
	// draw background
	tinydraw.FilledRectangle(dsp, 0, 0, w, h, background)
	tinydraw.Line(dsp, 0, bottom+r+1, w, bottom+r+1, color.RGBA{})
}

func getGroundLevel(x int16) int16 {
	if x > 160 {
		return h - 50
	}
	if x < 50 {
		return h - 100
	}
	if x > 80 && x < 130 {
		return h - 70
	}
	return bottom
}

func drawFloor() {
	// draw floor
	for x := int16(0); x < w; x++ {
		floor := getGroundLevel(x)
		if floor < bottom {
			dsp.SetPixel(x, floor+r+1, color.RGBA{})
		}
	}
}

func CheckKeyPress() {
	if gbaLib.Buttons.Right.IsPressed() {
		xSpeed = 3
	}
	if gbaLib.Buttons.Left.IsPressed() {
		xSpeed = -3
	}
	if gbaLib.Buttons.A.IsPressed() && !jumping {
		ySpeed = -12
		jumping = true
	}
}

func update() {
	CheckKeyPress()
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
	drawFloor()
	dsp.Display()
	move()
}

func move() {
	// update position
	newX := x + xSpeed
	if newX < w-11 && newX > 11 {
		x = newX
	}
	groundLevel := getGroundLevel(x)
	newY := y + ySpeed
	if newY > 11 {
		if newY < bottom && y > groundLevel || newY < groundLevel {
			y = newY
			jumping = true
		} else if y < groundLevel { // move to bottom
			y = groundLevel
		} else if y > groundLevel && y < bottom {
			y = bottom
		}
	} else { // stop movement at top
		ySpeed = 0
	}
	// update speed
	if (y == groundLevel || y == bottom) && ySpeed > 0 {
		ySpeed = 0
		jumping = false
	} else {
		if ySpeed > 12 {
			ySpeed--
		}
		if ySpeed < 12 {
			ySpeed++
		}
	}
	if !jumping {
		if xSpeed > 0 {
			xSpeed--
		}
		if xSpeed < 0 {
			xSpeed++
		}
	}
}

func main() {
	gbaLib.SetVBlankInterrupt(update)
	for !gbaLib.Buttons.Start.IsPressed() {
	}
	tinyfont.WriteLine(dsp, &tinyfont.Picopixel, 20, h>>1, "Bye Bye !!", foreground)
	gbaLib.Stop()
}
