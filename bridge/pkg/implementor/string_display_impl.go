package implementor

import (
	"fmt"
	"strings"
)

type StringDisplay struct {
	Str   string
	Width int
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{
		Str:   str,
		Width: len(str),
	}
}

func (s *StringDisplay) RawOpen() {
	s.PrintLine()
}

func (s *StringDisplay) RawPrint() {
	fmt.Printf("|%s|\n", s.Str)
}

func (s *StringDisplay) RawClose() {
	s.PrintLine()
	fmt.Println("")
}

func (s *StringDisplay) PrintLine() {
	line := strings.Repeat("-", s.Width)
	fmt.Println(line)
}
