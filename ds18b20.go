package main

import (
	"os/exec"
	"log"
	"os"
	"strings"
	"strconv"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)



func main() {
	runService := new(bool)
	*runService = true

	go stopProgram(runService)

	const CLK int = 6
	const DIO int = 5

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
		// out = "ae 01 4b 46 7f ff 0c 10 76 : crc=76 YES\nae 01 4b 46 7f ff 0c 10 76 t=26875\n"
		out, err := exec.Command("cat", "/sys/bus/w1/devices/28-05170143ccff/w1_slave").Output()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// a = {"ae 01 4b 46 7f ff 0c 10 76 : crc=76 YES", "ae 01 4b 46 7f ff 0c 10 76 t=26875"}
		a := strings.Split(string(out), "\n")

		// b = "ae 01 4b 46 7f ff 0c 10 76 t=26875"
		b := strings.Split(a[1], "=")

		// c = "26875"
		c := b[len(b) - 1]

		// d = 26875
		d, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// e = 26.875
		e := float64(d) / 1000

		// g = 27
		g := int(e + 0.5)

		// d0 = 2
		d0 := int(g / 10)

		// d1 = 7
		d1 := int(g % 10)

		tm1637.show([4]int{d0, d1, 16, 12})

		time.Sleep(time.Second * 5)
	}

	tm1637.setBrightness(BrightDarkest)
	tm1637.clear()
}
