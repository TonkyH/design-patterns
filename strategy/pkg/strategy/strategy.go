package strategy

import (
	"math/rand"
	"strategy/pkg/hand"
)

type IStrategy interface {
	NextHand() hand.Hand
	Study()
}

type WinningStrategy struct {
	random   int64
	won      bool
	prevHand hand.Hand
}

func NewWinningStrategy(seed int64) *WinningStrategy {
	rand.Seed(seed)
	return &WinningStrategy{
		random: rand.Int63n(3),
		won:    false,
	}
}

func (s WinningStrategy) NextHand()  hand.Hand {
    if (!s.won) {
        s.prevHand = hand.GetHand(rand.Int63n(3))
    }
    return s.prevHand
}

func (s WinningStrategy) Study(win bool) {
    s.won = win
}
