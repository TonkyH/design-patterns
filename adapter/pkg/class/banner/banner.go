package banner

import "fmt"

type Banner struct {
	Text string
}

func NewBanner(text string) *Banner {
	return &Banner{Text: text}
}

func (b *Banner) ShowWithPaten() {
	fmt.Println("(" + b.Text + ")")
}

func (b *Banner) ShowWithSter() {
	fmt.Println("*" + b.Text + "*")
}
