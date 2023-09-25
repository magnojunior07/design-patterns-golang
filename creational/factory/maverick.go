package main

type Maverick struct {
	Gun
}

func newMaverick() IGun {
	return &Maverick{
		Gun: Gun{
			name:  "Maverick",
			power: 5,
		},
	}
}