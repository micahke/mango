package scene

import (
	"math"

	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/util"
)


type PlanetMap struct {
  WINDOW_WIDTH int
  WINDOW_HEIGHT int

  tileSize float32
  xTiles int
  yTiles int

  xOffset float64
  yOffset float64

}

func (planetMap *PlanetMap) Init() {
  planetMap.tileSize = 50
  planetMap.xTiles = planetMap.WINDOW_WIDTH / int(planetMap.tileSize)
  planetMap.yTiles = planetMap.WINDOW_HEIGHT / int(planetMap.tileSize)
}


func (planetMap *PlanetMap) Update(deltaTime float32) {
  planetMap.xOffset += 10.0 * float64(deltaTime)
}

func (planetMap *PlanetMap) Draw() {

  planetMap.drawDebugBG()

}


func (planetMap *PlanetMap) drawDebugBG() {
  xOffsetBlocks := math.Floor(planetMap.xOffset / float64(planetMap.tileSize))

  for x := 0; x < planetMap.xTiles + 2; x++ {
    for y := 0; y < planetMap.yTiles + 2; y++ {
      xCoord := int64(x) + int64(xOffsetBlocks)
      yCoord := y * int(planetMap.tileSize)

      finalX := xCoord * int64(planetMap.tileSize)

      mango.IM.FillRect(float32(finalX) - float32(planetMap.xOffset) + 1, float32(yCoord) + 1, planetMap.tileSize - 2, planetMap.tileSize - 2, util.NewColorRGBf(0.5, 0.5, 0.5))
    }
  }
}
