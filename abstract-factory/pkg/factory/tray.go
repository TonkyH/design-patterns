package factory

type Tray struct {
	*Item
	ItemTray []Item
}

func NewTray(caption string) *Tray {
	return &Tray{
		Item:     NewItem(caption),
		ItemTray: []Item{},
	}
}

func (t *Tray) Add(item Item) {
	t.ItemTray = append(t.ItemTray, item)
}
