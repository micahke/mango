package scene

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
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

	cameraSpeed float32

	showTileMap     bool
	renderPlanets   bool
	renderCrosshair bool
	renderShip      bool

	bgColor util.Color
}

func (planetMap *PlanetMap) Init() {
	galaxy.Init()
	planetMap.tileSize = 50.0
	planetMap.xTiles = planetMap.WINDOW_WIDTH / int(planetMap.tileSize)
	planetMap.yTiles = planetMap.WINDOW_HEIGHT / int(planetMap.tileSize)
	planetMap.showTileMap = false
	planetMap.renderPlanets = true
	planetMap.renderCrosshair = false
	planetMap.renderShip = true
	planetMap.cameraSpeed = 500.0
	planetMap.bgColor = util.NewColorRGBi(30, 26, 29)

	// mango.IM.SetBackgroundColor(planetMap.bgColor)

	// Debug
	planetMap.mapDebugPanel = debug.NewPlanetMapDebugPanel(&planetMap.showTileMap, &planetMap.tileSize, &planetMap.renderPlanets, &planetMap.renderCrosshair, &planetMap.renderShip, &planetMap.cameraSpeed)

	util.ImguiRegisterPanel("planetMap", planetMap.mapDebugPanel)

}

func (planetMap *PlanetMap) Update(deltaTime float32) {

  planetMap.controlShip()

	if input.MouseRightPressed {
		util.ImguiActivatePanel("planetMap")
	}

	// planetMap.xOffset += 50.0 * float64(deltaTime)

	planetMap.xTiles = planetMap.WINDOW_WIDTH / int(planetMap.tileSize)
	planetMap.yTiles = planetMap.WINDOW_HEIGHT / int(planetMap.tileSize)


}

func (planetMap *PlanetMap) Draw() {
	xOffsetBlocks := math.Floor(planetMap.xOffset / float64(planetMap.tileSize))
	yOffsetBlocks := math.Floor(planetMap.yOffset / float64(planetMap.tileSize))

  if !planetMap.showTileMap {
    planetMap.drawBG()
  }

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


			system := galaxy.NewSystem(xCoord, yCoord, false)

			if system.Exists() {

				systemSize := system.Size() * planetMap.tileSize
				xOff := system.Offset()[0] * planetMap.tileSize
				yOff := system.Offset()[1] * planetMap.tileSize

        screenX := float32(finalX)-float32(planetMap.xOffset)
        screenY := float32(finalY)-float32(planetMap.yOffset)

        coords := planetMap.calculateScreenDistance(screenX, screenY)
        coords = coords.Mul(0.8)

        if planetMap.showTileMap {
          coords = coords.Mul(0)
        }


				if planetMap.renderPlanets {
					// mango.IM.DrawCircle(float32(finalX) - float32(planetMap.xOffset) + xOff, float32(finalY) - float32(planetMap.yOffset) + yOff, systemSize, systemSize, system.Color())
					// mango.IM.DrawCircle(screenX+xOff - coords[0], screenY+yOff-coords[1], systemSize, systemSize, system.Color())
					mango.IM.DrawSprite(screenX+xOff - coords[0], screenY+yOff-coords[1], systemSize, systemSize, "pixel-system.png")
          // mango.IM.DrawSprite(float32(finalX)-float32(planetMap.xOffset)+xOff, float32(finalY)-float32(planetMap.yOffset)+yOff, systemSize, systemSize, system.Color())
				}
			}

		}
	}

	transWhite := util.NewColorRGBAf(1.0, 1.0, 1.0, 0.5)

	if planetMap.renderCrosshair {

		mango.IM.FillRect(0, float32(input.MouseY), float32(planetMap.WINDOW_WIDTH), 0.5, transWhite)
		mango.IM.FillRect(float32(input.MouseX), 0, 0.5, float32(planetMap.WINDOW_HEIGHT), transWhite)
	}

	if planetMap.renderShip {
		mango.IM.DrawSprite((float32(planetMap.WINDOW_WIDTH)/2.0)-(planetMap.tileSize/2.0), (float32(planetMap.WINDOW_HEIGHT)/2.0)-(planetMap.tileSize/2), planetMap.tileSize, planetMap.tileSize, "tie-fighter.png")
	}

}


func (planetMap *PlanetMap) drawBG() {
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
      alpha := float32(galaxy.PerlinValueAtCoords(xCoord, yCoord, true))
      bgTileColor := util.NewColorRGBAf(planetMap.bgColor.X(), planetMap.bgColor.Y(), planetMap.bgColor.Z(), alpha)
      mango.IM.FillRect(float32(finalX) - float32(planetMap.xOffset), float32(finalY) - float32(planetMap.yOffset), planetMap.tileSize, planetMap.tileSize, bgTileColor)
    }
  }
}

func (planetMap *PlanetMap) calculateScreenDistance(x, y float32) mgl32.Vec2 {
  distanceX := float32(planetMap.WINDOW_WIDTH / 2.0) - x
  distanceY := float32(planetMap.WINDOW_HEIGHT / 2.0) - y

  return mgl32.Vec2{distanceX, distanceY}

}

func (planetMap *PlanetMap) drawDebugBG(x, y float32, xCoord, yCoord int64) {

	normedPValue := galaxy.PerlinValueAtCoords(xCoord, yCoord, true)
	color := util.NewColorRGBf(float32(normedPValue), float32(normedPValue), float32(normedPValue))

	mango.IM.FillRect(x-float32(planetMap.xOffset)+1, y-float32(planetMap.yOffset)+1, planetMap.tileSize-2, planetMap.tileSize-2, color)

}


func (planetMap *PlanetMap) controlShip() {
  change := float64(planetMap.cameraSpeed) * float64(mango.Time.DeltaTime())
  if input.GetKey(input.KEY_A) {
    planetMap.xOffset -= change 
  }
  if input.GetKey(input.KEY_D) {
    planetMap.xOffset += change
  }
  if input.GetKey(input.KEY_W) {
    planetMap.yOffset += change
  }
  if input.GetKey(input.KEY_S) {
    planetMap.yOffset -= change
  }
}

