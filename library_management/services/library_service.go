package services 

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(id int, book models.Book) error
	AddMember(member models.Member) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}
func NewLibrary() *Library{
	return &Library{
		Books: make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(id int, book models.Book) error {
	if _, exists := l.Books[id]; exists {
		return errors.New("book with this ID already exists")
	}
	l.Books[id] = book
	return nil
}

func (l *Library) AddMember(member models.Member) error{
	if _, exists := l.Members[member.ID]; exists {
		return errors.New("member with this ID already exists")
	}
	l.Members[member.ID] = member
	return nil
}
func (l *Library)RemoveBook (bookID int) error{
	if _,exists := l.Books[bookID];exists{
		delete(l.Books, bookID)
		return nil
	}
	return errors.New("book not found")
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	if book.Status == "Borrowed"{
		return errors.New("book is already borrowed")
	}
	member , exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks , book)
	l.Members[memberID] = member
	l.Books[bookID] = book

	return nil
}

func (l *Library)ReturnBook(bookID int , memberID int) error{
	book , exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	if book.Status == "Available"{
		return errors.New("book is already available")
	}

	member , exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}
	book.Status = "Available"
	member.BorrowedBooks = removeByValue(member.BorrowedBooks, book)
	l.Members[memberID] = member
	l.Books[bookID] = book
	return nil
}

func removeByValue(arr []models.Book, book models.Book) []models.Book {
	for i, v := range arr {
		if v.ID == book.ID { // Compare by ID only
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
func (l *Library)ListAvailableBooks() []models.Book{
	books := []models.Book{}
	for _, book := range l.Books{
		if book.Status == "Available"{
			books = append(books , book)
		}
	}

	return books
}

func (l *Library ) ListBorrowedBooks(memberID int) []models.Book{
	member , exists := l.Members[memberID]
	if !exists {
		return nil
	}

	return member.BorrowedBooks
}