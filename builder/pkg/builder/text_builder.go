package builder

import "strings"

type TextBuilder struct {
	sb strings.Builder
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{sb: strings.Builder{}}
}

func (t *TextBuilder) MakeTitle(title string) {
	t.sb.WriteString("========================")
	t.sb.WriteString("「")
	t.sb.WriteString(title)
	t.sb.WriteString("」\n\n")
}

func (t *TextBuilder) MakeString(str string) {
	t.sb.WriteString("*")
	t.sb.WriteString(str)
	t.sb.WriteString("\n\n")
}

func (t *TextBuilder) MakeItems(items []string) {
	for _, s := range items {
		t.sb.WriteString(" ・")
		t.sb.WriteString(s)
		t.sb.WriteString("\n")
	}
	t.sb.WriteString("\n")
}

func (t *TextBuilder) Close() {
	t.sb.WriteString("========================\n")
}

func (t *TextBuilder) GetTextResult() string {
    return t.sb.String()
}
