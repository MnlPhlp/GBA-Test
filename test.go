package main

import (
	"image/color"

	"github.com/MnlPhlp/gbaLib/pkg/buttons"
	"github.com/MnlPhlp/gbaLib/pkg/gbaDraw"
	"github.com/MnlPhlp/gbaLib/pkg/interrupts"
)

const (
	w         = 240
	h         = 160
	wBG       = 300
	hBG       = 160
	xBgMax    = wBG - w
	r         = 5
	bottom    = h - (r + 3)
	xMax      = w - r - 1
	xMin      = r
	moveRange = 50 + r
	xMaxSpeed = 3
)

var (
	x, y, oldX, oldY, frameCount int16
	xBG, yBG, xBgOld             int16
	dsp                          = gbaDraw.Display
	background                   = gbaDraw.ToRgba15(color.RGBA{G: 125})
	foreground                   = gbaDraw.ToRgba15(color.RGBA{G: 125})
	floorColor                   = uint16(0)
	xSpeed, ySpeed               int16
	jumping                      bool
	floors                       = getFloors()
)

func init() {
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
	gbaDraw.Filled2PointRect(0, 0, w, h, background)
	y1 := int16(bottom + r + 1)
	y2 := int16(bottom + r + 2)
	for xTmp := int16(0); xTmp < 240; xTmp++ {
		dsp.SetPixel(xTmp, y1, 0)
		dsp.SetPixel(xTmp, y2, 0)
	}
}

func getFloors() [][]uint8 {
	floors := make([][]uint8, 2)
	floor := make([]uint8, wBG)
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
	floors[1] = make([]uint8, wBG)
	for i := 50; i < 150; i++ {
		floors[1][i] = 50
	}
	return floors
}

func drawFloor() {
	// delete old Floor
	for _, floor := range floors {
		for x := int16(0); x < w; x++ {
			y := floor[x+xBG]
			yOld := floor[x+xBgOld]
			keep := y == yOld || yOld == 0
			for i := int16(1); i <= 3; i++ {
				if !keep {
					// delete old Floor
					dsp.SetPixel(x, int16(yOld)+r+i, background)
				}
				if y != 0 {
					dsp.SetPixel(x, int16(y)+r+i, floorColor)
				}
			}
		}
	}
	xBgOld = xBG
}

func CheckKeyPress() {
	if buttons.Right.IsPressed() {
		xSpeed = xMaxSpeed
	}
	if buttons.Left.IsPressed() {
		xSpeed = -xMaxSpeed
	}
	if !jumping {
		if buttons.A.IsPressed() {
			ySpeed = -12
			jumping = true
		} else if buttons.B.IsPressed() && y < bottom {
			y++ // drop trough the floor
		}
	}
}

func drawFigure(xOld, yOld, x, y int16) {
	gbaDraw.Filled2PointRect(x-r, y-r, x+r, y+r, foreground)
}

func update() {
	buttons.Poll()
	CheckKeyPress()
	move()
	dsp.Blank()
	drawFigure(oldX, oldY, x, y)
	oldX = x
	oldY = y
	drawFloor()
	dsp.Display()
}

func move() {
	// update position
	newX := x + xSpeed
	// update screen pos
	if x < moveRange && xBG > 0 {
		xBG -= xMaxSpeed
		newX += xMaxSpeed
		if xBG < 0 {
			newX += xBG
			xBG = 0
		}
	}
	if x > w-moveRange && xBG < xBgMax {
		xBG += xMaxSpeed
		newX -= xMaxSpeed
		if xBG > xBgMax {
			newX += xBG - xBgMax
			xBG = xBgMax
		}
	}
	if newX < xMin {
		x = xMin
	} else if newX > xMax {
		x = xMax
	} else {
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
				xFloor := x + xBG
				if uint8(y) <= floors[i][xFloor] && uint8(newY) >= floors[i][xFloor] {
					//if floor is in move fall to floor
					landing = true
					floor = int16(floors[i][xFloor])
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
	dsp.Configure()
	interrupts.SetVBlankInterrupt(update)
	for !buttons.Start.IsPressed() {
	}
	interrupts.Stop()
}
