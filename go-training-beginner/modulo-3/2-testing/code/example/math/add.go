package math

func Add(a, b int) int {
	return a + b
}

func AddMore(a, b int) int {
	if a > b+10 {
		panic("B must be greater than A")
	}

	return a + b
}
