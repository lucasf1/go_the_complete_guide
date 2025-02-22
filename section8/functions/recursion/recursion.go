package recursion

import "fmt"

func main() {
	fact := factorial_loop(5)
	fmt.Println(fact)

	fact = factorial_recur(5)
	fmt.Println(fact)
}

func factorial_loop(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}

func factorial_recur(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial_recur(number-1)
}
