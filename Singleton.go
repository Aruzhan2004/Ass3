package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	value int
}

var (
	once     sync.Once
	instance *Singleton
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func (s *Singleton) DoSomething() {
	fmt.Printf("Doing something with value %d\n", s.value)
}

func main() {
	object := GetInstance()
	object.value = 1000

	secondObject := GetInstance()
	secondObject.value = 2000

	object.DoSomething()
	secondObject.DoSomething()
}
