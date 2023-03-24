package im

import (
	"github.com/micahke/infinite-universe/mango/opengl"
	"github.com/micahke/infinite-universe/mango/util/loaders"
)

type TextureCache map[string]*opengl.Texture

var textureCache TextureCache

func InitTextureCache() {
	textureCache = make(map[string]*opengl.Texture)
}

// Get the texture assigned a a specific path
func getTexture(texturePath string, deferred bool) *opengl.Texture {
	// If we find the texture, return it
	if texture, ok := textureCache[texturePath]; ok {
		return texture
	}

	// Otherwise, create a texture with the image date from the png loader
	imageData := loaders.LoadPNG(texturePath)
	texture := opengl.NewTextureFromData(texturePath, imageData, deferred)

	// cache the data
	textureCache[texturePath] = texture

	return texture

}

func GetTextureCache() *TextureCache {
	return &textureCache
}
