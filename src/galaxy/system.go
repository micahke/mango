package galaxy

import (
	"math/rand"

	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/util"
)

var (
	SYSTEM_GENERATION_THRESHOLD int32 = 0
)

type System struct {
	systemType SystemClass
	exists     bool

	size   float32
	offset glm.Vec2
	color  util.Color
}

type SystemClass int

const (
	SYSTEM_CLASS_O SystemClass = 0
	SYSTEM_CLASS_B SystemClass = 1
	SYSTEM_CLASS_A SystemClass = 2
	SYSTEM_CLASS_F SystemClass = 3
	SYSTEM_CLASS_G SystemClass = 4
	SYSTEM_CLASS_K SystemClass = 5
	SYSTEM_CLASS_M SystemClass = 6
)

var (
	CLASS_O_COLOR util.Color = util.NewColorRGBi(0, 194, 209)
	CLASS_B_COLOR util.Color = util.NewColorRGBi(214, 252, 255)
	CLASS_A_COLOR util.Color = util.NewColorRGBi(240, 231, 216)
	CLASS_F_COLOR util.Color = util.NewColorRGBi(225, 194, 56)
	CLASS_G_COLOR util.Color = util.NewColorRGBi(255, 224, 86)
	CLASS_K_COLOR util.Color = util.NewColorRGBi(255, 145, 0)
	CLASS_M_COLOR util.Color = util.NewColorRGBi(255, 49, 46)
)

func NewSystem(xCoord, yCoord int64, fullGeneration bool) *System {
	system := new(System)

	perlinValue := PerlinValueAtCoords(xCoord, yCoord, true)

	if perlinValue < 0.5 {
		system.exists = false
		return system
	}

	seed := (xCoord&0xFFFF)<<16 | (yCoord & 0xFFFF)
	rand.Seed(seed)

	if rand.Int31n(20) > SYSTEM_GENERATION_THRESHOLD {
		system.exists = false
		return system
	}

	system.exists = true
	// system.size = (float32(rand.Int63n(50)) + 30.0) / 100.0

	system.Classify()
	system.GenerateSize()

	if !fullGeneration {
		return system
	}

	return system

}

// Updates a system's Stellar Classification information
func (system *System) Classify() {
	// generate a number between 0 and 1
	odds := rand.Float64()

	if odds <= 0.005 {
		system.systemType = SYSTEM_CLASS_O
		system.color = CLASS_O_COLOR
	} else if odds <= 0.015 {
		system.systemType = SYSTEM_CLASS_B
		system.color = CLASS_B_COLOR
	} else if odds <= 0.045 {
		system.systemType = SYSTEM_CLASS_A
		system.color = CLASS_A_COLOR
	} else if odds <= 0.125 {
		system.systemType = SYSTEM_CLASS_F
		system.color = CLASS_F_COLOR
	} else if odds <= 0.285 {
		system.systemType = SYSTEM_CLASS_G
		system.color = CLASS_G_COLOR
	} else if odds <= 0.605 {
		system.systemType = SYSTEM_CLASS_K
		system.color = CLASS_K_COLOR
	} else {
		system.systemType = SYSTEM_CLASS_M
		system.color = CLASS_M_COLOR
	}

}

func (system *System) GenerateSize() {

	// if we have a class M system , return 1.0
	if system.systemType == SYSTEM_CLASS_O {
		system.size = 1.0
		return
	}

	lowerBound := 30
	upperBound := 40

	if system.systemType == SYSTEM_CLASS_B {
		lowerBound = 80
		upperBound = 90
	}
	if system.systemType == SYSTEM_CLASS_A {
		lowerBound = 70
		upperBound = 80
	}
	if system.systemType == SYSTEM_CLASS_F {
		lowerBound = 60
		upperBound = 70
	}
	if system.systemType == SYSTEM_CLASS_F {
		lowerBound = 50
		upperBound = 60
	}
	if system.systemType == SYSTEM_CLASS_F {
		lowerBound = 40
		upperBound = 50
	}

	var size float32 = (float32(rand.Intn(upperBound-lowerBound)) + float32(lowerBound)) / 100.0
	system.size = size

	maxOffset := 1 - system.size

	system.offset[0] = rand.Float32() * maxOffset
	system.offset[1] = rand.Float32() * maxOffset

}

func (system *System) Exists() bool {
	return system.exists
}

func (system *System) Size() float32 {
	return system.size
}

func (system *System) Offset() glm.Vec2 {
	return system.offset
}
func (system *System) Color() util.Color {
	return system.color
}
