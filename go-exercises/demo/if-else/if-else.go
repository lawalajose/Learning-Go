package main

import (
	"fmt"
	"time"
)

// Book represents a book in the library
type Book struct {
	ID         int
	Title      string
	Author     string
	IsBorrowed bool
}

// Member represents a library member
type Member struct {
	ID   int
	Name string
}

// BorrowRecord represents a borrowing record
type BorrowRecord struct {
	ID         int
	BookID     int
	MemberID   int
	BorrowDate time.Time
	DueDate    time.Time
}

// Library represents the library system
type Library struct {
	books         map[int]Book
	members       map[int]Member
	borrowRecords map[int]BorrowRecord
	nextBookID    int
	nextMemberID  int
	nextBorrowID  int
}

// NewLibrary creates a new library
func NewLibrary() *Library {
	return &Library{
		books:         make(map[int]Book),
		members:       make(map[int]Member),
		borrowRecords: make(map[int]BorrowRecord),
		nextBookID:    1,
		nextMemberID:  1,
		nextBorrowID:  1,
	}
}

// AddBook adds a new book to the library
func (l *Library) AddBook(title, author string) int {
	book := Book{
		ID:         l.nextBookID,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}
	l.books[l.nextBookID] = book
	l.nextBookID++
	return book.ID
}

// AddMember adds a new member to the library
func (l *Library) AddMember(name string) int {
	member := Member{
		ID:   l.nextMemberID,
		Name: name,
	}
	l.members[l.nextMemberID] = member
	l.nextMemberID++
	return member.ID
}

// BorrowBook allows a member to borrow a book
func (l *Library) BorrowBook(memberID, bookID int, dueDate time.Time) bool {
	book, exists := l.books[bookID]
	if !exists || book.IsBorrowed {
		return false
	}

	member, exists := l.members[memberID]
	if !exists {
		return false
	}

	record := BorrowRecord{
		ID:         l.nextBorrowID,
		BookID:     bookID,
		MemberID:   memberID,
		BorrowDate: time.Now(),
		DueDate:    dueDate,
	}
	l.borrowRecords[l.nextBorrowID] = record
	l.nextBorrowID++
	book.IsBorrowed = true
	l.books[bookID] = book

	fmt.Printf("Book '%s' borrowed by %s. Due on %s\n", book.Title, member.Name, dueDate.Format("2006-01-02"))
	return true
}

// ReturnBook allows a member to return a book
func (l *Library) ReturnBook(bookID int) bool {
	book, exists := l.books[bookID]
	if !exists || !book.IsBorrowed {
		return false
	}

	book.IsBorrowed = false
	l.books[bookID] = book

	fmt.Printf("Book '%s' returned.\n", book.Title)
	return true
}

func main() {
	library := NewLibrary()

	// Add some books
	book1ID := library.AddBook("The Go Programming Language", "Alan A. A. Donovan")
	book2ID := library.AddBook("Effective Go", "Robert Griesemer")

	// Add some members
	member1ID := library.AddMember("John Doe")
	member2ID := library.AddMember("Jane Smith")

	// Borrow a book
	dueDate := time.Now().AddDate(0, 0, 14) // Due in 14 days
	library.BorrowBook(member1ID, book1ID, dueDate)

	// Return a book
	library.ReturnBook(book1ID)
}
