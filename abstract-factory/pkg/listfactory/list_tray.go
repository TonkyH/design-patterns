package listfactory

import (
	"abstract-factoy/pkg/factory"
	"strings"
)

type ListTray struct {
	factory.Tray
}

func NewListTray(caption string) *factory.Tray {
	return factory.NewTray(caption)
}

func (l *ListTray) MakeHTML() string {
	var sb strings.Builder
	sb.WriteString("<li>\n")
	sb.WriteString(l.Caption)
	sb.WriteString("\n<ul>\n")

	for i := 0; i < len(l.ItemTray); i++ {
		sb.WriteString(l.ItemTray[i].MakeHTML())
	}

	sb.WriteString("</ul>\n")
	sb.WriteString("</li>\n")

	return sb.String()
}
