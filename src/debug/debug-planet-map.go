package debug

import (
	"fmt"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/infinite-universe/mango"
	"github.com/micahke/infinite-universe/mango/util"
	"github.com/micahke/infinite-universe/src/galaxy"
)



type PlanetMapDebugPanel struct {
  str string

  showTileMap *bool
  tileSize *float32
}


func NewPlanetMapDebugPanel(showTileMap *bool, tileSize *float32) *PlanetMapDebugPanel {
  panel := new(PlanetMapDebugPanel)

  panel.showTileMap = showTileMap
  panel.tileSize = tileSize


  return panel

}


func (panel *PlanetMapDebugPanel) RenderPanel() {

  imgui.SetNextWindowSize(imgui.Vec2{X: 300, Y: 400})

  fpsCounter := fmt.Sprint("Frametime: ", fmt.Sprintf("%.2f", mango.Time.FrameTime() * 1000)) + "ms"
  fpsCounter += fmt.Sprint(", ", fmt.Sprintf("%.2f", mango.Time.FPS())) + " FPS"

  imgui.BeginV("Debug", util.ImguiPanelStatus("planetMap"), 0)
  {
    imgui.Text(fpsCounter)
  }

  imgui.Spacing()
  imgui.Separator()
  imgui.Spacing()

  if imgui.TreeNode("Tilemap Settings") {

    imgui.Checkbox("Show Tile Map", panel.showTileMap)
    imgui.SameLine()
    if imgui.Button("Reset Tile Size") {
      *panel.tileSize = 50
    }
    imgui.SliderFloat("Tile Size", panel.tileSize, 1, 100)

    imgui.TreePop()
  }

  if imgui.TreeNode("Galaxy Settings") {
    imgui.SliderFloat("Alpha", &galaxy.GALAXY_ALPHA, 0, 5)
    imgui.SliderFloat("Beta", &galaxy.GALAXY_BETA, 0, 3)
    imgui.SliderInt("Iterations", &galaxy.GALAXY_N, 1, 10)
    imgui.InputInt("Seed", &galaxy.GALAXY_SEED)

      galaxy.Rebuild()

    imgui.TreePop()
  }

  imgui.Spacing()
  imgui.Separator()
  imgui.Spacing()


  imgui.End()


}
