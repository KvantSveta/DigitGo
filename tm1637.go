package main

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

const AutoAddr int = 0x40  // 64
//const FixedAddr int = 0x44 // 68
const StartAddr int = 0xC0 // 192

const BrightDarkest int = 0
//const BrightTypical int = 2
const BrightHighest int = 7

var digitToSegment = []int{
	63,  // 0b0111111, // 0
	6,   // 0b0000110, // 1
	91,  // 0b1011011, // 2
	79,  // 0b1001111, // 3
	102, // 0b1100110, // 4
	109, // 0b1101101, // 5
	125, // 0b1111101, // 6
	7,   // 0b0000111, // 7
	127, // 0b1111111, // 8
	111, // 0b1101111, // 9
	119, // 0b1110111, // A
	124, // 0b1111100, // b
	57,  // 0b0111001, // C
	94,  // 0b1011110, // d
	121, // 0b1111001, // E
	113, // 0b1110001, // F
	99,  // 0b1100011, // degree
	64,  // 0b1000000, // minus
	0,   // 0b0000000, // nothing
}

type TM1637 struct {
	clkPin rpio.Pin
	dataPin rpio.Pin
	brightness int
	doublePoint bool
	currentData [4]int
}

func (tm *TM1637) clear() {
	tm.doublePoint = false
	tm.show([4]int{18, 18, 18, 18})
}

func (tm *TM1637) show(data [4]int) {
	tm.currentData = data
	tm.start()
	tm.writeByte(AutoAddr)
	tm.stop()
	tm.start()
	tm.writeByte(StartAddr)

	for _, v := range data {
		if tm.doublePoint {
			tm.writeByte(digitToSegment[v] + 0x80)
		} else {
			tm.writeByte(digitToSegment[v])
		}
	}

	tm.stop()
	tm.start()
	tm.writeByte(0x88 + tm.brightness)
	tm.stop()
}

func (tm *TM1637) setBrightness(brightness int) {
	if brightness >= BrightHighest {
		brightness = BrightHighest
	} else if brightness < BrightDarkest {
		brightness = BrightDarkest
	}

	if tm.brightness != brightness {
		tm.brightness = brightness
		tm.show(tm.currentData)
	}
}

func (tm *TM1637) showDoublePoint(on bool) {
	tm.doublePoint = on

	tm.show(tm.currentData)
}

func intToBool(d int) bool {
	// 0 = 0 & 0x01; 2 & 0x01;
	// 1 = 1 & 0x01; 3 & 0x01;
	r := d & 0x01

	if r == 1 {
		return true
	} else {
		return false
	}
}

func (tm *TM1637) writeByte(data int) {
	for i := 0; i < 8; i++ {
		tm.clkPin.Low()
		if intToBool(data) {
			tm.dataPin.High()
		} else {
			tm.dataPin.Low()
		}
		data = data >> 1
		tm.clkPin.High()
	}

	tm.clkPin.Low()
	tm.dataPin.High()
	tm.clkPin.High()
	tm.dataPin.Input()

	for tm.dataPin.Read() == 1 {
		time.Sleep(time.Millisecond)
		if tm.dataPin.Read() == 1 {
			tm.dataPin.Output()
			tm.dataPin.Low()
			tm.dataPin.Input()
		}
	}

	tm.dataPin.Output()
	tm.dataPin.Low()
}

func (tm *TM1637) start() {
	tm.clkPin.High()
	tm.dataPin.High()
	tm.dataPin.Low()
	tm.clkPin.Low()
}

func (tm *TM1637) stop() {
	tm.clkPin.Low()
	tm.dataPin.Low()
	tm.clkPin.High()
	tm.dataPin.High()
}
