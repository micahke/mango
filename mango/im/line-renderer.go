package im

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"github.com/micahke/infinite-universe/mango/opengl"
	"github.com/micahke/infinite-universe/mango/util/color"
)

type LineRenderer struct {
	vao    *opengl.VertexArray
	vbo    *opengl.VertexBuffer
	layout *opengl.VertexBufferLayout

	shader      *opengl.Shader
	modelMatrix glm.Mat4
}

var points []float32 = []float32{
	0.0, 0.0,
	0.0, 0.0,
}

func InitLineRenderer() *LineRenderer {

	renderer := new(LineRenderer)

	renderer.vao = opengl.NewVertexArray()
	renderer.vbo = opengl.NewVertexBuffer(points)
	renderer.layout = opengl.NewVertexBufferLayout()
	renderer.layout.Pushf(2)

	renderer.vao.AddBuffer(*renderer.vbo, *renderer.layout)

	// renderer.shader = opengl.NewShaderG("LineVertex.glsl", "LineFragment.glsl", "LineGeometry.glsl")
	renderer.shader = opengl.NewShader("LineVertex.glsl", "LineFragment.glsl")

	renderer.modelMatrix = glm.Ident4()

	return renderer

}

func (renderer *LineRenderer) RenderLine(p1, p2 glm.Vec2, color color.Color, thickness float32, projectionMatrix, viewMatrix glm.Mat4) {

	renderer.vbo.SetData([]float32{
		p1.X(), p1.Y(),
		p2.X(), p2.Y(),
	})

	renderer.shader.Bind()
	renderer.shader.SetUniformMat4f("projection", projectionMatrix)
	renderer.shader.SetUniformMat4f("view", viewMatrix)
	renderer.shader.SetUniformMat4f("model", renderer.modelMatrix)

	tk := thickness / float32(1300)

	renderer.shader.SetUniform1f("thickness", tk)

	renderer.shader.SetUniform4f("uColor", color.Vec4[0], color.Vec4[1], color.Vec4[2], color.Vec4[3])

	renderer.vao.Bind()

	gl.DrawArrays(gl.LINES, 0, 2)

}



func plotLineLow(x0, y0, x1, y1, size int) []float32 {
    pts := []float32{}
    dx := x1 - x0
    dy := y1 - y0
    yi := size
    if dy < 0 {
        yi = -size
        dy = -dy
    }
    D := (2 * dy) - dx
    y := y0

    for x := x0; x <= x1; x += size {
        pts = append(pts, float32(x), float32(y))
        if D > 0 {
            y = y + yi
            D = D + (2 * (dy - dx))
        } else {
            D = D + 2*dy
        }
    }
  return pts
}

func plotLineHigh(x0, y0, x1, y1, size int) []float32 {
    pts := []float32{}
    dx := x1 - x0
    dy := y1 - y0
    xi := size
    if dx < 0 {
        xi = -size
        dx = -dx
    }
    D := (2 * dx) - dy
    x := x0

    for y := y0; y <= y1; y += size {
        pts = append(pts, float32(x), float32(y))
        if D > 0 {
            x = x + xi
            D = D + (2 * (dx - dy))
        } else {
            D = D + 2*dx
        }
    }
  return pts
}

func (renderer *LineRenderer) GenerateBresenhamPoints(x0, y0, x1, y1, size int) []float32 {
    if abs(y1 - y0) < abs(x1 - x0) {
        if x0 > x1 {
            return plotLineLow(x1, y1, x0, y0, size)
        } else {
            return plotLineLow(x0, y0, x1, y1, size)
        }
    } else {
        if y0 > y1 {
            return plotLineHigh(x1, y1, x0, y0, size)
        } else {
            return plotLineHigh(x0, y0, x1, y1, size)
        }
    }
}


func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
