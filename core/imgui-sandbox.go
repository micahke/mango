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
  sbImage *image.Image
  textureID uint32
}


func InitImguiSandbox() *ImguiSandbox {

  
  sandbox := new(ImguiSandbox)

  imageLoader := loaders.NewImageLoader(loaders.PNG)
  img, err := imageLoader.LoadImage("quicktime.png")
  if err != nil {
    fmt.Println("Error loading sandbox image")
  }
  imageLoader.FlipImageV()

  sandbox.sbImage = img
  data, ok := (*sandbox.sbImage).(*image.NRGBA)
  if !ok {
    logging.DebugLogError("Error running sandbox image")
    fmt.Println()
  }

  texture := opengl.NewTextureFromData("quicktime.png", data, false)
  sandbox.textureID = texture.GetID()

  return sandbox
}


func (sandbox *ImguiSandbox) RenderPanel() {
  imgui.SetNextWindowSizeV(util.ImguiGenVec2(400, 400), imgui.ConditionOnce)
  imgui.Begin("Sandbox")

  imgui.Image(imgui.TextureID(sandbox.textureID), util.ImguiGenVec2(400, 400))

  sandbox.endFrame()
}


func (sandbox *ImguiSandbox) endFrame() {
  imgui.End()
}
