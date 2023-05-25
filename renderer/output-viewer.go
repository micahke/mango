package renderer

import (
	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/util"
)


type OutputViewer struct {
  textureID *uint32
}


func NewOutputViewer(id *uint32) *OutputViewer {
  return &OutputViewer{
    textureID: id,
  }
}


func (viewer *OutputViewer) RenderPanel() {

  imgui.Begin("Hello")

  imgui.Text("Hello")
  imgui.Image(imgui.TextureID(*viewer.textureID), util.ImguiGenVec2(400, 400))

  imgui.End()
}



