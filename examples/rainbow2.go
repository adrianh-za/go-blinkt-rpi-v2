package examples

import (
	"go-blinkt-rpi-v2/lib"
	"time"
)

func Rainbow2() {
	colorsys := lib.ColorSys{}
	blinkt := lib.BlinktGpio{}.NewBlinkt(0.2)
	blinkt.ClearOnExit = true
	blinkt.CaptureExit = true
	blinkt.Setup()

	var hue = int64(0)
	var hueStep = int64(3)

	for {

		//Hue must be between 0 and 359
		hue = hue % 360

		//Calculate hue for each pixel on Blinkt
		for pixelCount := 0; pixelCount <= 7; pixelCount++ {
			var offset = pixelCount * 30
			var pixelHue = (hue + int64(offset)) % 360
			var r, g, b = colorsys.Hsv2Rgb(float64(pixelHue), 1.0, 1.0)
			blinkt.SetPixel(pixelCount, int(r), int(g), int(b))

			//Debug stuff
			//fmt.Println("hue: ", hue, " pixelCount: ", pixelCount, " pixelHue", pixelHue, "RGB", r, "|", g, "|", b)
		}

		//Increment hue by step value
		hue = hue + hueStep

		//Show and do a small delay
		blinkt.Show()
		time.Sleep(5 * time.Millisecond)
	}
}