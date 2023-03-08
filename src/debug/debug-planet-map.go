package debug

import (
  "github.com/AllenDang/imgui-go"
  "github.com/micahke/infinite-universe/mango/util"
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

  // str := fmt.Sprint("Frametime: ", fmt.Sprintf("%.2f", mango.Time.FrameTime() * 1000)) + "ms"
  // str += fmt.Sprint(", ", fmt.Sprintf("%.2f", mango.Time.FPS())) + " FPS"

  imgui.Begin("Debug")
  {
    imgui.Checkbox("Show Tile Map", panel.showTileMap)
    imgui.SameLine()
    if imgui.Button("Reset Tile Size") {
      *panel.tileSize = 50
    }
    imgui.SliderFloat("Tile Size", panel.tileSize, 1, 100)

  }

  if imgui.Button("Close") {
    util.ImguiDeactivatePanel("planetMap")
  }

  imgui.End()


}
