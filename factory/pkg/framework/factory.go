package framework

type IFactory interface {
	CreateProduct(owner string) Product
	RegisterProduct()
}

type Factory struct {
	IFactory
}

func (f *Factory) Create(owner string) Product {
	p := f.CreateProduct(owner)
	f.RegisterProduct()
	return p
}
