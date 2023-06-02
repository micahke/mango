package loaders

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/disintegration/imaging"
)

type ImageType byte

const (
	PNG ImageType = iota
	JPEG
)

type ImageLoader struct {
  Type ImageType
  Image *image.Image
}

func NewImageLoader(fileType ImageType) *ImageLoader {
	return &ImageLoader{
    Type: fileType,
  }
}

// Returns the raw, flipped version of the image
func (loader *ImageLoader) LoadImageRaw(name string) (*image.Image, error) {
	data, err := GetFileData(name)
	if err != nil {
		return nil, err
	}

  var img *image.Image
  img, err = loader.decodeImage(data)
  if err != nil {
    return nil, err
  }

  loader.Image = img
  
  return img, nil
}


func (loader *ImageLoader) LoadImage(name string) (*image.Image, error) {
  img, err := loader.LoadImageRaw(name)
  if err != nil {
    return nil, err
  }

  loader.FlipImageV()

  return img, nil

}


func (loader *ImageLoader) FlipImageV() {
  data := imaging.FlipV(*loader.Image)
  *loader.Image = image.Image(data)
}


func (loader *ImageLoader) FlipImageH() {
  data := imaging.FlipH(*loader.Image)
  *loader.Image = image.Image(data)
}


func (loader *ImageLoader) decodeImage(data []byte) (*image.Image, error) {
  var image image.Image
  var err error
  buffer := bytes.NewBuffer(data)
  switch loader.Type {
  case PNG:
    image, err = png.Decode(buffer)
  case JPEG:
    image, err = jpeg.Decode(buffer)
  default:
  return nil, fmt.Errorf("Image type not supported")
  }

  if err != nil {
    return nil, err
  }
  
  return &image, nil

}
