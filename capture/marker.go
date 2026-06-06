package capture

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

// fontData holds 5x7 pixel patterns for basic numbers and symbols
var fontData = map[rune][5]byte{
	'0': {0x3E, 0x51, 0x49, 0x45, 0x3E},
	'1': {0x00, 0x42, 0x7F, 0x40, 0x00},
	'2': {0x42, 0x61, 0x51, 0x49, 0x46},
	'3': {0x21, 0x41, 0x45, 0x4B, 0x31},
	'4': {0x18, 0x14, 0x12, 0x7F, 0x10},
	'5': {0x27, 0x45, 0x45, 0x45, 0x39},
	'6': {0x3C, 0x4A, 0x49, 0x49, 0x30},
	'7': {0x01, 0x71, 0x09, 0x05, 0x03},
	'8': {0x36, 0x49, 0x49, 0x49, 0x36},
	'9': {0x06, 0x49, 0x49, 0x29, 0x1E},
	'd': {0x38, 0x44, 0x44, 0x3F, 0x40},
	'p': {0x7F, 0x09, 0x09, 0x09, 0x06},
	'x': {0x41, 0x22, 0x1C, 0x22, 0x41},
	'=': {0x14, 0x14, 0x14, 0x14, 0x14},
	'L': {0x7F, 0x40, 0x40, 0x40, 0x40},
	'R': {0x7F, 0x09, 0x19, 0x29, 0x46},
	'T': {0x01, 0x01, 0x7F, 0x01, 0x01},
	'B': {0x7F, 0x49, 0x49, 0x49, 0x36},
	':': {0x00, 0x36, 0x36, 0x00, 0x00},
	' ': {0x00, 0x00, 0x00, 0x00, 0x00},
}

func drawChar(img *image.RGBA, char rune, startX, startY int, clr color.Color) {
	pattern, ok := fontData[char]
	if !ok {
		pattern = fontData[' ']
	}
	bounds := img.Bounds()
	for col := 0; col < 5; col++ {
		b := pattern[col]
		for row := 0; row < 7; row++ {
			if (b & (1 << row)) != 0 {
				px := startX + col
				py := startY + row
				if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
					img.Set(px, py, clr)
				}
			}
		}
	}
}

func drawString(img *image.RGBA, text string, startX, startY int, clr color.Color) {
	for i, char := range text {
		drawChar(img, char, startX+i*6, startY, clr)
	}
}

// DrawClickMarker draws the standard target ring on click position (kept for backward compatibility)
func DrawClickMarker(inputPath string, outputPath string, clickX, clickY int) error {
	return DrawMeasurementMarker(inputPath, outputPath, clickX, clickY, 0, 0)
}

