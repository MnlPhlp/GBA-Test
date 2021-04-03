package main

import (
	"github.com/MnlPhlp/gbaLib/pkg/buttons"
	"github.com/MnlPhlp/gbaLib/pkg/gbaDraw"
)

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
