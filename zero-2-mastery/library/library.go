package main

import (
	"fmt"
	"time"
)

type Book string

type Member string

type CheckedOutTime time.Time

type ReturnedTime time.Time

type Issue struct {
	book       Book
	member     Member
	checkedOut CheckedOutTime
	returned   ReturnedTime
}

type Library struct {
	books   []Book
	members []Member
	issues  map[Book]Issue
}

func checkOut(book Book, member Member, library *Library) {
	library.issues[book] = Issue{book: book, member: member, checkedOut: CheckedOutTime(time.Now())}
}

func returnBook(book Book, member Member, library *Library) {

	issue, found := library.issues[book]

	if found {
		issue.returned = ReturnedTime(time.Now())
		library.issues[book] = issue
	} else {
		fmt.Println("Book could not found!")
	}
}

func main() {
	books := []Book{Book("Ice And Fire"), Book("Bhaukaal"), Book("Mahakaal"), Book("Amrit")}
	members := []Member{Member("CJS"), Member("Rupesh"), Member("Singh")}

	library := Library{books: books, members: members, issues: make(map[Book]Issue)}
	fmt.Println(library)

	checkOut(books[1], members[1], &library)
	fmt.Println(library.issues)

	returnBook(books[1], members[1], &library)
	fmt.Println(library.issues)

	checkOut(books[0], members[0], &library)
	fmt.Println(library.issues)

	returnBook(books[0], members[0], &library)
	fmt.Println(library.issues)

}
