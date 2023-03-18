package util


type UVSpriteMap struct {

  whiteChannel Color
  blackChannel Color

}



func (uv *UVSpriteMap) White() Color {
  return uv.whiteChannel
}

func (uv *UVSpriteMap) Black() Color {
  return uv.blackChannel
}
