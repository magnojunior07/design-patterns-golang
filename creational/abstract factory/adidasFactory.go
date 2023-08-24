package main

type AdidasShoe struct {
	Shoe
}
type AdidasShirt struct {
	Shirt
}

type AdidasFactory struct {}

func (a *AdidasFactory) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe {
		logo: "adidas",
		size: 14,
		},
	}
}
func (a *AdidasFactory) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt {
			logo: "adidas",
			size: "XL",
		},
	}
}