package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

func main() {

	data := []rune{'a', 'b', 'c', 'd'}

	var capitalized []rune

	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))

		fmt.Printf("%c done!\n", r)
	}

	for i := 0; i < len(data); i++ {
		char := data[i]

		go capIt(char)
	}

	time.Sleep(500 * time.Microsecond)

	fmt.Printf("Capitalized: %c\n", capitalized)

	wait := func(ms time.Duration) {
		time.Sleep(ms * time.Millisecond)

		fmt.Println("sleep duration", ms*time.Millisecond)
	}

	fmt.Println("Launching goroutines")

	go wait(500)
	go wait(100)
	go wait(1000)

	fmt.Println("Launched goroutines")

	time.Sleep(1100 * time.Millisecond)

	fmt.Println("Waited 1100")

	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}

	sum := 0

	for _, file := range files {

		go getSumOfFile(file, &sum)

	}

	time.Sleep(5000 * time.Millisecond)

	fmt.Println("Overall files Sum", sum)

}

func getSumOfFile(filepath string, overallSum *int) (int, error) {
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("readFile", err)
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		textNum := scanner.Text()
		numInLine, err := strconv.Atoi(textNum)

		if err != nil {
			fmt.Println("numInLine err", err)
			continue
		}

		sum += numInLine
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner err", err)
	}

	fmt.Println("Sum of", filepath, "is", sum)

	*overallSum += sum

	return sum, err

}
