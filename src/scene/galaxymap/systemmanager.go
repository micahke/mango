package galaxymap

import (
	"math"

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/input"
	"github.com/micahke/infinite-universe/mango/util"
	"github.com/micahke/infinite-universe/src/galaxy"
)

type SystemManager struct {
	systemData []*SystemPositionData
}

type SystemPositionData struct {
	system *galaxy.System

	pixelSize float64

	parallaxOffset float64

	tileOffset    glm.Vec2
	tilemapCoords glm.Vec2

	parallaxCoords glm.Vec2

	distanceFromMouse float32

	closest bool
}

var (
	CLOSEST_GALAXY_TO_MOUSE *SystemPositionData
)

// +++++++ HIGH LEVEL FUNCTIONS ++++++++++++=

func InitSystemManager() *SystemManager {

	manager := new(SystemManager)

	return manager

}

func (m *SystemManager) Update() {
	// reset the position buffer
	m.systemData = []*SystemPositionData{}

	// generate systems
	m.generateSystems()

}

func (m *SystemManager) Draw() {

	for _, data := range m.systemData {

		// Build UV map
		uv := util.UVSpriteMap{}

		// if this is the closest system, draw the uv sprite with a clear background and white border
		if data.closest {
			uv.SetWhiteChannel(util.NewColorRGBAf(0.0, 0.0, 0.0, 0.0))
			uv.SetBlackChannel(util.NewColorRGBAf(1.0, 1.0, 1.0, 0.85))
			var offset float32 = 30.0
			mango.IM.DrawUVSprite(float32(data.parallaxCoords[0])-offset/2, float32(data.parallaxCoords[1])-offset/2, float32(data.pixelSize)+offset, float32(data.pixelSize)+offset, "thinner-border.png", uv)
		}

		uv.SetWhiteChannel(data.system.Color())
		uv.SetBlackChannel(util.DarkenColor(data.system.Color(), 0.5))

		mango.IM.DrawUVSprite(float32(data.parallaxCoords[0]), float32(data.parallaxCoords[1]), float32(data.pixelSize), float32(data.pixelSize), "minimal.png", uv)

	}

}

// +++++++++++++++++ HELPER FUNCTIONS ++++++++++++++++

func (m *SystemManager) generateSystems() {

	var minDistance float64 = math.Inf(1)
	var minDistSystem *SystemPositionData = new(SystemPositionData)

	for _, data := range tilemap.tilePositions {

		system := galaxy.NewSystem(data.x, data.y, false)

		if system.Exists() {
			systemPosition := new(SystemPositionData)

			systemPosition.system = system
			systemPosition.pixelSize = tilemap.tileSize * float64(system.Size())

			systemPosition.tileOffset[0] = tilemap.tileSize * system.Offset()[0]
			systemPosition.tilemapCoords[0] = data.screenCoords[0] + systemPosition.tileOffset[0]

			systemPosition.tileOffset[1] = tilemap.tileSize * system.Offset()[1]
			systemPosition.tilemapCoords[1] = data.screenCoords[1] + systemPosition.tileOffset[1]

			coords := m.calcVectorDistance(systemPosition.tilemapCoords[0], systemPosition.tilemapCoords[1])
			coords = coords.Mul(float64(system.ParallaxEffect()))

			if DEBUG_PANEL.RenderPerlinNoise {
				coords = coords.Mul(0.0)
			}

			systemPosition.parallaxCoords[0] = systemPosition.tilemapCoords[0] - coords[0]
			systemPosition.parallaxCoords[1] = systemPosition.tilemapCoords[1] - coords[1]

			// Calculate the distance from the center of the screen
			distance := m.calcDistanceFrom(glm.Vec2{input.MouseX, input.MouseY}, systemPosition.parallaxCoords)
			systemPosition.distanceFromMouse = float32(distance)

			// If this is the first system, set it as the closest
			if minDistance == 0.0 {
				minDistance = distance
				minDistSystem = systemPosition
			}

			// If this system is closer than the previous closest, set it as the closest
			if distance < minDistance {
				minDistance = distance
				minDistSystem = systemPosition
			}

			m.systemData = append(m.systemData, systemPosition)
		}

	}
	minDistSystem.closest = true
	CLOSEST_GALAXY_TO_MOUSE = minDistSystem
}

func (m *SystemManager) calcVectorDistance(x, y float64) glm.Vec2 {
	distanceX := (float64(width) / 2.0) - x
	distanceY := (float64(height) / 2.0) - y

	return glm.Vec2{distanceX, distanceY}
}

func (m *SystemManager) calcDistanceFrom(p1, p2 glm.Vec2) float64 {
	return math.Sqrt(math.Pow(p1.X()-p2.X(), 2.0) + math.Pow(p1.Y()-p2.Y(), 2.0))
}
