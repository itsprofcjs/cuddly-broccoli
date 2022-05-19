package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1

	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new

	increment(counter)
}

const (
	Active   = 1
	Inactive = 0
)

type ItemState struct {
	name  string
	state int
}

func activate(item *ItemState) {
	item.state = Active
}

func deactivate(item *ItemState) {
	item.state = Inactive
}

func checkout(items []ItemState) {
	for idx := range items {
		deactivate(&items[idx])
	}
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)

	items := []ItemState{
		{name: "First Item", state: Active},
		{name: "Second Item", state: Active},
		{name: "Third Item", state: Active},
		{name: "Fourth Item", state: Active},
	}
	fmt.Println("Items", items)

	deactivate(&items[3])
	fmt.Println("Items deactivate", items)

	checkout(items)
	fmt.Println("Checkout", items)
}
