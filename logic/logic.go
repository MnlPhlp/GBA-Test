package logic

import (
	"GBA-Test/constants"
	"GBA-Test/floors"
	"GBA-Test/sound"
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
			ySpeed = -constants.YMaxSpeed
			jumping = true
			sound.Jump()
		} else if buttons.B.IsPressed() && state.Y < constants.Bottom {
			state.Y++ // drop trough the floor
		}
	}
}

func updateBgPos(pos, bgPos, move, size, bgSize int16) (int16, int16) {
	// update screen pos
	if pos < constants.MoveRange && bgPos > 0 {
		bgPos -= move
		pos += move
		if bgPos < 0 {
			pos += bgPos
			bgPos = 0
		}
	} else if pos > size-constants.MoveRange && bgPos < bgSize-size {
		bgPos += move
		pos -= move
		if bgPos > bgSize {
			pos += bgPos - bgSize
			bgPos = bgSize
		}
	}
	return pos, bgPos
}

func updateXPos() {
	newX := state.X + xSpeed
	// update screen pos
	newX, state.XBg = updateBgPos(newX, state.XBg, constants.XMaxSpeed, constants.W, floors.WBg)
	// update figure pos
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
		ySpeed = 0
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
	} else if newY < constants.R && state.YBg == 0 {
		// stop movement at top
		state.Y = constants.R
		ySpeed = 0
	} else {
		landing := false
		var floorHeight int16
		if ySpeed > 0 { // only check for floors when falling
			x := state.X + state.XBg
			for _, floor := range state.Floors {
				floorHeight = floor.Y - state.YBg
				if floor.X2 < x || floor.X1 > x || floorHeight < state.Y || floorHeight > newY {
					continue
				} else {
					//if floor is in move fall to floor
					landing = true
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
		// update screen pos
		move := int16(5)
		if ySpeed > 0 {
			move = ySpeed
		} else if ySpeed < 0 {
			move = -ySpeed
		}
		state.Y, state.YBg = updateBgPos(state.Y, state.YBg, move, constants.H, floors.HBg)
	}
}

func Move() (gameOver bool) {
	buttons.Poll()
	CheckKeyPress()
	// update position
	updateXPos()
	updateYPos()
	// update speed
	updateSpeeds()
	return state.Y >= constants.Bottom
}
