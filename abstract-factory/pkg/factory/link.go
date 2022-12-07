package factory

type Link struct {
	*Item
	Url string
}

func NewLink(caption string, url string) *Link {
	return &Link{
		Item: NewItem(caption),
		Url:  url,
	}
}
