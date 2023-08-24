package main

type NikeShoe struct {
	Shoe
}

type NikeShirt struct {
	Shirt
}

type NikeFactory struct {}

func (n *NikeFactory) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe {
		logo: "nike",
		size: 14,
		},
		}
}

func (n *NikeFactory) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt {
			logo: "nike",
			size: "L",
		},
	}
}