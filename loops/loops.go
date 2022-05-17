package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Sum =", sum)

	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement. Sum is", sum)
	}

	//--Summary:
	//  Implement the classic "FizzBuzz" problem using a `for` loop.
	//
	//--Requirements:
	//* Print integers 1 to 50, except:
	//  - Print "Fizz" if the integer is divisible by 3
	//  - Print "Buzz" if the integer is divisible by 5
	//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
	//
	//--Notes:
	//* The remainder operator (%) can be used to determine divisibility

	for i := 1; i <= 50; i++ {

		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		if divisibleBy3 && divisibleBy5 {
			fmt.Println("FizzBuzz", i)
		} else if divisibleBy3 {
			fmt.Println("Fizz", i)
		} else if divisibleBy5 {
			fmt.Println("Buzz", i)
		} else {
			fmt.Println(i)
		}
	}
}
