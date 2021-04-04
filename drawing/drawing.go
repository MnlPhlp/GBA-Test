package drawing

import (
	"GBA-Test/constants"
	"GBA-Test/state"

	"github.com/MnlPhlp/gbaLib/pkg/gbaDraw"
)

const (
	background = gbaDraw.Green
	foreground = gbaDraw.Blue
	floorColor = gbaDraw.Black
)

var (
	xBgOld  = state.XBg
	xOld    = state.X
	yOld    = state.Y
	xBgOld2 = state.XBg
	xOld2   = state.X
	yOld2   = state.Y
)

func Configure() {
	gbaDraw.Display.Configure()
	drawBackground()
	gbaDraw.Display.Display()
	drawBackground()
}

func drawBackground() {
	// draw background
	gbaDraw.Display.Filled2PointRect(0, 0, constants.W, constants.H, background)
	y1 := int16(constants.Bottom + constants.R + 1)
	y2 := int16(constants.Bottom + constants.R + 2)
	for xTmp := int16(0); xTmp < 240; xTmp++ {
		gbaDraw.Display.SetPixel(xTmp, y1, floorColor)
		gbaDraw.Display.SetPixel(xTmp, y2, floorColor)
	}
}

func drawFigure(x, y, xOld, yOld int16) {
	const r = constants.R
	gbaDraw.Display.FilledCircle(xOld, yOld, constants.R, background)
	gbaDraw.Display.FilledCircle(x, y, constants.R, foreground)
}

func drawFloor() {
	// delete old Floor
	for _, floor := range state.Floors {
		for x := int16(0); x < constants.W; x++ {
			y := floor[x+state.XBg]
			yOld := floor[x+xBgOld]
			keep := y == yOld || yOld == 0
			for i := int16(1); i <= 3; i++ {
				if !keep {
					// delete old Floor
					gbaDraw.Display.SetPixel(x, int16(yOld)+constants.R+i, background)
				}
				if y != 0 {
					gbaDraw.Display.SetPixel(x, y+constants.R+i, floorColor)
				}
			}
		}
	}
}

func Update() {
	drawFigure(state.X, state.Y, xOld, yOld)
	drawFloor()
	xOld = xOld2
	yOld = yOld2
	xBgOld = xBgOld2
	xOld2 = state.X
	yOld2 = state.Y
	xBgOld2 = state.XBg
	gbaDraw.Display.Display()
}
