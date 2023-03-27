package opengl

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/micahke/infinite-universe/mango/logging"
)

type Texture struct {
	m_RendererID      uint32
	m_FilePath        string
	m_LocalBuffer     []uint8
	m_Width, m_Height int32
	m_BPP             int
}

func NewTexture(path string) *Texture {
	texture := &Texture{}
	texture.m_RendererID = 0
	texture.m_FilePath = path
	texture.m_BPP = 0

	file, err := os.Open("res/" + path)
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("Error decoding file")
	}

	flippedImage := imaging.FlipV(img)
	texture.m_Height = int32(flippedImage.Rect.Dy())
	texture.m_Width = int32(flippedImage.Rect.Dx())
	texture.m_LocalBuffer = flippedImage.Pix

	gl.GenTextures(1, &texture.m_RendererID)
	gl.BindTexture(gl.TEXTURE_2D, texture.m_RendererID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, texture.m_Width, texture.m_Height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&texture.m_LocalBuffer[0]))

	gl.BindTexture(gl.TEXTURE_2D, 0)



	return texture
}

func NewTextureFromData(path string, data *image.NRGBA, deferredLoading bool) *Texture {
	texture := new(Texture)
	texture.m_FilePath = path
	texture.m_BPP = 0

	flippedImage := data
	texture.m_Height = int32(flippedImage.Rect.Dy())
	texture.m_Width = int32(flippedImage.Rect.Dx())
	texture.m_LocalBuffer = flippedImage.Pix

	gl.GenTextures(1, &texture.m_RendererID)
	gl.BindTexture(gl.TEXTURE_2D, texture.m_RendererID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

  ptr := gl.Ptr(&texture.m_LocalBuffer[0])

  if deferredLoading {
    ptr = nil
  }

  gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, texture.m_Width, texture.m_Height, 0, gl.RGBA, gl.UNSIGNED_BYTE, ptr)

	gl.BindTexture(gl.TEXTURE_2D, 1)

	return texture
}


// A really minimal and lightweight texture packet
func NewDataTexture(width, height int32, data []uint8) *Texture {

  texture := new(Texture)

	gl.GenTextures(1, &texture.m_RendererID)
	gl.BindTexture(gl.TEXTURE_2D, texture.m_RendererID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

  gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&data[0]))

  return texture

}

func (texture *Texture) SetSubTexture(x, y, width, height int, data []uint8) []uint8 {
  gl.BindTexture(gl.TEXTURE_2D, texture.m_RendererID)
  gl.PixelStorei(gl.UNPACK_ROW_LENGTH, int32(texture.m_Width))

  // Set the sub-image data
  gl.TexSubImage2D(gl.TEXTURE_2D, 0, int32(x), int32(y), int32(width), int32(height), gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&data[0]))

  // Get the updated sub-texture data and return it
  pixels := make([]uint8, width*height*4)
  gl.GetTexImage(gl.TEXTURE_2D, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&pixels[0]))

  return pixels
}

func (texture *Texture) UpdateSubImage(x, y, width, height int) {
  pixels := texture.m_LocalBuffer
  // gl.BindTexture(gl.TEXTURE_2D, texture.m_RendererID)
  gl.PixelStorei(gl.UNPACK_ROW_LENGTH, int32(texture.m_Width))

  // Calculate the starting index of the buffer based on the x and y position
  startIndex := (y * int(texture.m_Width) + x) * 4

  realPix := pixels[startIndex:startIndex+width*height]

  gl.TexSubImage2D(gl.TEXTURE_2D, 0, int32(x), int32(y), int32(width), int32(height), gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&realPix[0]))
}


func (texture *Texture) SetSubImageData(x, y, width, height int, data []uint8) {
  // gl.BindTexture(gl.TEXTURE_2D, texture.m_RendererID)
  gl.PixelStorei(gl.UNPACK_ROW_LENGTH, int32(texture.m_Width))

  gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, int32(width), int32(height), gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&data[0]))
}


func (texture *Texture) GetSubtextureData(x, y, width, height int) []uint8 {
	pixels := make([]uint8, width*height*4)

	gl.PixelStorei(gl.UNPACK_ROW_LENGTH, int32(texture.m_Width))
	gl.GetTexImage(gl.TEXTURE_2D, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))
  
  logging.DebugLog(pixels)
  return pixels


}

func (texture *Texture) Bind(slot uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + slot)
	gl.BindTexture(gl.TEXTURE_2D, texture.m_RendererID)
}

func (texture *Texture) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (texture *Texture) GetWidth() int32 {
	return texture.m_Width
}

func (texture *Texture) GetHeight() int32 {
	return texture.m_Height
}

func (texture *Texture) GetID() uint32 {
  return texture.m_RendererID
}

