/*

A PWM example by @Drahoslav7, using the go-rpio library

Toggles a LED on physical pin 35 (GPIO pin 19)
Connect a LED with resistor from pin 35 to ground.

*/

package main

import (
	"os"
	"time"

	rpio "github.com/s-fairchild/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	pin := rpio.Pin(19)
	pin.Mode(rpio.Pwm)
	pin.Freq(64000)
	pin.DutyCycle(0, 32)
	// the LED will be blinking at 2000Hz
	// (source frequency divided by cycle length => 64000/32 = 2000)

	// five times smoothly fade in and out
	for i := 0; i < 5; i++ {
		for i := uint32(0); i < 32; i++ { // increasing brightness
			pin.DutyCycle(i, 32)
			time.Sleep(time.Second / 32)
		}
		for i := uint32(32); i > 0; i-- { // decreasing brightness
			pin.DutyCycle(i, 32)
			time.Sleep(time.Second / 32)
		}
	}
}
