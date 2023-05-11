package renderer

import "github.com/micahke/mango/util/color"

// Just some statics to hold vertex rendering data

var quadVertices []float32 = []float32{

	// x, y, z,           r, g, b, a

	0.0, 0.0, 0.0, 0.0, 0.0, 0.0,
	0.0, 1.0, 0.0, 0.0, 0.0, 0.0,
	1.0, 1.0, 0.0, 0.0, 0.0, 0.0,
	1.0, 0.0, 0.0, 0.0, 0.0, 0.0,
}

var quadIndeces []uint32 = []uint32{
	0, 1, 2,
	2, 3, 0,
}

// Will generate quad vertex arrays that edit the color of the quad
func generateQuadVertices(color color.Color) []float32 {
	colorValues := color.Vec4

	// Make a copy of the quadVertices array and replaces the appropriate rgba values

	vertices := make([]float32, len(quadVertices))

	vertices[2] = colorValues[0]
	vertices[3] = colorValues[1]
	vertices[4] = colorValues[2]
	vertices[5] = colorValues[3]

	vertices[8] = colorValues[0]
	vertices[9] = colorValues[1]
	vertices[10] = colorValues[2]
	vertices[11] = colorValues[3]

	vertices[14] = colorValues[0]
	vertices[15] = colorValues[1]
	vertices[16] = colorValues[2]
	vertices[17] = colorValues[3]

	vertices[20] = colorValues[0]
	vertices[21] = colorValues[1]
	vertices[22] = colorValues[2]
	vertices[23] = colorValues[3]

	return vertices

}
