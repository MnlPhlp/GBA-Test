package main

import (
	"image/color"
	"machine"

	"github.com/MnlPhlp/gbaLib/pkg/buttons"
	"github.com/MnlPhlp/gbaLib/pkg/interrupts"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

const (
	w      = 240
	h      = 160
	r      = 5
	bottom = h - (r + 3)
	xMax   = w - r
	xMin   = r
)

var (
	x, y, oldX, oldY, frameCount int16
	dsp                          = machine.Display
	background                   = color.RGBA{G: 125}
	foreground                   = color.RGBA{B: 255}
	xSpeed, ySpeed               int16
	jumping                      bool
	floors                       = getFloors(0)
)

func init() {
	dsp.Configure()
	x = w / 2
	y = h / 2
	oldX = x
	oldY = y
	xSpeed = 0
	ySpeed = 0
	drawBackground()
}

func drawBackground() {
	// draw background
	tinydraw.FilledRectangle(dsp, 0, 0, w, h, background)
	y1 := int16(bottom + r + 1)
	y2 := int16(bottom + r + 2)
	for xTmp := int16(0); xTmp < 240; xTmp++ {
		dsp.SetPixel(xTmp, y1, color.RGBA{})
		dsp.SetPixel(xTmp, y2, color.RGBA{})
	}
}

func getFloors(x int16) [][]uint8 {
	floors := make([][]uint8, 2)
	floor := make([]uint8, 240)
	for i := 0; i < 50; i++ {
		floor[i] = h - 100
	}
	for i := 80; i < 130; i++ {
		floor[i] = h - 70
	}
	for i := 160; i < 240; i++ {
		floor[i] = h - 50
	}
	floors[0] = floor
	floors[1] = make([]uint8, 240)
	for i := 50; i < 150; i++ {
		floors[1][i] = 50
	}
	return floors
}

func drawFloor() {
	// draw floor
	for _, floor := range floors {
		for x := int16(0); x < w; x++ {
			if floor[x] == 0 {
				continue
			}
			for i := int16(1); i <= 3; i++ {
				dsp.SetPixel(x, int16(floor[x])+r+i, color.RGBA{})
			}
		}
	}
}

func CheckKeyPress() {
	if buttons.Right.IsPressed() {
		xSpeed = 3
	}
	if buttons.Left.IsPressed() {
		xSpeed = -3
	}
	if buttons.A.IsPressed() && !jumping {
		ySpeed = -12
		jumping = true
	}
}

func update() {
	buttons.Poll()
	CheckKeyPress()
	tinydraw.FilledCircle(
		dsp,
		oldX,
		oldY,
		r,
		background,
	)
	oldX = x
	oldY = y
	// draw new
	tinydraw.FilledCircle(
		dsp,
		x,
		y,
		r,
		foreground,
	)
	drawFloor()
	dsp.Display()
	move()
}

func move() {
	// update position
	newX := x + xSpeed
	if newX < xMax && newX > xMin {
		x = newX
	}
	newY := y + ySpeed
	if newY > bottom {
		jumping = false
		ySpeed = 0
		y = bottom
	} else if newY > r {
		landing := false
		var floor int16
		if ySpeed > 0 { // only check for floors when falling
			for i := 0; i < len(floors); i++ {
				if uint8(y) <= floors[i][x] && uint8(newY) >= floors[i][x] {
					//if floor is in move fall to floor
					landing = true
					floor = int16(floors[i][x])
					break
				}
			}

		}
		if landing {
			ySpeed = 0
			y = floor
			jumping = false
		} else {
			y = newY
			jumping = true
		}
	} else { // stop movement at top
		ySpeed = 0
	}
	// update speed
	if jumping {
		if ySpeed > 12 {
			ySpeed--
		}
		if ySpeed < 12 {
			ySpeed++
		}
	} else {
		if xSpeed > 0 {
			xSpeed--
		}
		if xSpeed < 0 {
			xSpeed++
		}
	}
}

func main() {
	interrupts.SetVBlankInterrupt(update)
	for !buttons.Start.IsPressed() {
	}
	tinyfont.WriteLine(dsp, &tinyfont.Picopixel, 20, h>>1, "Bye Bye !!", foreground)
	interrupts.Stop()
}
