package idcard

import "fmt"

type IdCard struct {
	owner string
}

func NewIdCard(owner string) *IdCard {
	fmt.Printf("%sのカードを作成します\n", owner)
	return &IdCard{owner: owner}
}

func (i *IdCard) Use() {
	fmt.Printf("%s を使用します\n", i.String())
}

func (i *IdCard) String() string {
	return "[IDCard: " + i.owner + "]"
}
