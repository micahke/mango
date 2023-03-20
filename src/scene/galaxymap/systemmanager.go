package galaxymap

import (

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/micahke/infinite-universe/mango"
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

  tileOffset glm.Vec2
	tilemapCoords glm.Vec2

  parallaxCoords glm.Vec2
}

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
    uv.SetWhiteChannel(data.system.Color())
    uv.SetBlackChannel(util.DarkenColor(data.system.Color(), 0.5))


    mango.IM.DrawUVSprite(float32(data.parallaxCoords[0]), float32(data.parallaxCoords[1]), float32(data.pixelSize), float32(data.pixelSize), "pixel-system.png", uv)
    

	}

}

// +++++++++++++++++ HELPER FUNCTIONS ++++++++++++++++

func (m *SystemManager) generateSystems() {
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

      coords := m.calculateScreenDistance(systemPosition.tilemapCoords[0], systemPosition.tilemapCoords[1])
      coords = coords.Mul(float64(system.ParallaxEffect()))

      if DEBUG_PANEL.RenderPerlinNoise {
        coords = coords.Mul(0.0)
      }

      systemPosition.parallaxCoords[0] = systemPosition.tilemapCoords[0] - coords[0]
      systemPosition.parallaxCoords[1] = systemPosition.tilemapCoords[1] - coords[1]

			m.systemData = append(m.systemData, systemPosition)
		}

	}
}

func (m *SystemManager) calculateScreenDistance(x, y float64) glm.Vec2 {
  distanceX := (float64(width) / 2.0) - x
  distanceY := (float64(height) / 2.0) - y

  return glm.Vec2{distanceX, distanceY}
}
