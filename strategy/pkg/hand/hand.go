package hand

type Hand int64

const (
	ROCK = iota
	SCISSORS
	PAPER
)

func (h Hand) fight(hand Hand) int64 {
	if h == hand {
		return 0
	} else if (h+1)%3 == hand {
		return 1
	} else {
		return -1
	}
}

func (h *Hand) IsStrongThan(hand Hand) bool {
	return h.fight(hand) == 1
}

func (h *Hand) IsWeakerThan(hand Hand) bool {
	return h.fight(hand) == -1
}

func (h Hand) String() string {
	switch h {
	case ROCK:
		return "Rock"
	case SCISSORS:
		return "Scissors"
	case PAPER:
		return "Paper"
	default:
		return "Unknown"
	}
}
