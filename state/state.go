package state

import "GBA-Test/floors"

var (
	X, Y           int16
	XBg, YBg       int16
	XSpeed, YSpeed int16
	Jumping        bool
	Floors         floors.FloorList
)
