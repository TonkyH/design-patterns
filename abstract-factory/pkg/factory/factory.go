package factory

import "abstract-factoy/pkg/listfactory"

type IFactory interface {
	CreateLink(caption string, url string) *Link
	CreateTray(caption string) *Tray
	CreatePage(title string, author string) *Page
}

type Factory struct {
	IFactory
}

func GetFactory(className string) IFactory {
    switch className {
        case "ListFactory" {
            return &listfactory.NewListFa
        }
    }
}
