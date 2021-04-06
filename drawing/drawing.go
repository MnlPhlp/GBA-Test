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
	gbaDraw.Display.FilledDiamond(xOld, yOld, constants.R, background)
	gbaDraw.Display.FilledDiamond(x, y, constants.R, foreground)
}

func drawFloor() {
	var x1, x2, x1Old, x2Old, y int16
	for _, floor := range state.Floors {
		x1 = floor.X1 - state.XBg
		x2 = floor.X2 - state.XBg
		x1Old = floor.X1 - xBgOld
		x2Old = floor.X2 - xBgOld
		y = floor.Y + constants.R + 1
		// remove old lines
		if !(x2Old < 0 || x1Old > constants.W) {
			if x1Old < x1 {
				drawFloorLine(x1Old, x1, y, background)
			}
			if x2Old > x2 {
				drawFloorLine(x2, x2Old, y, background)
			}
		}
		if x2 < 0 || x1 > constants.W {
			continue
		}
		drawFloorLine(x1, x2, y, floorColor)
	}
}

func drawFloorLine(x1, x2, y int16, c gbaDraw.ColorIndex) {
	// limit x Values
	if x1 < 0 {
		x1 = 0
	}
	if x2 >= constants.W {
		x2 = constants.W - 1
	}
	gbaDraw.Display.HLine(x1, x2, y, c)
	gbaDraw.Display.HLine(x1, x2, y+1, c)
	gbaDraw.Display.HLine(x1, x2, y+2, c)
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
