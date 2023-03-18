package util

import (
	"math"

	glm "github.com/go-gl/mathgl/mgl32"
)

type Color struct {
	glm.Vec4

	red   float32
	green float32
	blue  float32
	alpha float32
}

var WHITE Color = NewColorRGBf(1.0, 1.0, 1.0)
var BLACK Color = NewColorRGBf(0.0, 0.0, 0.0)
var DRACULA Color = NewColorRGBi(45, 52, 54)
var MINT_GREEN Color = NewColorRGBi(0, 184, 148)
var PINK_GLAMOUR Color = NewColorRGBi(255, 118, 117)
var ELECTRON_BLUE Color = NewColorRGBi(9, 132, 227)

func NewColorRGBf(red, green, blue float32) Color {
	color := Color{}

	color.red = red
	color.green = green
	color.blue = blue
	color.Vec4 = glm.Vec4{red, green, blue, 1.0}

	return color
}

func NewColorRGBAf(red, green, blue, alpha float32) Color {
	color := Color{}

	color.red = red
	color.green = green
	color.blue = blue
	color.alpha = alpha
	color.Vec4 = glm.Vec4{red, green, blue, alpha}

	return color
}

func NewColorRGBi(red, green, blue int) Color {
	color := Color{}

	color.red = float32(red) / 255.0
	color.green = float32(green) / 255.0
	color.blue = float32(blue) / 255.0
	color.Vec4 = glm.Vec4{color.red, color.green, color.blue, 1.0}

	return color
}

func NewColorRGBAi(red, green, blue int, alpha float32) Color {
	color := Color{}

	color.red = float32(red) / 255.0
	color.green = float32(green) / 255.0
	color.blue = float32(blue) / 255.0
	color.alpha = alpha
	color.Vec4 = glm.Vec4{color.red, color.green, color.blue, alpha}

	return color
}


func DarkenColor(color Color, amount float32) Color {
  red := color.red * (1.0 - amount)
  rFinal := math.Max(0.0, math.Min(1.0, float64(red)))  

  green := color.green * (1.0 - amount)
  gFinal := math.Max(0.0, math.Min(1.0, float64(green)))  

  blue := color.blue * (1.0 - amount)
  bFinal := math.Max(0.0, math.Min(1.0, float64(blue)))  


  return NewColorRGBAf(float32(rFinal), float32(gFinal), float32(bFinal), color.alpha)


}
