package debug

import (
	"fmt"
	"inifinite-universe/galaxy"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango"
	"github.com/micahke/mango/util"
)

type PlanetMapDebugPanel struct {
	str string

	showTileMap     *bool
	renderPlanets   *bool
	renderCrosshair *bool
	renderShip      *bool
	tileSize        *float32
	cameraSpeed     *float32

  hoverCamera bool

	liveRebuild bool
}

func NewPlanetMapDebugPanel(showTileMap *bool, tileSize *float32, renderPlanets *bool, renderCrosshair *bool, renderShip *bool, cameraSpeed *float32) *PlanetMapDebugPanel {
	panel := new(PlanetMapDebugPanel)

	panel.showTileMap = showTileMap
	panel.tileSize = tileSize
	panel.renderPlanets = renderPlanets
	panel.renderCrosshair = renderCrosshair
	panel.renderShip = renderShip
	panel.cameraSpeed = cameraSpeed
	panel.liveRebuild = false

	return panel

}

func (panel *PlanetMapDebugPanel) RenderPanel() {

	// imgui.SetNextWindowSize(imgui.Vec2{X: 300, Y: 400})

	fpsCounter := fmt.Sprint("Frametime: ", fmt.Sprintf("%.2f", mango.Time.FrameTime()*1000)) + "ms"
	fpsCounter += fmt.Sprint(", ", fmt.Sprintf("%.2f", mango.Time.FPS())) + " FPS"

	imgui.BeginV("Debug", util.ImguiPanelStatus("planetMap"), 0)
	{
		imgui.Text(fpsCounter)
	}

	imgui.Spacing()
	imgui.Separator()
	imgui.Spacing()

	if imgui.TreeNode("Tilemap Settings") {

		imgui.Checkbox("Render Tile Map", panel.showTileMap)
		imgui.Checkbox("Render Planets", panel.renderPlanets)
		imgui.Checkbox("Render Crosshair", panel.renderCrosshair)
		imgui.Checkbox("Render Ship", panel.renderShip)
		if imgui.Button("Reset Tile Size") {
			*panel.tileSize = 50
		}
		imgui.SliderFloat("Tile Size", panel.tileSize, 1, 100)
		imgui.SliderFloat("Max Ship Speed", panel.cameraSpeed, 0.0, 1000.0)

		imgui.TreePop()
	}

	if imgui.TreeNode("Galaxy Settings") {
		imgui.Checkbox("Live Rebuild", &panel.liveRebuild)
		imgui.SliderFloat("Alpha", &galaxy.GALAXY_ALPHA, 0, 5)
		imgui.SliderFloat("Beta", &galaxy.GALAXY_BETA, 0, 3)
		imgui.SliderInt("Iterations", &galaxy.GALAXY_N, 1, 10)
		imgui.InputInt("Seed", &galaxy.GALAXY_SEED)
		imgui.SliderInt("System Frequency", &galaxy.SYSTEM_GENERATION_THRESHOLD, 0, 20)
		imgui.SliderFloat("System Scaling", &galaxy.GALAXY_FREQ, 0, 20)

		if panel.liveRebuild {
			galaxy.Rebuild()
		}

		imgui.TreePop()
	}

	imgui.Spacing()
	imgui.Separator()
	imgui.Spacing()

	imgui.End()

}
