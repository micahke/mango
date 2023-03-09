package scene

import (
	"math"

	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/input"
	"github.com/micahke/infinite-universe/mango/util"
	"github.com/micahke/infinite-universe/src/debug"
)


type PlanetMap struct {
  WINDOW_WIDTH int
  WINDOW_HEIGHT int

  tileSize float32
  xTiles int
  yTiles int

  xOffset float64
  yOffset float64

  mapDebugPanel *debug.PlanetMapDebugPanel

  showTileMap bool


}



func (planetMap *PlanetMap) Init() {
  planetMap.tileSize = 50
  planetMap.xTiles = planetMap.WINDOW_WIDTH / int(planetMap.tileSize)
  planetMap.yTiles = planetMap.WINDOW_HEIGHT / int(planetMap.tileSize)
  planetMap.showTileMap = true
  planetMap.mapDebugPanel = debug.NewPlanetMapDebugPanel(&planetMap.showTileMap, &planetMap.tileSize)

  util.ImguiRegisterPanel("planetMap", planetMap.mapDebugPanel)
  util.ImguiActivatePanel("planetMap")

}


func (planetMap *PlanetMap) Update(deltaTime float32) {
  if input.MouseRightPressed {
    util.ImguiActivatePanel("planetMap")
  }

  if input.MouseLeftPressed {
    clickedX = float32(input.MouseX)
    clickedY = float32(input.MouseY)
  }
}

func (planetMap *PlanetMap) Draw() {

  if (planetMap.showTileMap) {
    planetMap.drawDebugBG()
  }

}

var (
  clickedX float32
  clickedY float32
)

func (planetMap *PlanetMap) drawDebugBG() {
  xOffsetBlocks := math.Floor(planetMap.xOffset / float64(planetMap.tileSize))
  yOffsetBlocks := math.Floor(planetMap.yOffset / float64(planetMap.tileSize))

  for x := 0; x < planetMap.xTiles + 2; x++ {
    for y := 0; y < planetMap.yTiles + 2; y++ {
      xCoord := int64(x) + int64(xOffsetBlocks)
      yCoord := int64(y) + int64(yOffsetBlocks)

      finalX := xCoord * int64(planetMap.tileSize)
      finalY := yCoord * int64(planetMap.tileSize)

      color := util.NewColorRGBf(0.5, 0.5, 0.5)

      if float32(clickedX) >= float32(x) * planetMap.tileSize  && float32(clickedX) <= float32(x) * planetMap.tileSize + planetMap.tileSize {
        if float32(clickedY) >= float32(y) * planetMap.tileSize && float32(clickedY) <= float32(y) * planetMap.tileSize + planetMap.tileSize {
          color = util.MINT_GREEN
        }
      }
      mango.IM.FillRect(float32(finalX) - float32(planetMap.xOffset) + 1, float32(finalY) - float32(planetMap.yOffset) + 1, planetMap.tileSize - 2, planetMap.tileSize - 2, color)
    }
  }
}

