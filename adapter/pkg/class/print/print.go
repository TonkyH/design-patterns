package print

type Print interface {
	PrintWeak() string
	PrintStrong() string
}
