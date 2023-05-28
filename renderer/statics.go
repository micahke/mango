package renderer

import "github.com/micahke/mango/util/color"

// Just some statics to hold vertex rendering data

var quadVertices []float32 = []float32{

	// x, y           r, g, b, a        // uvX, uvY

	0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0,
	0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0,
	1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 1.0,
	1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0,
}

var quadIndeces []uint32 = []uint32{
	0, 1, 2,
	2, 3, 0,
}

// Will generate quad vertex arrays that edit the color of the quad
func generateQuadVertices(color color.Color) []float32 {
	colorValues := color.Vec4

	// Make a copy of the quadVertices array and replaces the appropriate rgba values

	vertices := quadVertices

	vertices[2] = colorValues[0]
	vertices[3] = colorValues[1]
	vertices[4] = colorValues[2]
	vertices[5] = colorValues[3]

  vertices[10] = colorValues[0]
  vertices[11] = colorValues[1]
  vertices[12] = colorValues[2]
  vertices[13] = colorValues[3]

  vertices[18] = colorValues[0]
  vertices[19] = colorValues[1]
  vertices[20] = colorValues[2]
  vertices[21] = colorValues[3]

  vertices[26] = colorValues[0]
  vertices[27] = colorValues[1]
  vertices[28] = colorValues[2]
  vertices[29] = colorValues[3]

	return vertices

}
