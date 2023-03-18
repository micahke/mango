package util


type UVSpriteMap struct {

  whiteChannel Color
  blackChannel Color

}

func (uv *UVSpriteMap) SetWhiteChannel(color Color) {
  uv.whiteChannel = color
}

func (uv *UVSpriteMap) SetBlackChannel(color Color) {
  uv.blackChannel = color
}

func (uv *UVSpriteMap) White() Color {
  return uv.whiteChannel
}

func (uv *UVSpriteMap) Black() Color {
  return uv.blackChannel
}
