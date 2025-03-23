package tasklist

type Adder struct {
	a int
	b int
}

func NewAdder(a, b int) *Adder {
	return &Adder{
		a: a,
		b: b,
	}
}

func (adder *Adder) Sum() int {
	return adder.a + adder.b
}
