package opengl

import (
	"embed"
	"io/fs"
	"strings"
)

// TODO: come back to this, I'm sure it's horrible code, I just want it to work rn

//go:embed shaders/*.glsl
var shaderFiles embed.FS

var ShaderCache map[string]string

// Loads the shaders into the cache
func LoadShaders() (map[string]string, error) {
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
        shaders[strings.Trim(file, "shaders/")] = string(contents)
    }

    ShaderCache = shaders

    return ShaderCache, nil
}
