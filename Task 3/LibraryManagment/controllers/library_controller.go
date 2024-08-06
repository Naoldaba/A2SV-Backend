package controllers

import (
    "bufio"
    "fmt"
    "library_managment/models"
    "library_managment/services"
    "os"
    "strconv"
)

func LibraryController(lib *services.Library) {

    reader := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("Welcome to A2SV lib Managment System")
        fmt.Println("1. Add a new book")
        fmt.Println("2. Remove an existing book")
        fmt.Println("3. Borrow a book")
        fmt.Println("4. Return a book")
        fmt.Println("5. List all available books")
        fmt.Println("6. List all borrowed books by a member")
        fmt.Println("7. Register as a member")
        fmt.Println("8. list all members")
        fmt.Println("9. Quit")

        fmt.Print("Enter your choice: ")
        reader.Scan()
        choice := reader.Text()

        switch choice {
            case "1":
                addBook(lib, reader)
            case "2":
                removeBook(lib, reader)
            case "3":
                borrowBook(lib, reader)
            case "4":
                returnBook(lib, reader)
            case "5":
                listAvailableBooks(lib)
            case "6":
                listBorrowedBooks(lib, reader)
            case "7":
                registerUser(lib, reader)
            case "8":
                listMembers(lib)
            case "9":
                return 
            default:
                fmt.Println()
                fmt.Println("Invalid choice, please try again.")
                fmt.Println()
        }
    }
}

func addBook(lib *services.Library, reader *bufio.Scanner) {

    id := lib.GetNextUniqueBookID() 

    fmt.Print("Enter book Title: ")
    reader.Scan()
    title := reader.Text()

    fmt.Print("Enter book Author: ")
    reader.Scan()
    author := reader.Text()

    book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
    lib.AddBook(book)

    fmt.Println()
    fmt.Println("Book is added successfully!")
    fmt.Println()
}

func removeBook(lib *services.Library, reader *bufio.Scanner) {

    fmt.Print("Enter book ID: ")
    reader.Scan()
    id, _ := strconv.Atoi(reader.Text())

    lib.RemoveBook(id)
    fmt.Println()
    fmt.Println("Book is removed successfully!")
    fmt.Println()
}

func borrowBook(lib *services.Library, reader *bufio.Scanner) {

    fmt.Print("Enter book ID: ")
    reader.Scan()
    bookID, _ := strconv.Atoi(reader.Text())

    fmt.Print("Enter member ID: ")
    reader.Scan()
    memberID, _ := strconv.Atoi(reader.Text())

    err := lib.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println()
        fmt.Println("Error:", err)
        fmt.Println()
    } else {
        fmt.Println()
        fmt.Println("Book is borrowed successfully!")
        fmt.Println()
    }
}

func returnBook(lib *services.Library, reader *bufio.Scanner) {
    fmt.Print("Enter book ID: ")
    reader.Scan()
    bookID, _ := strconv.Atoi(reader.Text())

    fmt.Print("Enter member ID: ")
    reader.Scan()
    memberID, _ := strconv.Atoi(reader.Text())

    err := lib.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println()
        fmt.Println("Book is returned successfully!")
        fmt.Println()
    }
}

func listAvailableBooks(lib *services.Library) {
    books := lib.ListAvailableBooks()
    if len(books) == 0 {
        fmt.Println()
        fmt.Println("There are no available books.")
        fmt.Println()
    } else {
        fmt.Println("Available books:")
        for _, book := range books {
            fmt.Println()
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
            fmt.Println()
        }
    }
}

func listBorrowedBooks(lib *services.Library, reader *bufio.Scanner) {
    fmt.Print("Enter member ID: ")
    reader.Scan()
    memberID, _ := strconv.Atoi(reader.Text())

    books := lib.ListBorrowedBooks(memberID)
    if len(books) == 0 {
        fmt.Println()
        fmt.Println("There are no borrowed books by this member.")
        fmt.Println()
    } else {
        fmt.Println("Borrowed books:")
        fmt.Println()
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
    }
}
func registerUser(lib *services.Library, reader *bufio.Scanner){
    id := lib.GetNextUniqueMemberID()
    
    fmt.Println("Pls Enter Your Name: ")
    reader.Scan()
    name := reader.Text()

    user := models.Member{ID: id, Name: name}
    lib.SubscribeMember(user)

    fmt.Println()
    fmt.Println("Member is added successfully!")
    fmt.Println()
}

func listMembers(lib *services.Library){
    members := lib.ListAllMembers()

    if len(members) == 0 {
        fmt.Println()
        fmt.Println("There are no registered Users.")
        fmt.Println()
    } else {
        fmt.Println("Users:")
        for _, user := range members{
            fmt.Println()
            fmt.Printf("ID: %d, Name: %s, BorrowedBooks: %v\n", user.ID, user.Name, user.BorrowedBooks)
            fmt.Println()
        }
    }
}
