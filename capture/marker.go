package capture

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

// DrawClickMarker は (clickX, clickY) を中心に、視線誘導用の半透明赤丸を描画します
func DrawClickMarker(inputPath string, outputPath string, clickX, clickY int) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	markedImg := image.NewRGBA(bounds)
	draw.Draw(markedImg, bounds, img, image.Point{}, draw.Src)

	redColor := color.RGBA{R: 255, G: 0, B: 0, A: 200}
	redSoftColor := color.RGBA{R: 255, G: 0, B: 0, A: 60}

	radiusOuter := 35
	radiusInner := 15

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dx := float64(x - clickX)
			dy := float64(y - clickY)
			dist := math.Sqrt(dx*dx + dy*dy)

			if (dist >= float64(radiusOuter-2) && dist <= float64(radiusOuter)) || 
			   (dist >= float64(radiusInner-1) && dist <= float64(radiusInner+1)) {
				markedImg.Set(x, y, redColor)
			}
			if dist < float64(radiusInner-1) {
				original := markedImg.At(x, y)
				blended := blendColors(original, redSoftColor)
				markedImg.Set(x, y, blended)
			}
		}
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, markedImg)
}

func blendColors(bg color.Color, fg color.RGBA) color.Color {
	r1, g1, b1, a1 := bg.RGBA()
	r1Val := uint8(r1 >> 8)
	g1Val := uint8(g1 >> 8)
	b1Val := uint8(b1 >> 8)
	a1Val := uint8(a1 >> 8)

	alpha := float64(fg.A) / 255.0
	r := uint8(float64(fg.R)*alpha + float64(r1Val)*(1.0-alpha))
	g := uint8(float64(fg.G)*alpha + float64(g1Val)*(1.0-alpha))
	b := uint8(float64(fg.B)*alpha + float64(b1Val)*(1.0-alpha))
	  
	return color.RGBA{R: r, G: g, B: b, A: a1Val}
}
