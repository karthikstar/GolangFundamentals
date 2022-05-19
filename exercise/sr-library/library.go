//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change

//* There must only ever be one copy of the library in memory at any time - NEED USE POINTERS

//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

//type Aliases
type Title string // book title
type Name string  // library memmber name

type LendAudit struct { // track when book is checked in or checked out
	checkOut time.Time
	checkIn  time.Time // makes use of tiem package
	//if checkin = 0 we know it hasnt been returned yet
	// if both checkout and checkin have a time, we know the book has been returned
}

type Member struct {
	name  Name
	books map[Title]LendAudit // books is of map type, having keys - title, and values - lendAudit
}

//for tracking of no of books. we only have a finite no of books
type BookEntry struct {
	total  int // total books opwned by the library
	lended int // total books lended out
}

type Library struct {
	members map[Name]Member     // member type has the lend audit information. library will be able to determine who has what books
	books   map[Title]BookEntry // can tell total no of books owned and lent out
}

// audit for 1 member
func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var ReturnTime string
		if audit.checkIn.IsZero() { // if its 0, means hasnt been returned yet
			ReturnTime = "[not returned yet]"
		} else {
			ReturnTime = audit.checkIn.String() // will display info in a string format
		}
		fmt.Println(member.name, title, ":", audit.checkOut.String(), "through", ReturnTime)
	}
}

// audit for ALL The members
func printLibraryAudit(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

// prints the books out
func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/lended:", book.lended) // track which books have been lent out, and how many books the library has
	}
	fmt.Println()
}

// checkout a book
// req.s library, title of book, and member who is checking out the book
// bool will indicate if the checkout is succesful
func checkoutBook(library *Library, title Title, member *Member) bool {
	// check if title of book exists in library
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	// if book exists, we need check if there is any avail copies of the book is available
	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}
	//there is available copy of the book to lend
	book.lended += 1 // update

	library.books[title] = book //reassign the book back to its corresp title as we were working on a copy just now

	// update member lending information
	member.books[title] = LendAudit{checkOut: time.Now()} // checkIn time will have a default value of 0
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	audit, found := member.books[title]
	if !found {
		fmt.Println("member did not check out this book")
		return false
	}
	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit // reassin audit

	return true
}
func main() {
	//	lib := Library{{"Tom's land", "Mr Cook", "Baby says hi", "Peter pan"}, {"John", "Alice", "Peter"}}
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	// creating books
	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The Little Go Book"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Let's learn go"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Go Bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}

	//adding some members for the library
	library.members["Jayson"] = Member{"Jayson", make(map[Title]LendAudit)}
	library.members["Billy"] = Member{"Billy", make(map[Title]LendAudit)}
	library.members["Susan"] = Member{"Susan", make(map[Title]LendAudit)}

	fmt.Println("\nInitial")
	printLibraryBooks(&library)
	printLibraryAudit(&library)

	member := library.members["Jayson"]
	checkedOut := checkoutBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck out a book:")

	if checkedOut {
		printLibraryBooks(&library)
		printLibraryAudit(&library)
	}

	returned := returnBook(&library, "Go Bootcamp", &member)
	fmt.Println("\nCheck in a book:")
	if returned {
		printLibraryBooks(&library)
		printLibraryAudit(&library)
	}

	// initially we shld have some books , some members 0 lent out
	// then we let one member check out a book
	// then later we get the member to return it
}
