package ecs

import "testing"


type ecs_test struct {
  ecs *ECS
}


func TestECS(t *testing.T) {

  test := ecs_test{}
  test.setup()

  // Test to see whether adding an entity works
  t1 := test.addEntity("test")
  if !t1 {
    t.Fatalf("Could not add entity to ECS")
  }

  // Get the added entity
  t2 := test.getEntity("test")
  if !t2 {
    t.Fatalf("Could not get entity from ECS")
  }

}



func (test ecs_test) setup() {

  // Create new ECS
  test.ecs = new(ECS)
  
}


func (test ecs_test) addEntity(id string) bool {

  entity := test.ecs.CreateEntity(id)

  if entity != nil {
    return true
  }

  return false

}


// Check to see whether we get an entity that we know exists
func (test ecs_test) getEntity(id string) bool {

  entity := test.ecs.GetEntity(id)
  if entity != nil {
    return true
  }


  return false

}


