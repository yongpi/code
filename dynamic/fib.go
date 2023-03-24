package dynamic

func Fib(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

func Fib2(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	first, second := 1, 1
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}

func Stairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return Stairs(n-1) + Stairs(n-2)
}

func Cow(n int) int {
	if n <= 0 {
		return 0
	}

	if n <= 3 {
		return n
	}

	return Cow(n-1) + Cow(n-3)
}
