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

func GetFloors() [][]int16 {
	floors := make([][]int16, 2)
	floor := make([]int16, constants.WBg)
	for i := 0; i < 50; i++ {
		floor[i] = constants.H - 100
	}
	for i := 80; i < 130; i++ {
		floor[i] = constants.H - 70
	}
	for i := 160; i < 240; i++ {
		floor[i] = constants.H - 50
	}
	floors[0] = floor
	floors[1] = make([]int16, constants.WBg)
	for i := 50; i < 150; i++ {
		floors[1][i] = 50
	}
	return floors
}

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
		var floor int16
		if ySpeed > 0 { // only check for state.Floors when falling
			for i := 0; i < len(state.Floors); i++ {
				xFloor := state.X + state.XBg
				if state.Y <= state.Floors[i][xFloor] && newY >= state.Floors[i][xFloor] {
					//if floor is in move fall to floor
					landing = true
					floor = int16(state.Floors[i][xFloor])
					break
				}
			}

		}
		if landing {
			ySpeed = 0
			state.Y = floor
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
