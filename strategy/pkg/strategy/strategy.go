package strategy

import (
	"strategy/pkg/hand"
)

type IStrategy interface {
	NextHand() hand.Hand
	study()
}

type WinningStrategy struct {
	random   int64
	won      bool
	prevHand hand.Hand
}
