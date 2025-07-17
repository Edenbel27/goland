package controllers

import(
	"fmt"
	"bufio"
    "os"
	"library_management/services"
	"library_management/models"
)

func DisplayMenu(lib *services.Library){
	i := true
	for i{
	fmt.Print(`
		1. Add Book 
		2. Add Member 
		3. Remove Book 
		4. Borrow Book 
		5. Return Book 
		6. List Available Books  
		7. List Borrowed Books 
		8. Exit 

	`)
		fmt.Print("Enter your choice : ")
		var choice int
		fmt.Scan(& choice)
		fmt.Scanln()
		switch choice{
			case 1:
				addBook(lib)
			case 2:
				addMember(lib)
			case 3:
				removeBook(lib)
			case 4:
				borrowBook(lib)
			case 5:
				returnBook(lib)
			case 6:
				listAvailableBooks(lib)
			case 7:
				listBorrowedBooks(lib)
			case 8:
				fmt.Println("Exiting... Goodbye!")
				i = false
				return
			default:
				fmt.Print("Invalid choice. Try again.")

		}

	}

	
}
	func addBook(lib *services.Library) {

	fmt.Print("Enter Book Id : ")
	var id int
	fmt.Scanf("%d", &id)
	fmt.Scanln()

	fmt.Print("Enter Book Title : ")
	var title string
	fmt.Scanf("%s", &title)
	fmt.Scanln()

	fmt.Print("Enter Book Author : ")
	var author string
	fmt.Scanf("%s", &author)
	fmt.Scanln()

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	err := lib.AddBook(id, book)
	if err != nil {
		fmt.Println("Error adding book:", err)
	} else {
		fmt.Println("Book added successfully.")
	}
}

func addMember(lib *services.Library){
	var id int
	fmt.Print("Enter Member ID : ")
	fmt.Scanf("%d", &id)
	fmt.Scanln()
	
	fmt.Print("Enter Member Name : ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]
	member := models.Member{
		ID:           id,
		Name:         name,
		BorrowedBooks: []models.Book{},
	}
	err := lib.AddMember(member)
	if err != nil {
		fmt.Println("Error adding member:", err)
	} else {
		fmt.Println("Member added successfully.")
	}
}
func removeBook(lib *services.Library){
	fmt.Print("Enter the id of the book : ")
	var id int
	fmt.Scanf("%d", &id)
	err := lib.RemoveBook(id)
	if err != nil {
		fmt.Println("Error removing book:", err)
	} else {
		fmt.Println("Book removed successfully.")
	}
}

func borrowBook(lib *services.Library){
	fmt.Print("Enter the id of the book to borrow : ")
	var bookID int
	fmt.Scanf("%d", &bookID)
	fmt.Scanln()
	fmt.Print("Enter your member ID : ")
	var memberID int
	fmt.Scanf("%d", &memberID)
	fmt.Scanln()
	err := lib.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}

}

func returnBook(lib *services.Library){
	fmt.Print("Enter the id of the book to return : ")
	var bookID int
	fmt.Scanf("%d", &bookID)
	fmt.Scanln()
	fmt.Print("Enter your member ID : ")
	var memberID int
	fmt.Scanf("%d", &memberID)
	fmt.Scanln()
	err := lib.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error returning book:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func listAvailableBooks(lib *services.Library){
	availableBooks := lib.ListAvailableBooks()
	for _, book := range availableBooks{
		fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID, book.Title, book.Author, book.Status)
	}
}

func listBorrowedBooks(lib *services.Library){
	fmt.Print("Enter your member ID to list borrowed books: ")
	var memberID int
	fmt.Scanf("%d", &memberID)
	borrowedBooks := lib.Members[memberID].BorrowedBooks
	if borrowedBooks == nil {
		fmt.Println("No books borrowed by this member.")
		return
	}
	for _, book := range borrowedBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID, book.Title, book.Author, book.Status)
	}
}