package opengl

import (
	"embed"
	"io/fs"
	"strings"

	"github.com/micahke/mango/logging"
)

// TODO: come back to this, I'm sure it's horrible code, I just want it to work rn

//go:embed shaders/*.glsl
var shaderFiles embed.FS

var ShaderCache map[string]string

var ShaderNames []string

// Loads the shaders into the cache
func LoadShaders() (map[string]string, error) {

	logging.DebugLog("Loading shaders")

	shaders := make(map[string]string)
	files, err := fs.Glob(shaderFiles, "shaders/*.glsl")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		contents, err := fs.ReadFile(shaderFiles, file)
		if err != nil {
			return nil, err
		}
		// shaders[strings.TrimSuffix(file, ".glsl")] = string(contents)
		pathName := strings.Trim(file, "shaders/")
		shaders[pathName] = string(contents)
    ShaderNames = append(ShaderNames, pathName)
	}

	ShaderCache = shaders

	return ShaderCache, nil
}

func getShaderFromCache(name string) string {
	return ShaderCache[name]
}
