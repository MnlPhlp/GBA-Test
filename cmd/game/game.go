package main

import (
	"GBA-Test/drawing"
	"GBA-Test/floors"
	"GBA-Test/logic"
	"GBA-Test/state"
	"image/color"

	"github.com/MnlPhlp/gbaLib/pkg/buttons"
	"github.com/MnlPhlp/gbaLib/pkg/gbaBios"
	"github.com/MnlPhlp/gbaLib/pkg/gbaDraw"
	"github.com/MnlPhlp/gbaLib/pkg/interrupts"
	"tinygo.org/x/tinyfont"
)

func init() {
	state.X = floors.StartX
	state.Y = floors.StartY
	state.XBg = floors.StartXBg
	state.YBg = floors.StartYBg
	state.XSpeed = 0
	state.YSpeed = 0
	state.Floors = floors.Floors
}

func main() {
	drawing.Configure()
	for !buttons.Start.IsPressed() {
		update()
	}
	finish()
}

func update() {
	// calculate next frame
	logic.Move()
	drawing.Update()
	// wait for Display refresh
	gbaDraw.VSync()
	// Display new frame
	gbaDraw.Display.Display()
}

func finish() {
	gbaDraw.Display.Filled2PointRect(0, 0, 240, 160, gbaDraw.ToColorIndex(color.RGBA{R: 150, G: 150, B: 200}))
	tinyfont.WriteLine(gbaDraw.Display, &tinyfont.TomThumb, 50, 50, "Bye Bye !!", color.RGBA{})
	gbaDraw.VSync()
	gbaDraw.Display.Display()
	interrupts.Disable()
	gbaBios.Stop()
}
