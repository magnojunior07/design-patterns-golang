package main

import "fmt"

func main() {
	normalHouseBuilder, errNormalHouse := getBuider("normal")
	if(errNormalHouse != nil) {
		fmt.Println(errNormalHouse)
		return
	}

	iglooHouseBuilder, errIglooHouse := getBuider("igloo")
	if(errIglooHouse != nil) {
		fmt.Println(errIglooHouse)
		return
	}

	fmt.Println("Builder pattern example")
	fmt.Println("========================")

	director := newDirector(normalHouseBuilder)
	normalHouse := director.buildHouse()

	fmt.Println("Normal House:")
	fmt.Printf("Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Num Floor: %d\n", normalHouse.floor)

	fmt.Println(" ")

	director.setBuilder(iglooHouseBuilder)
	iglooHouse := director.buildHouse()

	fmt.Println("Igloo House:")
	fmt.Printf("Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Num Floor: %d\n", iglooHouse.floor)

}