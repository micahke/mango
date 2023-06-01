package core

import (
	"fmt"
	"image"

	"github.com/AllenDang/imgui-go"
	"github.com/micahke/mango/logging"
	"github.com/micahke/mango/opengl"
	"github.com/micahke/mango/util"
	"github.com/micahke/mango/util/loaders"
)


type ImguiSandbox struct {
  Image *image.Image
}


func InitImguiSandbox() *ImguiSandbox {

  
  sandbox := new(ImguiSandbox)

  imageLoader := loaders.NewImageLoader(loaders.PNG)
  img, err := imageLoader.LoadImage("quicktime.png")
  imageLoader.FlipImageV()
  if err != nil {
    fmt.Println("Error loading sandbox image")
  }

  sandbox.Image = img
  fmt.Println(sandbox.Image)

  return sandbox
}


func (sandbox *ImguiSandbox) RenderPanel() {
  imgui.SetNextWindowSizeV(util.ImguiGenVec2(400, 400), imgui.ConditionOnce)
  imgui.Begin("Sandbox")

  data, ok := (*sandbox.Image).(*image.NRGBA)
  if !ok {
    logging.DebugLogError("Error running sandbox image")
    fmt.Println()
  }

  texture := opengl.NewTextureFromData("quicktime.png", data, false)

  imgui.Image(imgui.TextureID(texture.GetID()), util.ImguiGenVec2(400, 400))

  sandbox.endFrame()
}


func (sandbox *ImguiSandbox) endFrame() {
  imgui.End()
}
