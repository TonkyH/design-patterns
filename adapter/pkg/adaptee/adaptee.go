package adaptee

import "fmt"

type Adaptee struct {
	Text string
}

func NewAdaptee(text string) *Adaptee {
	return &Adaptee{Text: text}
}

func (b *Adaptee) ShowWithPaten() {
	fmt.Println("(" + b.Text + ")")
}

func (b *Adaptee) ShowWithSter() {
	fmt.Println("*" + b.Text + "*")
}
