package main

import (
	"errors"
	"fmt"
)

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuider(builderType string) (IBuilder, error) {
	switch builderType {
	case "normal":
		return &NormalHouseBuilder{}, nil
	case "igloo":
		return &IglooHouseBuilder{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Not supported builder type"))		
	}
}