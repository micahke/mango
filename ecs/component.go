package ecs


type Component interface {

  // This function initializes the component
  Init()

  // This function updated every frame
  Update()

}