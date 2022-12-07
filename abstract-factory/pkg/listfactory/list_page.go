package listfactory

import (
	"abstract-factoy/pkg/factory"
	"strings"
)

type ListPage struct {
	factory.Page
}

func NewListPage(title string, author string) *factory.Page {
	return factory.NewPage(title, author)
}

func (l *ListPage) MakeHTML() string {
	var sb strings.Builder
	sb.WriteString("<!DOCUMENTTYPE html> \n")
	sb.WriteString("<html><head><title>\n")
	sb.WriteString(l.Title)
	sb.WriteString("</title></head>\n")
	sb.WriteString("<body>\n")
	sb.WriteString("<h1>")
	sb.WriteString(l.Title)
	sb.WriteString("</h1>\n")
	sb.WriteString("<ul>\n")

	for i := 0; i < len(l.Content); i++ {
		sb.WriteString(l.Content[i].MakeHTML())
	}
	sb.WriteString("</ul>\n")
	sb.WriteString("<hr><address>")
	sb.WriteString(l.Author)
	sb.WriteString("</address>\n")
	sb.WriteString("</body></html>\n")

	return sb.String()
}
