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

  mango.IM.DrawLine(float32(width) / 2.0, float32(height) / 2.0, float32(input.MouseX), float32(input.MouseY), util.WHITE)

}




func plotLine(x0, y0, x1, y1 int, size int) {
	dx := x1 - x0
	dy := y1 - y0
	xi := 1
	yi := 1
	if dx < 0 {
		xi = -1
		dx = -dx
	}
	if dy < 0 {
		yi = -1
		dy = -dy
	}

	x := x0
	y := y0

	if dx > dy {
		// line is more horizontal
		D := (2 * dy) - dx
		for x != x1 {
			if D > 0 {
				y += yi
				D += 2*(dy - dx)
			} else {
				D += 2 * dy
			}
			mango.IM.FillRect(float32(x-size/2), float32(y-size/2), float32(size), float32(size), util.WHITE)
			x += xi
		}
	} else {
		// line is more vertical
		D := (2 * dx) - dy
		for y != y1 {
			if D > 0 {
				x += xi
				D += 2*(dx - dy)
			} else {
				D += 2 * dx
			}
			mango.IM.FillRect(float32(x-size/2), float32(y-size/2), float32(size), float32(size), util.WHITE)
			y += yi
		}
	}
}


