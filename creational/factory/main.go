package main

func main() {
	gunFactory := GunFactory{}

	ak47, errAk47 := gunFactory.makeGun("ak47")
	if errAk47 != nil {
		fmt.Println(errAk47)
	}
	
	maverick, errMaverick := gunFactory.makeGun("maverick")
	if errMaverick != nil {
		fmt.Println(errMaverick)
	}

	printGunDetails(ak47)
	printGunDetails(maverick)
}

func printGunDetails(g IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}