package util

import "github.com/micahke/infinite-universe/mango/util/color"


type UVSpriteMap struct {
	whiteChannel color.Color
	blackChannel color.Color
}

func (uv *UVSpriteMap) SetWhiteChannel(color color.Color) {
	uv.whiteChannel = color
}

func (uv *UVSpriteMap) SetBlackChannel(color color.Color) {
	uv.blackChannel = color
}

func (uv *UVSpriteMap) White() color.Color {
	return uv.whiteChannel
}

func (uv *UVSpriteMap) Black() color.Color {
	return uv.blackChannel
}
