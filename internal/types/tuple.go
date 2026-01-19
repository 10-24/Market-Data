package types

type Tuple[A, B any] struct {
	A A
	B B
}

func Tup[A, B any](a A, b B) Tuple[A, B] {
	return Tuple[A, B]{
		A: a,
		B: b,
	}
}
