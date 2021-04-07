package main

import (
	"fmt"
	"image/color"

	"github.com/MnlPhlp/gbaLib/pkg/buttons"
	"github.com/MnlPhlp/gbaLib/pkg/gbaDraw"
	"github.com/MnlPhlp/gbaLib/pkg/gbaSound"
	"github.com/MnlPhlp/gbaLib/pkg/registers"
	"tinygo.org/x/tinyfont"
)

var (
	dsp              = gbaDraw.Display
	sound            = gbaSound.Sound1
	yPos             = 10
	frequency uint16 = 100
)

func println(str string) {
	tinyfont.WriteLine(dsp, &tinyfont.TomThumb, 10, int16(yPos), str, color.RGBA{G: 255, B: 255})
	yPos += int(tinyfont.TomThumb.YAdvance)
}

func showRegister() {
	yPos = 10
	dsp.Blank()
	println(fmt.Sprintf("  (REG_SOUND1CNT_L) = %v", registers.Sound.Sound1Cnt_L.Get()))
	println(fmt.Sprintf("  (REG_SOUND1CNT_H) = %v", registers.Sound.Sound1Cnt_H.Get()))
	println(fmt.Sprintf("  (REG_SOUND1CNT_X) = %v", registers.Sound.Sound1Cnt_X.Get()))
	println(fmt.Sprintf("  (Buff_SOUND1CNT_H) = %v", sound.BuffLen.Get()))
	println(fmt.Sprintf("  (Buff_SOUND1CNT_X) = %v", sound.BuffFreq.Get()))
	println(fmt.Sprint("Frequency: "))
	regVal := registers.Sound.Sound1Cnt_X.Get() & 0x7ff
	regVal2 := uint16(2048 - 131072/int(300))
	println(fmt.Sprintf("   Register real: %v", regVal))
	println(fmt.Sprintf("   Register wanted: %v", regVal2))
	hz := 4194304 / (32 * (2048 - int(regVal)))
	println(fmt.Sprintf("   Frequency real: %v", hz))
	println(fmt.Sprintf("   Frequency wanted: %v", frequency))
	// time
	println(fmt.Sprint("Length: "))
	regVal = registers.Sound.Sound1Cnt_H.Get() & 0x3f
	println(fmt.Sprintf("   Register: %v", regVal))
	sec := (64 - float32(regVal)) * (1 / 256)
	println(fmt.Sprintf("   Length real: %.2f sec", sec))
	println(fmt.Sprintf("   Length wanted: %v sec", 1))
	gbaDraw.VSync()
	dsp.Display()
}

func wait() {
	buttons.Poll()
	for !buttons.A.IsPressed() {
		buttons.Poll()
	}
}

func main() {
	dsp.Configure()
	sound.Enable()
	gbaSound.Enable()
	sound.Enable()
	sound.SetLength(1000)
	sound.SetLoop(false)
	for {
		for dutyCyle := 0; dutyCyle < 4; dutyCyle++ {
			sound.SetFrequency(frequency)
			sound.SetWaveDutyCyle(uint16(dutyCyle))
			sound.Flush()
			showRegister()
			wait()
		}
		frequency *= 2
	}
}
