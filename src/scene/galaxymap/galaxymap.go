package galaxymap

import (
	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/input"
	"github.com/micahke/infinite-universe/mango/util"
	"github.com/micahke/infinite-universe/src/galaxy"
)


type GalaxyMap struct {}

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
}
