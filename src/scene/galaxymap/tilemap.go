package galaxymap

import (
	"math"

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/util"
	"github.com/micahke/infinite-universe/src/galaxy"
	// "github.com/micahke/infinite-universe/mango/util"
)

type Tilemap struct {
	width  int
	height int

	tileSize      float64
	proxyTileSize float32
	numTiles      [2]int

	positions []glm.Vec2
	offset    glm.Vec2

	tilePositions []*TileData
}

type TileData struct {
	x int
	y int

	screenCoords glm.Vec2
}

// +++++++++++++++ HIGH LEVEL FUNCTIONS +++++++++++++++++++

func InitTilemap(width, height int) *Tilemap {

	tilemap := new(Tilemap)

	tilemap.tileSize = 50.0
	tilemap.proxyTileSize = 50.0
	tilemap.width = width
	tilemap.height = height

	return tilemap

}

func (t *Tilemap) Update() {
	t.tileSize = float64(t.proxyTileSize)

	tilemap.offset[0] += 50.0 * mango.Time.DeltaTime()
	tilemap.offset[1] += 50.0 * mango.Time.DeltaTime()

	// Reset the tile data
	t.tilePositions = []*TileData{}

	// figure out how many rows and columns to render
	t.calculateNumberOfTiles()

	// calculate draw positions
	for x := 0; x < t.numTiles[0]+2; x++ {
		for y := 0; y < t.numTiles[1]+2; y++ {
			tileData := new(TileData)
			tileData.x = x + int(math.Floor(t.offset[0]/t.tileSize))
			tileData.y = y + int(math.Floor(t.offset[1]/t.tileSize))
			tileData.screenCoords[0] = (float64(tileData.x) * t.tileSize) - t.offset[0]
			tileData.screenCoords[1] = (float64(tileData.y) * t.tileSize) - t.offset[1]

			t.tilePositions = append(t.tilePositions, tileData)
		}
	}
}

func (t *Tilemap) Draw() {
	for _, data := range t.tilePositions {
		pValue := galaxy.PerlinValueAtCoords(data.x, data.y, true)
		color := util.NewColorRGBf(0.5, 0.5, 0.5)
		gap := glm.Vec2{1.0, 2.0}
		if DEBUG_PANEL.RenderPerlinNoise {
			gap = glm.Vec2{1.0, 2.0}
			color = util.NewColorRGBAf(float32(pValue), float32(pValue), float32(pValue), float32(pValue))
		}
		if DEBUG_PANEL.RenderBackground {
			gap = gap.Mul(0.0)
			rawColor := BG_COLOR.Mul(float32(pValue))
			color = util.NewColorRGBAf(rawColor[0], rawColor[1], rawColor[2], 1.0)
			color = util.NewColorRGBAf(BG_COLOR.Vec4[0], BG_COLOR.Vec4[1], BG_COLOR.Vec4[2], float32(galaxy.PerlinValueAtCoords(data.x, data.y, true)))
		}
		mango.IM.FillRect(float32(data.screenCoords[0])+float32(gap[0]), float32(data.screenCoords[1])+float32(gap[0]), float32(t.tileSize)-float32(gap[1]), float32(t.tileSize)-float32(gap[1]), color)
	}
}

// +++++++++++++++ HELPER FUNCTIONS +++++++++++++++++++++++++

func (t *Tilemap) calculateNumberOfTiles() {
	// Calculate x component
	t.numTiles[0] = int(math.Ceil(float64(t.width) / t.tileSize))
	t.numTiles[1] = int(math.Ceil(float64(t.height) / t.tileSize))
}

