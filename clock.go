package main

import (
	"os"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

func main() {
	runService := new(bool)
	*runService = true

	go stopProgram(runService)

	const CLK int = 23
	const DIO int = 24

	tm1637 := TM1637{
		clkPin: rpio.Pin(CLK),
		dataPin: rpio.Pin(DIO),
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

	doublePoint := false

	for *runService {
		// date = 2017-10-02 17:39:22.2612555 +0300 MSK m=+0.003000200
		date := time.Now()

		// hour = 17
		hour := date.Hour()
		// do = 1
		d0 := int(hour / 10)
		// d1 = 7
		d1 := int(hour % 10)

		// minute = 39
		minute := date.Minute()
		// d2 = 3
		d2 := int(minute / 10)
		// d3 = 9
		d3 := int(minute % 10)

		// second = 22
		second := date.Second()

		if second % 2 == 1 {
			doublePoint = true
		} else {
			doublePoint = false
		}

		tm1637.show([4]int{d0, d1, d2, d3})
		tm1637.showDoublePoint(doublePoint)

		time.Sleep(time.Second)
	}

	tm1637.setBrightness(BrightDarkest)
	tm1637.clear()
}
