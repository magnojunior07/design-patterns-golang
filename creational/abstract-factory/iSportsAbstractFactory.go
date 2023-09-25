package main

import (
	"errors"
	"fmt"
)

type ISportsAbstractFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsAbstractFactory, error) {
	switch brand {
	case "adidas":
		return &AdidasFactory{}, nil
	case "nike":
		return &NikeFactory{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Wrong brand type passed"))
	}
}