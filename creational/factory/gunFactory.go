package main

import "fmt"

type GunFactory struct {}

func (g *GunFactory) makeGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "maverick":
		return newMaverick(), nil
	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}