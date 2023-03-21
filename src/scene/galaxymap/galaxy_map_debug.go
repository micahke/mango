package galaxymap

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/util"
	"github.com/micahke/infinite-universe/src/galaxy"
)


type GalaxyMapDebugPanel struct {

  galaxyMap *GalaxyMap

  RenderPerlinNoise bool
  RenderBackground bool
  RenderPlanets bool
  RenderTilemap bool

  LiveRebuild bool

}



func InitGMDebugPanel(galaxyMap *GalaxyMap) *GalaxyMapDebugPanel {

  panel := new(GalaxyMapDebugPanel)
  panel.galaxyMap = galaxyMap
  panel.RenderBackground = true
  panel.RenderPlanets = true
  panel.RenderTilemap = true
  panel.LiveRebuild = false

  return panel

}


func (panel *GalaxyMapDebugPanel) RenderPanel() {


  imgui.BeginV("Galaxy Map Settings", util.ImguiPanelStatus("galaxyMap"), 0)

    
	fpsCounter := fmt.Sprint("Frametime: ", fmt.Sprintf("%.2f", mango.Time.FrameTime()*1000)) + "ms"
	fpsCounter += fmt.Sprint(", ", fmt.Sprintf("%.2f", mango.Time.FPS())) + " FPS"

	{
		imgui.Text(fpsCounter)
	}

	imgui.Spacing()
	imgui.Separator()
	imgui.Spacing()

	if imgui.TreeNode("Tilemap Settings") {

    imgui.Checkbox("Render Tilemap", &panel.RenderTilemap)
    imgui.Checkbox("Render Perlin Layer", &panel.RenderPerlinNoise)
    imgui.Checkbox("Render Background", &panel.RenderBackground)
    imgui.Checkbox("Render Planets", &panel.RenderPlanets)
    imgui.SliderFloat("Tile Size", &tilemap.proxyTileSize, 5.0, 100.0)

		imgui.TreePop()
	}

	if imgui.TreeNode("Galaxy Settings") {
    imgui.Checkbox("Live Rebuild", &panel.LiveRebuild)
		imgui.SliderFloat("Alpha", &galaxy.GALAXY_ALPHA, 0, 5)
		imgui.SliderFloat("Beta", &galaxy.GALAXY_BETA, 0, 3)
		imgui.SliderInt("Iterations", &galaxy.GALAXY_N, 1, 10)
		imgui.InputInt("Seed", &galaxy.GALAXY_SEED)
		imgui.SliderInt("System Frequency", &galaxy.SYSTEM_GENERATION_THRESHOLD, 0, 20)
		imgui.SliderFloat("System Scaling", &galaxy.GALAXY_FREQ, 0, 20)


    if panel.LiveRebuild {
      galaxy.Rebuild()
    }

    imgui.TreePop()
  }

  imgui.End()


}