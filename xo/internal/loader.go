package internal

import "fmt"

// The loader interface
type ILoader interface {
	SayHello()
}

var AllLoaders = map[LoaderType]ILoader{}

// The loader implementation
// Drivers like mysql will create object for this.
type LoaderImp struct {
	Message string
}

func (lt *LoaderImp) SayHello() {
	fmt.Println("Hey from loader", lt.Message)
}
