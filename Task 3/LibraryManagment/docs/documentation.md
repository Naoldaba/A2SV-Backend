# Console-Based Library Management System

## Overview

This is console-based library management system written in Go. It demonstrates the use of structs, interfaces, methods, slices, and maps. The system allows users to manage books and members, including adding, removing, borrowing, and returning books.

## Features

- Add new books to the library.
- Remove existing books from the library.
- Borrow books for members.
- Return borrowed books.
- List all available books.
- List all books borrowed by a specific member.

## Structs

### Books
- ID
- Title
- Author
- Status

### Members
- ID
- Name
- BorrowedBooks

### Library
- map[int]Books
- map[int]Members


## Folders and Files

- `main.go`: Entry point of the application.
- `controllers/library_controller.go`: Handles console input and invokes the appropriate service methods.
- `models/book.go`: Defines the `Book` struct.
- `models/member.go`: Defines the `Member` struct.
- `services/library_service.go`: Contains business logic and data manipulation functions.
- `docs/documentation.md`: Contains system documentation and other related information.
- `go.mod`: Defines the module and its dependencies.