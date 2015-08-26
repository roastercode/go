/*

BeagleBone Led Brightness - GoBot - beaglebone_led_brightness.go

*/

package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/beaglebone"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()
	beagleboneAdaptor := beaglebone.NewBeagleboneAdaptor("beaglebone")
	led := gpio.NewLedDriver(beagleboneAdaptor, "led", "p9_16")

	work := func() {
		brightness := uint8(0)
		fadeAmount := uint8(5)

		gobot.Every(100*time.Millisecond, func() {
			led.Brightness(brightness)
			brightness = brightness + fadeAmount
			if brightness == 0 || brightness == 255 {
				fadeAmount = -fadeAmount
			}
		})
	}

	robot := gobot.NewRobot("pwmBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{led},
		work)

	gbot.AddRobot(robot)

	gbot.Start()
}
