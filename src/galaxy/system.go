package galaxy

import (
	"math/rand"

	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/util"
)

var (
  SYSTEM_GENERATION_THRESHOLD int32 = 1
)


type System struct {

  exists bool
  
  size float32
  offset glm.Vec2
  color util.Color


}


func NewSystem(xCoord, yCoord int64) *System {
  system := new(System)

  perlinValue := PerlinValueAtCoords(xCoord, yCoord, true)

  if perlinValue < 0.3 {
    system.exists = false
    return system
  }
  
  seed := (xCoord & 0xFFFF) << 16 | (yCoord & 0xFFFF)
  rand.Seed(seed)

  if rand.Int31n(20) > SYSTEM_GENERATION_THRESHOLD {
    system.exists = false
    return system
  }

  system.exists = true
  system.size = (float32(rand.Int63n(50)) + 30.0) / 100.0
  maxOffset := 1 - system.size

  system.offset[0] = rand.Float32() * maxOffset
  system.offset[1] = rand.Float32() * maxOffset


  return system

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
