package factory

type IItem interface {
	MakeHTML() string
}

type Item struct {
	IItem
	Caption string
}

func NewItem(caption string) *Item {
	return &Item{Caption: caption}
}
