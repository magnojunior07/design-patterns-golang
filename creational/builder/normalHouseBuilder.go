package main

type NormalHouseBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalHouseBuilder() *NormalHouseBuilder {
	return &NormalHouseBuilder{}
}

func (b *NormalHouseBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *NormalHouseBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *NormalHouseBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalHouseBuilder) getHouse() House {
	return House {
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}