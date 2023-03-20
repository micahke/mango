package galaxymap

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/util"
)


type GalaxyMapDebugPanel struct {

  galaxyMap *GalaxyMap

  RenderPerlinNoise bool
  RenderBackground bool
  RenderPlanets bool
  RenderTilemap bool

}



func InitGMDebugPanel(galaxyMap *GalaxyMap) *GalaxyMapDebugPanel {

  panel := new(GalaxyMapDebugPanel)
  panel.galaxyMap = galaxyMap
  panel.RenderBackground = true
  panel.RenderPlanets = true
  panel.RenderTilemap = true

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


  imgui.End()


}
