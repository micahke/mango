package shape


type Rect struct {

  Width float32
  Height float32

}


func (rect *Rect) GetShapeName() string {
  return "Rect"
}
