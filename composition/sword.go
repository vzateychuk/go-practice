package main

import "fmt"

type Sword struct {
	Name string
}

func (Sword) Damage() int {
	return 2 // Damage returns the damage dealt by this sword.
}

func (s Sword) String() string {
	return fmt.Sprintf("%s is a sword that can deal %d points of damage", s.Name, s.Damage())
}
