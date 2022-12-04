package concreteprototype

import (
	"fmt"
	"prototype/pkg/prototype"
)

type MessageBox struct {
    decochar string
}

func NewMessageBox(decochar string) *MessageBox {
    return &MessageBox{decochar: decochar}
}

func (m *MessageBox) Use(s string) {
    decolen := 1 + len(s) + 1
    for i := 0; i < decolen; i++ {
        fmt.Print(m.decochar)
    }
    fmt.Println()

    fmt.Println(m.decochar + s + m.decochar)

    for i := 0; i <decolen; i++ {
        fmt.Print(m.decochar)
    }
    fmt.Println()
}

func (m *MessageBox) CreateCopy() prototype.Product {
    c := *m 
    return &c
}
