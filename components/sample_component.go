package components


// A sample component that we can use around the engine for various tasks
type SampleComponent struct {
  
}



func (component *SampleComponent) Init() {}


func (component *SampleComponent) Update() {}


func (component *SampleComponent) GetComponentName() string {
  return "Sample Component"
}
