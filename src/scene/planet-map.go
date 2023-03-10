package scene

import (
	"math"

	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/input"
	"github.com/micahke/infinite-universe/mango/util"
	"github.com/micahke/infinite-universe/src/debug"
	"github.com/micahke/infinite-universe/src/galaxy"
)

type PlanetMap struct {
	WINDOW_WIDTH  int
	WINDOW_HEIGHT int

	tileSize float32
	xTiles   int
	yTiles   int

	xOffset float64
	yOffset float64

	mapDebugPanel *debug.PlanetMapDebugPanel

	showTileMap bool
}

func (planetMap *PlanetMap) Init() {
  galaxy.Init()
	planetMap.tileSize = 50.0
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

  // planetMap.xOffset += 50.0 * float64(deltaTime)

	planetMap.xTiles = planetMap.WINDOW_WIDTH / int(planetMap.tileSize)
	planetMap.yTiles = planetMap.WINDOW_HEIGHT / int(planetMap.tileSize)

  planetMap.xOffset += 50.0 * float64(deltaTime)

}

func (planetMap *PlanetMap) Draw() {
  xOffsetBlocks := math.Floor(planetMap.xOffset / float64(planetMap.tileSize))
  yOffsetBlocks := math.Floor(planetMap.yOffset / float64(planetMap.tileSize))


  for x := 0; x < planetMap.xTiles+2; x++ {
    for y := 0; y < planetMap.yTiles+2; y++ {
      xCoord := int64(x) + int64(xOffsetBlocks)
      yCoord := int64(y) + int64(yOffsetBlocks)
      // finalX := xCoord * int64(planetMap.tileSize)
      // finalY := yCoord * int64(planetMap.tileSize)
      finalX := int64(math.Floor(float64(xCoord) * float64(planetMap.tileSize)))
      finalY := int64(math.Floor(float64(yCoord) * float64(planetMap.tileSize)))

      if planetMap.showTileMap {
        planetMap.drawDebugBG(float32(finalX), float32(finalY), xCoord, yCoord)
      }

      system := galaxy.NewSystem(xCoord, yCoord)

      if system.Exists() {

      systemSize := system.Size() * planetMap.tileSize
        xOff := system.Offset()[0] * planetMap.tileSize
        yOff := system.Offset()[1] * planetMap.tileSize

        mango.IM.DrawCircle(float32(finalX) - float32(planetMap.xOffset) + xOff, float32(finalY) - float32(planetMap.yOffset) + yOff, systemSize, systemSize, util.PINK_GLAMOUR)
      }

    }
  }



}



func (planetMap *PlanetMap) drawDebugBG(x, y float32, xCoord, yCoord int64) {


  normedPValue := galaxy.PerlinValueAtCoords(xCoord, yCoord, true)
  color := util.NewColorRGBf(float32(normedPValue), float32(normedPValue), float32(normedPValue))



  mango.IM.FillRect(x - float32(planetMap.xOffset) + 1, y - float32(planetMap.yOffset) + 1, planetMap.tileSize-2, planetMap.tileSize-2, color)



}
