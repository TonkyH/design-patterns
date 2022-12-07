package listfactory

import "abstract-factoy/pkg/factory"

type ListLink struct {
	factory.Link
}

func NewListLink(caption string, url string) *factory.Link {
	return factory.NewLink(caption, url)
}

func (l *ListLink) MakeHTML() string {
	return "<li> <a href=\"" + l.Url + "\">" + l.Caption + "</a></li>\n"
}
