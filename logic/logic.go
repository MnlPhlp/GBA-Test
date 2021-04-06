package logic

import (
	"GBA-Test/constants"
	"GBA-Test/state"

	"github.com/MnlPhlp/gbaLib/pkg/buttons"
)

var (
	xSpeed, ySpeed int16
	jumping        bool
)

func CheckKeyPress() {
	if buttons.Right.IsPressed() {
		xSpeed = constants.XMaxSpeed
	}
	if buttons.Left.IsPressed() {
		xSpeed = -constants.XMaxSpeed
	}
	if !jumping {
		if buttons.A.IsPressed() {
			ySpeed = -12
			jumping = true
		} else if buttons.B.IsPressed() && state.Y < constants.Bottom {
			state.Y++ // drop trough the floor
		}
	}
}

func updateXPos() {
	newX := state.X + xSpeed
	// update screen pos
	if state.X < constants.MoveRange && state.XBg > 0 {
		state.XBg -= constants.XMaxSpeed
		newX += constants.XMaxSpeed
		if state.XBg < 0 {
			newX += state.XBg
			state.XBg = 0
		}
	}
	if state.X > constants.W-constants.MoveRange && state.XBg < constants.XBgMax {
		state.XBg += constants.XMaxSpeed
		newX -= constants.XMaxSpeed
		if state.XBg > constants.WBg {
			newX += state.XBg - constants.WBg
			state.XBg = constants.WBg
		}
	}
	if newX < constants.XMin {
		state.X = constants.XMin
	} else if newX > constants.XMax {
		state.X = constants.XMax
	} else {
		state.X = newX
	}
}

func updateSpeeds() {
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

func updateYPos() {
	newY := state.Y + ySpeed
	if newY > constants.Bottom {
		jumping = false
		ySpeed = 0
		state.Y = constants.Bottom
	} else if newY > constants.R {
		landing := false
		var floorHeight int16
		if ySpeed > 0 { // only check for state.Floors when falling
			x := state.X + state.XBg
			for _, floor := range state.Floors {
				if floor.X2 < x || floor.X1 > x {
					continue
				}
				if state.Y <= floor.Y && newY >= floor.Y {
					//if floor is in move fall to floor
					landing = true
					floorHeight = floor.Y
					break
				}
			}
		}
		if landing {
			ySpeed = 0
			state.Y = floorHeight
			jumping = false
		} else {
			state.Y = newY
			jumping = true
		}
	} else { // stop movement at top
		ySpeed = 0
	}
}

func Move() {
	buttons.Poll()
	CheckKeyPress()
	// update position
	updateXPos()
	updateYPos()
	// update speed
	updateSpeeds()
}
