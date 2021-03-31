package gbaLib

import (
	"machine"
	"runtime/interrupt"
)

type InterruptHandler struct {
	f   func()
	num int
}

var interrupts = make(map[interrupt.Interrupt]InterruptHandler)

func isr(i interrupt.Interrupt) {
	if handler, ok := interrupts[i]; ok {
		handler.f()
	}
}

func setupInterrupt(i interrupt.Interrupt, num int, f func()) {
	interrupts[i] = InterruptHandler{
		f:   f,
		num: num,
	}
	i.Enable()
}

func SetKeypadInterrupt(f func()) {
	// enable the interrupt
	Register.KeyCnt.SetBits(1 << 0xE)
	// enable all keys
	Register.KeyCnt.SetBits(0b1111111111)
	// create a new Interrupt and store the function
	i := interrupt.New(machine.IRQ_KEYPAD, isr)
	setupInterrupt(i, machine.IRQ_KEYPAD, f)
}

func SetVBlankInterrupt(f func()) {
	// enable the interrupt
	Register.DispStat.SetBits(1 << 3)
	// create a new Interrupt and store the function
	i := interrupt.New(machine.IRQ_VBLANK, isr)
	setupInterrupt(i, machine.IRQ_VBLANK, f)
}

func Finish() {
	// disable interrupts
	Register.IE.Set(0)
	// keep running
	for {
	}
}
