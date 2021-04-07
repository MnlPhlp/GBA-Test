package floors

import (
	"fmt"
)

type Floor struct {
	X1, X2, Y int16
}

type FloorList []Floor

func NewFloors() FloorList { return make([]Floor, 0) }

// func GetFloors() FloorList {
// 	floors := newFloors()
// 	floors = floors.AddFloor(Floor{
// 		X1: 0,
// 		X2: 50,
// 		Y:  constants.H - 100,
// 	})
// 	floors = floors.AddFloor(Floor{
// 		X1: 50,
// 		X2: 120,
// 		Y:  constants.H - 30,
// 	})
// 	floors = floors.AddFloor(Floor{
// 		X1: 70,
// 		X2: 100,
// 		Y:  constants.H - 50,
// 	})
// 	floors = floors.AddFloor(Floor{
// 		X1: 120,
// 		X2: 200,
// 		Y:  constants.H - 120,
// 	})
// 	return floors
// }

func (floors FloorList) AddFloor(f Floor) FloorList {
	if f.X1 < 0 {
		panic(fmt.Errorf("invalid x1 Value. Must be bigger than 0"))
	}
	if f.X2 > WBg {
		panic(fmt.Errorf("invalid x2 Value. Must be smaller than %v", WBg))
	}
	if f.Y < 0 {
		panic(fmt.Errorf("invalid y Value. Must be bigger than 0"))
	}
	if f.Y > HBg {
		panic(fmt.Errorf("invalid y Value. Must be smaller than %v", HBg))
	}
	return append(floors, f)
}
