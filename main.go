package main

import (
	"GBA-Test/constants"
	"GBA-Test/drawing"
	"GBA-Test/logic"
	"GBA-Test/state"

	"github.com/MnlPhlp/gbaLib/pkg/buttons"
	"github.com/MnlPhlp/gbaLib/pkg/interrupts"
	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

func init() {
	state.X = constants.W / 2
	state.Y = constants.H / 2
	state.XSpeed = 0
	state.YSpeed = 0
	state.Floors = logic.GetFloors()
}

func main() {
	drawing.Configure()
	interrupts.SetVBlankInterrupt(update)
	registers.Timer.Tm0Cnt.Set(1 << 7)
	for !buttons.Start.IsPressed() {
	}
	interrupts.Stop()
}

var framecount = 0

func update() {
	logic.Move()
	drawing.Update()
}
