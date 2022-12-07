package listfactory

import (
	"abstract-factoy/pkg/factory"
)

type ListFactory struct {
}

func (l *ListFactory) CreateLink(caption string, url string) *factory.Link {
	return NewListLink(caption, url)
}

func (l *ListFactory) CreateTray(caption string) *factory.Tray {
	return NewListTray(caption)
}

func (l *ListFactory) CreatePage(title string, author string) *factory.Factory {
	return NewListPage(title, author)
}
