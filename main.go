package main

import (
	"github.com/stianeikeland/go-rpio"
	"fmt"
	"os"
	"time"
)

func main() {
	const ClockPin int = 6
	const DataPin int = 5

	tm1637 := TM1637{
		clkPin: rpio.Pin(ClockPin),
		dataPin: rpio.Pin(DataPin),
		brightness: BrightDarkest,
		doublePoint: false,
		currentData: [4]int{18, 18, 18, 18},
	}

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	tm1637.clkPin.Output()
	tm1637.dataPin.Output()

	tm1637.clkPin.Low()
	tm1637.dataPin.Low()

	for i := 0; i < 18; i++ {
		tm1637.show([4]int{i, i, i, i})
		if (i % 2) == 1 {
			tm1637.showDoublePoint(true)
		} else {
			tm1637.showDoublePoint(false)
		}

		time.Sleep(time.Second)
	}

	tm1637.setBrightness(BrightDarkest)
	tm1637.clear()
}
