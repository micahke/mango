package loaders

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

type PNG_LOADER struct {
	assetPath string

	cache map[string]*image.NRGBA
}

var png_loader *PNG_LOADER

// Initialize the png loader instance
func InitPNGLoader() {

	png_loader = new(PNG_LOADER)

	png_loader.cache = make(map[string]*image.NRGBA)
	png_loader.assetPath = "assets"

	files, err := getPngFiles()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(files); i++ {
		load_png(files[i])
	}

}

// Loads an image from the image cache
// If no image is found, load from assets folder
func LoadPNG(path string) *image.NRGBA {
	filePath := png_loader.assetPath + "/" + path
	if value, ok := png_loader.cache[filePath]; ok {
		return value
	}
	return load_png(filePath)
}

// Loads an individual PNG file and returns it's data
// Also adds image data to cache
func load_png(path string) *image.NRGBA {

	file, err := os.Open(path)
	if err != nil {
		// TODO: Log the error
		fmt.Println("Error reading the file")
	}

	// Load the image data
	img, err := png.Decode(file)
	if err != nil {
		// TODO: Log the error
		fmt.Println("Error decoding the file")
	}

	flippedImage := imaging.FlipV(img)
	png_loader.cache[path] = flippedImage
	return flippedImage

}

// Walk through all the files in the directory and extract the png files
func getPngFiles() ([]string, error) {
	var pngFiles []string

	// Walk through the directory and get the path of all files
	err := filepath.Walk(png_loader.assetPath, func(path string, info os.FileInfo, err error) error {
		// Check if the file is a PNG file
		if filepath.Ext(path) == ".png" {
			pngFiles = append(pngFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return pngFiles, nil
}
