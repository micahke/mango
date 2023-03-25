package galaxymap

import (
	"fmt"

	glm "github.com/go-gl/mathgl/mgl32"
	glm64 "github.com/go-gl/mathgl/mgl64"
	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/input"
	"github.com/micahke/infinite-universe/mango/util"
	"github.com/micahke/infinite-universe/src/galaxy"
)

type GalaxyMap struct{}

var tilemap *Tilemap
var systemManager *SystemManager

var width int
var height int

var BG_COLOR util.Color = util.NewColorRGBi(48, 42, 47)

var DEBUG_PANEL *GalaxyMapDebugPanel

func (gm *GalaxyMap) Init() {

	// Initialize the galaxy
	galaxy.Init()

	width, height = mango.GetWindow().GetSize()

	// Initialize systems
	tilemap = InitTilemap(width, height)
	systemManager = InitSystemManager()

	// Debug Panel
	DEBUG_PANEL = InitGMDebugPanel(gm)
	util.ImguiRegisterPanel("galaxyMap", DEBUG_PANEL)

}

func (gm *GalaxyMap) Update(deltaTime float64) {

	if input.MouseRightPressed {
		util.ImguiTogglePanel("galaxyMap")
	}

	tilemap.Update()
	systemManager.Update()
}

func (gm *GalaxyMap) Draw() {

	if DEBUG_PANEL.RenderTilemap {
		tilemap.Draw()
	}

	if DEBUG_PANEL.RenderPlanets {
		systemManager.Draw()
	}

	mango.IM.DrawSprite((float32(width)/2.0)-(float32(tilemap.tileSize)/2.0), (float32(height)/2.0)-(float32(tilemap.tileSize)/2.0), float32(tilemap.tileSize), float32(tilemap.tileSize), "tie-fighter.png")

	var planetCoords glm.Vec2
	planetCoords[0] = float32(CLOSEST_GALAXY_TO_MOUSE.parallaxCoords[0]) + float32(CLOSEST_GALAXY_TO_MOUSE.pixelSize/2.0)
	planetCoords[1] = float32(CLOSEST_GALAXY_TO_MOUSE.parallaxCoords[1]) + float32(CLOSEST_GALAXY_TO_MOUSE.pixelSize/2.0)

	if systemManager.calcDistanceFrom(glm64.Vec2{input.MouseX, input.MouseY}, glm64.Vec2{float64(width) / 2.0, float64(height) / 2.0}) >= 50 {
		// mango.IM.DrawLine(xp2, yp2, xp, yp, util.WHITE, 1.0)
	}

  closestSystem := CLOSEST_GALAXY_TO_MOUSE.system
  str := "x:" + fmt.Sprint(closestSystem.GetCoords()[0]) + ", y:" + fmt.Sprint(closestSystem.GetCoords()[1])
  if DEBUG_PANEL.RenderCoordinates {
  mango.IM.DrawText(10, float32(height) - 30, 30, "Infinite Universe")
  mango.IM.DrawText(10, float32(height) - 60, 30, str)
  }

}