// DrawMeasurementMarker draws click markers, guidelines to screen edges, and step-to-step distance arrow lines.
func DrawMeasurementMarker(inputPath string, outputPath string, clickX, clickY, prevX, prevY int) error {
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

	redColor := color.RGBA{R: 255, G: 0, B: 0, A: 220}
	redSoftColor := color.RGBA{R: 255, G: 0, B: 0, A: 60}
	guideColor := color.RGBA{R: 255, G: 255, B: 255, A: 100} // White translucent guide
	stepLineColor := color.RGBA{R: 46, G: 204, B: 113, A: 220} // Emerald Green

	// 1. Draw dashed crosshair guidelines from click coordinates to screen bounds
	drawDashedLine(markedImg, bounds.Min.X, clickY, bounds.Max.X, clickY, guideColor)
	drawDashedLine(markedImg, clickX, bounds.Min.Y, clickX, bounds.Max.Y, guideColor)

	// 2. Draw measurement distance texts from click to bounds
	drawString(markedImg, fmt.Sprintf("L:%dpx", clickX), bounds.Min.X+10, clickY-12, guideColor)
	drawString(markedImg, fmt.Sprintf("R:%dpx", bounds.Max.X-clickX), bounds.Max.X-80, clickY-12, guideColor)
	drawString(markedImg, fmt.Sprintf("T:%dpx", clickY), clickX+10, bounds.Min.Y+10, guideColor)
	drawString(markedImg, fmt.Sprintf("B:%dpx", bounds.Max.Y-clickY), clickX+10, bounds.Max.Y-20, guideColor)

	// 3. Draw step-to-step distance line if there's a valid previous step coordinate
	if prevX > 0 && prevY > 0 && (prevX != clickX || prevY != clickY) {
		drawLine(markedImg, prevX, prevY, clickX, clickY, stepLineColor)
		// Calculate Euclidean distance
		dx := float64(clickX - prevX)
		dy := float64(clickY - prevY)
		distance := math.Sqrt(dx*dx + dy*dy)
		
		// Draw label at middle of the line
		midX := (prevX + clickX) / 2
		midY := (prevY + clickY) / 2
		drawString(markedImg, fmt.Sprintf("d=%dpx", int(distance)), midX+10, midY-6, stepLineColor)
		
		// Draw small anchor ring at previous point
		drawRing(markedImg, prevX, prevY, 6, stepLineColor)
	}

	// 4. Draw main target rings at current click position
	radiusOuter := 30
	radiusInner := 12
	drawRing(markedImg, clickX, clickY, radiusOuter, redColor)
	drawRing(markedImg, clickX, clickY, radiusInner, redColor)
	
	// Blend soft red inside inner ring
	for y := clickY - radiusInner; y <= clickY+radiusInner; y++ {
		for x := clickX - radiusInner; x <= clickX+radiusInner; x++ {
			dx := float64(x - clickX)
			dy := float64(y - clickY)
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < float64(radiusInner) {
				if x >= bounds.Min.X && x < bounds.Max.X && y >= bounds.Min.Y && y < bounds.Max.Y {
					original := markedImg.At(x, y)
					markedImg.Set(x, y, blendColors(original, redSoftColor))
				}
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

func drawRing(img *image.RGBA, cx, cy, radius int, clr color.Color) {
	bounds := img.Bounds()
	for y := cy - radius - 1; y <= cy+radius+1; y++ {
		for x := cx - radius - 1; x <= cx+radius+1; x++ {
			dx := float64(x - cx)
			dy := float64(y - cy)
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist >= float64(radius-1) && dist <= float64(radius+1) {
				if x >= bounds.Min.X && x < bounds.Max.X && y >= bounds.Min.Y && y < bounds.Max.Y {
					img.Set(x, y, clr)
				}
			}
		}
	}
}

func drawDashedLine(img *image.RGBA, x1, y1, x2, y2 int, clr color.Color) {
	bounds := img.Bounds()
	dx := math.Abs(float64(x2 - x1))
	dy := math.Abs(float64(y2 - y1))
	sx, sy := 1, 1
	if x1 > x2 {
		sx = -1
	}
	if y1 > y2 {
		sy = -1
	}
	err := dx - dy

	x, y := x1, y1
	count := 0
	for {
		// Draw only on even 6-pixel chunks to make it dashed
		if (count/6)%2 == 0 {
			if x >= bounds.Min.X && x < bounds.Max.X && y >= bounds.Min.Y && y < bounds.Max.Y {
				img.Set(x, y, clr)
			}
		}
		count++

		if x == x2 && y == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x += sx
		}
		if e2 < dx {
			err += dx
			y += sy
		}
	}
}

func drawLine(img *image.RGBA, x1, y1, x2, y2 int, clr color.Color) {
	bounds := img.Bounds()
	dx := math.Abs(float64(x2 - x1))
	dy := math.Abs(float64(y2 - y1))
	sx, sy := 1, 1
	if x1 > x2 {
		sx = -1
	}
	if y1 > y2 {
		sy = -1
	}
	err := dx - dy

	x, y := x1, y1
	for {
		if x >= bounds.Min.X && x < bounds.Max.X && y >= bounds.Min.Y && y < bounds.Max.Y {
			img.Set(x, y, clr)
		}
		if x == x2 && y == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x += sx
		}
		if e2 < dx {
			err += dx
			y += sy
		}
	}
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
