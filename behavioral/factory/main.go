package main

func main() {
	gunFactory := GunFactory{}
	ak47, _ := gunFactory.makeGun("ak47")
	maverick, _ := gunFactory.makeGun("maverick")
	
	printGunDetails(ak47)
	printGunDetails(maverick)
}

func printGunDetails(g IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}