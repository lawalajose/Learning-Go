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
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Name string
type Title string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func prinLibraryBook(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkOutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not found in library")
		return false
	}
	if book.lended == book.total {
		fmt.Println("No more book available to lend")
		return false
	}
	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of the library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not checkout this book")
		return false
	}
	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		members: make(map[Name]Member),
		books:   make(map[Title]BookEntry),
	}

	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The little Go book"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Let's learn Go"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Go bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.members["Jayson"] = Member{"Jayson", make(map[Title]LendAudit)}
	library.members["Billy"] = Member{"Billy", make(map[Title]LendAudit)}
	library.members["Nunez"] = Member{"Nunez", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	prinLibraryBook(&library)
	printMemberAudits(&library)

	member := library.members["Jayson"]
	checkOut := checkOutBook(&library, "Go bootcamp", &member)
	fmt.Println("\nCheck out a book!")
	if checkOut {
		prinLibraryBook(&library)
		printMemberAudits(&library)
	}

	returned := returnBook(&library, "Go bootcamp", &member)
	fmt.Println("\nCheck in a book!")
	if returned {
		prinLibraryBook(&library)
		printMemberAudits(&library)
	}

}

// type Member struct {
// 	Name string
// }

// type Library struct {
// 	Members []Member
// 	Books   map[string]int
// }

// func printLibrary(Library Library) {
// 	mem := Library.Members
// 	fmt.Println("___ Dev Library Members ___")
// 	for _, members := range mem {
// 		fmt.Println(members)
// 	}
// 	fmt.Println()
// 	book := Library.Books
// 	for key, value := range book {
// 		fmt.Println("Ref No:", value, "Name : ", key)
// 	}
// }

// func checkout(library *Library, bookName string, RefNo int) {
// 	book := library.Books
// 	delete(book, bookName)
// 	checkOut := make(map[string]int)

// }

// func main() {
// 	books := make(map[string]int)

// 	members := []Member{
// 		Member{"Lawal Ajose"},
// 		Member{"Bashir AbdulSalam"},
// 		Member{"Yusuf Idris"},
// 	}
// 	books["Java"] = 1
// 	books["Golang"] = 2
// 	books["Javascript"] = 3
// 	books["Rust"] = 4

// 	devLibrary := Library{
// 		Members: members,
// 		Books:   books,
// 	}
// 	printLibrary(devLibrary)
// }
