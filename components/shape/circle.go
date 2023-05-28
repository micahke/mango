package shape



type Circle struct {
  Radius float32
}


func NewCircle(radius float32) *Circle {
  return &Circle{
    Radius: radius,
  }
}


func (circle *Circle) GetShapeName() string {
  return "Circle"
}
