package galaxymap

import (
	"math"

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

	// xp, yp := getLinePoints([]float32{planetCoords[0], planetCoords[1]}, []float32{float32(width) / 2.0, float32(height) / 2.0, float32(input.MouseX)}, 50)
	// xp2, yp2 := getLinePoints([]float32{float32(width) / 2.0, float32(height) / 2.0, float32(input.MouseX)}, []float32{planetCoords[0], planetCoords[1]}, 50)

	if systemManager.calcDistanceFrom(glm64.Vec2{input.MouseX, input.MouseY}, glm64.Vec2{float64(width) / 2.0, float64(height) / 2.0}) >= 50 {
		// mango.IM.DrawLine(xp2, yp2, xp, yp, util.WHITE, 1.0)
	}

  mango.IM.DrawSprite(0, 0, 300, 300, "BitmapFont.png")

}

func getLinePoints(p1, p2 []float32, dist float32) (float32, float32) {
	var dx, dy float32
	if p1[0] > p2[0] {
		slope := (p1[1] - p2[1]) / (p1[0] - p2[0])
		angle := math.Atan(float64(slope))
		dx = dist * float32(math.Cos(angle))
		dy = dist * float32(math.Sin(angle))
	} else {
		slope := (p2[1] - p1[1]) / (p2[0] - p1[0])
		angle := math.Atan(float64(slope))
		dx = -dist * float32(math.Cos(angle))
		dy = -dist * float32(math.Sin(angle))
	}
	x := p2[0] + dx
	y := p2[1] + dy

	return x, y
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
				D += 2 * (dy - dx)
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
				D += 2 * (dx - dy)
			} else {
				D += 2 * dx
			}
			mango.IM.FillRect(float32(x-size/2), float32(y-size/2), float32(size), float32(size), util.WHITE)
			y += yi
		}
	}
}
