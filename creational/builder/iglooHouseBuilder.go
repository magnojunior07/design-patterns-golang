package main

type IglooHouseBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooHouseBuilder() *IglooHouseBuilder {
	return &IglooHouseBuilder{}
}

func (b *IglooHouseBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooHouseBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooHouseBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooHouseBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}