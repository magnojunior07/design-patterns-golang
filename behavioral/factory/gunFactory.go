package main

import "fmt"

type GunFactory struct {}

func (g *GunFactory) makeGun(gunType string) IGun {
	switch gunType {
	case "ak47":
		newAk47()
	case "maverick":
		newMaverick()
	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}