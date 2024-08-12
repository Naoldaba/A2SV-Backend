package controllers

import (
	"bufio"
	"fmt"
	"library_managment/models"
	"library_managment/services"
	"os"
	"strconv"
	"strings"
	"regexp"
)

func LibraryController(lib services.LibraryManager) {

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

func isValidString(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	return re.MatchString(s)
}

func addBook(lib services.LibraryManager, reader *bufio.Scanner) {
	id := lib.GetNextUniqueBookID()
	fmt.Print("Enter book Title: ")
	if !reader.Scan() {
		if err := reader.Err(); err != nil {
			fmt.Println("Error reading book title:", err)
		} else {
			fmt.Println("Unexpected error reading book title.")
		}
		return
	}
	title := strings.TrimSpace(reader.Text())
	if title == "" {
		fmt.Println("Book title cannot be empty.")
		return
	}
    
	fmt.Print("Enter book Author: ")
	if !reader.Scan() {
		if err := reader.Err(); err != nil {
			fmt.Println("Error reading book author:", err)
		} else {
			fmt.Println("Unexpected error reading book author.")
		}
		return
	}

	author := strings.TrimSpace(reader.Text())
	if author == "" {
		fmt.Println("Book author cannot be empty.")
		return
	}

    if !isValidString(author){
        fmt.Println("\nBook author must only contain letters and spaces.\n")
		return
    }

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	if err := lib.AddBook(book); err != nil {
		fmt.Println("Error adding book to the library:", err)
		return
	}

	fmt.Println("\nBook is added successfully!\n")
}

func removeBook(lib services.LibraryManager, reader *bufio.Scanner) {
	fmt.Print("Enter book ID: ")
	if !reader.Scan() {
		if err := reader.Err(); err != nil {
			fmt.Println("Error reading book ID:", err)
		} else {
			fmt.Println("Unexpected error reading book ID.")
		}
		return
	}
	idStr := strings.TrimSpace(reader.Text())
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid book ID:", err)
		return
	}

	if err := lib.RemoveBook(id); err != nil {
		fmt.Println("Error removing book from the library:", err)
		return
	}

	fmt.Println("\nBook is removed successfully!\n")
}

func borrowBook(lib services.LibraryManager, reader *bufio.Scanner) {
	fmt.Print("Enter book ID: ")
	if !reader.Scan() {
		if err := reader.Err(); err != nil {
			fmt.Println("Error reading book ID:", err)
		} else {
			fmt.Println("Unexpected error reading book ID.")
		}
		return
	}
	bookIDStr := strings.TrimSpace(reader.Text())
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println("Invalid book ID:", err)
		return
	}

	fmt.Print("Enter member ID: ")
	if !reader.Scan() {
		if err := reader.Err(); err != nil {
			fmt.Println("Error reading member ID:", err)
		} else {
			fmt.Println("Unexpected error reading member ID.")
		}
		return
	}
	memberIDStr := strings.TrimSpace(reader.Text())
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID:", err)
		return
	}

	err = lib.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
	} else {
		fmt.Println("\nBook is borrowed successfully!\n")
	}
}

func returnBook(lib services.LibraryManager, reader *bufio.Scanner) {
	fmt.Print("Enter book ID: ")
	if !reader.Scan() {
		if err := reader.Err(); err != nil {
			fmt.Println("Error reading book ID:", err)
		} else {
			fmt.Println("Unexpected error reading book ID.")
		}
		return
	}
	bookIDStr := strings.TrimSpace(reader.Text())
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println("Invalid book ID:", err)
		return
	}

	fmt.Print("Enter member ID: ")
	if !reader.Scan() {
		if err := reader.Err(); err != nil {
			fmt.Println("Error reading member ID:", err)
		} else {
			fmt.Println("Unexpected error reading member ID.")
		}
		return
	}
	memberIDStr := strings.TrimSpace(reader.Text())
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID:", err)
		return
	}

	err = lib.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Printf("\nError %v\n", err)
	} else {
		fmt.Println("\nBook is returned successfully!\n")
	}
}

func listAvailableBooks(lib services.LibraryManager) {
	books := lib.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("\nThere are no available books.\n")
	} else {
		fmt.Println("\nAvailable Books:")
		fmt.Println("--------------------------------------------------")
		fmt.Printf("| %-5s | %-30s | %-20s |\n", "ID", "Title", "Author")
		fmt.Println("--------------------------------------------------")
		for _, book := range books {
			fmt.Printf("| %-5d | %-30s | %-20s |\n", book.ID, book.Title, book.Author)
		}
		fmt.Println("--------------------------------------------------")
	}
}

func listBorrowedBooks(lib services.LibraryManager, reader *bufio.Scanner) {
	fmt.Print("Enter member ID: ")
	if !reader.Scan() {
		fmt.Println("Error reading member ID.")
		return
	}

	memberIDStr := strings.TrimSpace(reader.Text())
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID.")
		return
	}

	books := lib.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("\nThere are no borrowed books by this member.\n")
	} else {
		fmt.Println("\nBorrowed Books:")
		fmt.Println("------------------------------------------------------------------")
		fmt.Printf("| %-5s | %-30s | %-20s |\n", "ID", "Title", "Author")
		fmt.Println("------------------------------------------------------------------")
		for _, book := range books {
			fmt.Printf("| %-5d | %-30s | %-20s |\n", book.ID, book.Title, book.Author)
		}
		fmt.Println("--------------------------------------------------")
	}
}

func registerUser(lib services.LibraryManager, reader *bufio.Scanner) {
	id := lib.GetNextUniqueMemberID()

	fmt.Print("Please enter your name: ")
	if !reader.Scan() {
		fmt.Println("Error reading input. Please try again.")
		return
	}
	name := strings.TrimSpace(reader.Text())

	if name == "" {
		fmt.Println("Error: Name cannot be empty.")
		return
	}

	user := models.Member{ID: id, Name: name}
	
	err := lib.SubscribeMember(user)
	if err != nil {
		fmt.Printf("Error registering member: %v\n", err)
		return
	}

	fmt.Println("\nMember is added successfully!\n")
}

func listMembers(lib services.LibraryManager) {
	members := lib.ListAllMembers()

	if len(members) == 0 {
		fmt.Println("\nThere are no registered users.\n")
		return
	}

	fmt.Println("Registered Users:")
    fmt.Println("-------------------------------")
    fmt.Printf("| %-5s | %-20s |\n", "ID", "Name")
    fmt.Println("-------------------------------")
    for _, user := range members {
        fmt.Printf("| %-5d | %-20s |\n", user.ID, user.Name)
    }
    fmt.Println("-------------------------------")
}