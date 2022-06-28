package main

func Fibonacci(num int) int {
	if num <= 2 {
		return 1
	}
	first, second, result := 1, 1, 1
	for i := 2; i < num; i++ {
		result, first, second = first+second, second, first+second
	}
	return result
}
