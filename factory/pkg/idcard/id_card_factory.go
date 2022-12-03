package idcard

import (
	"factory/pkg/framework"
	"fmt"
)

type IdCardFactory struct {
	*framework.Factory
}

func NewIdCardFactory() *IdCardFactory {
	IdCardFactory := &IdCardFactory{
		&framework.Factory{},
	}

	IdCardFactory.IFactory = IdCardFactory
	return IdCardFactory
}

func (i *IdCardFactory) CreateProduct(owner string) framework.Product {
	return NewIdCard(owner)
}

func (i *IdCardFactory) RegisterProduct() {
	fmt.Println("登録しました")
}
