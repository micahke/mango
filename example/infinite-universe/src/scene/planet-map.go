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

	xVelo float64
	yVelo float64

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
	// planetMap.bgColor = util.NewColorRGBi(30, 26, 29)
	planetMap.bgColor = util.NewColorRGBi(60, 52, 58)
	planetMap.bgColor = util.NewColorRGBi(48, 42, 47)
	// planetMap.bgColor = util.PINK_GLAMOUR

	// mango.IM.SetBackgroundColor(planetMap.bgColor)

	// Debug
	planetMap.mapDebugPanel = debug.NewPlanetMapDebugPanel(&planetMap.showTileMap, &planetMap.tileSize, &planetMap.renderPlanets, &planetMap.renderCrosshair, &planetMap.renderShip, &planetMap.cameraSpeed)

	util.ImguiRegisterPanel("planetMap", planetMap.mapDebugPanel)

}

func (planetMap *PlanetMap) Update(deltaTime float64) {

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

			system := galaxy.NewSystem(int(xCoord), int(yCoord), false)

			if system.Exists() {

				systemSize := system.Size() * float64(planetMap.tileSize)
				xOff := system.Offset()[0] * float64(planetMap.tileSize)
				yOff := system.Offset()[1] * float64(planetMap.tileSize)

				screenX := float32(finalX) - float32(planetMap.xOffset)
				screenY := float32(finalY) - float32(planetMap.yOffset)

				coords := planetMap.calculateScreenDistance(screenX, screenY)
				coords = coords.Mul(system.ParallaxEffect())

				if planetMap.showTileMap {
					coords = coords.Mul(0)
				}

				if planetMap.renderPlanets {

					darkerColor := util.DarkenColor(system.Color(), 0.5)

					uvMap := util.UVSpriteMap{}
					uvMap.SetWhiteChannel(system.Color())
					uvMap.SetBlackChannel(darkerColor)
					mango.IM.DrawUVSprite(screenX+float32(xOff)-coords[0], screenY+float32(yOff)-coords[1], float32(systemSize), float32(systemSize), "pixel-system.png", uvMap)
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
			alpha := float32(galaxy.PerlinValueAtCoords(int(xCoord), int(yCoord), true))
			bgTileColor := util.NewColorRGBAf(planetMap.bgColor.X(), planetMap.bgColor.Y(), planetMap.bgColor.Z(), float32(alpha))
			mango.IM.FillRect(float32(finalX)-float32(planetMap.xOffset), float32(finalY)-float32(planetMap.yOffset), planetMap.tileSize, planetMap.tileSize, bgTileColor)
		}
	}
}

func (planetMap *PlanetMap) calculateScreenDistance(x, y float32) mgl32.Vec2 {
	distanceX := float32(planetMap.WINDOW_WIDTH/2.0) - x
	distanceY := float32(planetMap.WINDOW_HEIGHT/2.0) - y

	return mgl32.Vec2{distanceX, distanceY}
}

func (planetMap *PlanetMap) drawDebugBG(x, y float32, xCoord, yCoord int64) {

	normedPValue := galaxy.PerlinValueAtCoords(int(xCoord), int(yCoord), true)
	color := util.NewColorRGBf(float32(normedPValue), float32(normedPValue), float32(normedPValue))

	mango.IM.FillRect(x-float32(planetMap.xOffset)+1, y-float32(planetMap.yOffset)+1, planetMap.tileSize-2, planetMap.tileSize-2, color)

}

func (planetMap *PlanetMap) controlShip() {
	// change := float64(planetMap.cameraSpeed) * float64(mango.Time.DeltaTime())
	thrust := 2.5
	if input.GetKey(input.KEY_A) {
		// planetMap.xOffset -= change
		planetMap.xVelo -= thrust
	}
	if input.GetKey(input.KEY_D) {
		// planetMap.xOffset += change
		planetMap.xVelo += thrust
	}
	if input.GetKey(input.KEY_W) {
		// planetMap.yOffset += change
		planetMap.yVelo += thrust
	}
	if input.GetKey(input.KEY_S) {
		// planetMap.yOffset -= change
		planetMap.yVelo -= thrust
	}

	planetMap.xVelo = math.Max(float64(-1.0*planetMap.cameraSpeed), math.Min(float64(planetMap.cameraSpeed), planetMap.xVelo))
	planetMap.yVelo = math.Max(float64(-1.0*planetMap.cameraSpeed), math.Min(float64(planetMap.cameraSpeed), planetMap.yVelo))

	planetMap.xOffset += planetMap.xVelo * mango.Time.DeltaTime()
	planetMap.yOffset += planetMap.yVelo * mango.Time.DeltaTime()

}
