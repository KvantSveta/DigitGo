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

	const CLK int = 18
	const DIO int = 25

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

	for *runService {
		// date = 2017-10-02 17:39:22.2612555 +0300 MSK m=+0.003000200 0 2 1 0
		date := time.Now()

		day := date.Day()
		// d0 = 0
		d0 := int(day / 10)
		// d1 = 2
		d1 := int(day % 10)

		month := date.Month()
		// d2 = 1
		d2 := int(month / 10)
		// d3 = 0
		d3 := int(month % 10)

		tm1637.show([4]int{d0, d1, d2, d3})
		tm1637.showDoublePoint(true)

		time.Sleep(time.Second * 5)
	}

	tm1637.setBrightness(BrightDarkest)
	tm1637.clear()
}
