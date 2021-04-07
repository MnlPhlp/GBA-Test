package sound

import "github.com/MnlPhlp/gbaLib/pkg/gbaSound"

func Enable() {
	gbaSound.Enable()
	gbaSound.Sound1.Enable()
}

func Jump() {
	gbaSound.Sound1.SetFrequency(300)
	gbaSound.Sound1.SetLength(1000)
	gbaSound.Sound1.SetSweep(false, 5, 7)
	gbaSound.Sound1.Play()
}

func GameOver() {
	gbaSound.Sound1.SetFrequency(300)
	gbaSound.Sound1.SetLength(1000)
	gbaSound.Sound1.SetSweep(true, 5, 7)
	gbaSound.Sound1.Play()
}

func Disable() {
	gbaSound.Disable()
}
