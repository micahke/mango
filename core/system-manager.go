package core

import (
	"fmt"
	"reflect"
)

type MangoSystem interface {
	InitializeSystem()

	UpdateSystem()
}

type SystemManager struct {
	systems []MangoSystem
}

func NewSystemManager() *SystemManager {
	return &SystemManager{}
}

func (manager *SystemManager) Update() {
	for _, system := range manager.systems {
		system.UpdateSystem()
	}

}

func (manager *SystemManager) AddSystem(system MangoSystem) bool {
	if !manager.ContainsSystem(system) {
		system.InitializeSystem()
		manager.systems = append(manager.systems, system)
		return true
	}
	return false
}

func (manager *SystemManager) GetSystemOfType(t reflect.Type) (MangoSystem, error) {
	for _, s := range manager.systems {
		if reflect.TypeOf(s) == t {
			return s, nil
		}
	}
	return nil, fmt.Errorf("No system found")
}

func (manager *SystemManager) GetSystem(system MangoSystem) (MangoSystem, error) {
	for _, s := range manager.systems {
		if s == system {
			return system, nil
		}
	}
	return nil, fmt.Errorf("No system found")
}

func (manager *SystemManager) ContainsSystemOfType(t reflect.Type) bool {
	for _, s := range manager.systems {
		if reflect.TypeOf(s) == t {
			return true
		}
	}
	return false
}

func (manager *SystemManager) ContainsSystem(system MangoSystem) bool {
	for _, s := range manager.systems {
		if s == system {
			return true
		}
	}
	return false
}
