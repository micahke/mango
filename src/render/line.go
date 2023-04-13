package render

import "github.com/micahke/infinite-universe/mango"


type Line struct {

  x0, y0, x1, y1 int

} 




func DrawLine(line Line, size int) {

  PlotLine(line.x0, line.y0, line.x1, line.y1, size)

}



// func PlotLine(x0, y0, x1, y1 int, size int) {
//
//     pts := []float32{}
//
//     dx := x1 - x0
//     dy := y1 - y0
//     yi := size
//     if dy < 0 {
//         yi = -size
//         dy = -dy
//     }
//     D := (2 * dy) - dx
//     y := y0
//
//     for x := x0; x <= x1; x += size {
//         pts = append(pts, float32(x), float32(y))
//         if D > 0 {
//             y = y + yi
//             D = D + (2 * (dy - dx))
//         } else {
//             D = D + 2*dy
//         }
//     }
//
//   mango.IM.DrawPixels(pts, float32(size))
// }


func plotLineLow(x0, y0, x1, y1, size int) {
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
  mango.IM.DrawPixels(pts, float32(size))
}

func plotLineHigh(x0, y0, x1, y1, size int) {
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
  mango.IM.DrawPixels(pts, float32(size))
}

func PlotLine(x0, y0, x1, y1, size int) {
    if abs(y1 - y0) < abs(x1 - x0) {
        if x0 > x1 {
            plotLineLow(x1, y1, x0, y0, size)
        } else {
            plotLineLow(x0, y0, x1, y1, size)
        }
    } else {
        if y0 > y1 {
            plotLineHigh(x1, y1, x0, y0, size)
        } else {
            plotLineHigh(x0, y0, x1, y1, size)
        }
    }
}


func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
