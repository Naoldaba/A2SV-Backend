package services

import (
    "library_managment/models"
    "errors"
    "sync"
)

var mu sync.Mutex

type LibraryManager interface {
    AddBook(book models.Book) error
    RemoveBook(bookId int) error
    BorrowBook(bookId int, memberId int) error
    ReturnBook(bookId int, memberId int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberId int) []models.Book
    SubscribeMember(member models.Member) error
    UnsubscribeMember(memberId int) error
    ListAllMembers() []models.Member
    GetNextUniqueBookID() int
    GetNextUniqueMemberID() int
}

type Library struct {
    Books   map[int]models.Book
    Members map[int]models.Member
    nextBookID  int
    nextUserID  int
}

func NewLibrary() LibraryManager {
    return &Library{
        Books:   make(map[int]models.Book),
        Members: make(map[int]models.Member),
        nextBookID:  1,
        nextUserID: 1,
    }
}

func (l *Library) GetNextUniqueBookID() int {
    mu.Lock()
    defer mu.Unlock()
    id := l.nextBookID
    l.nextBookID++
    return id
}

func (l *Library) GetNextUniqueMemberID() int {
    mu.Lock()
    defer mu.Unlock()
    id := l.nextUserID
    l.nextUserID++
    return id
}

func (l *Library) AddBook(book models.Book) error {
    for _, mem_book := range l.Books{
        if mem_book.ID == book.ID{
            return errors.New("book already added")
        }  
    }
    l.Books[book.ID] = book
    return nil
}

func (l *Library) ListAllMembers() []models.Member{
    var Members []models.Member
    for _, book := range l.Members {
        Members = append(Members, book)
    }
    return Members
}

func (l *Library) SubscribeMember(member models.Member) error {
    for _, mem_user := range l.Members{
        if mem_user.ID == member.ID{
            return errors.New("user already registered")
        }  
    }
    l.Members[member.ID] = member
    return nil
}

func (l *Library) UnsubscribeMember(memberId int) error {
    for _, mem_user := range l.Books{
        if mem_user.ID == memberId{
            delete(l.Members, memberId)
            return nil
        }
    }
    return errors.New("no such member")
    
}

func (l *Library) RemoveBook(bookId int) error {
    for _, mem_book := range l.Books{
        if mem_book.ID == bookId{
            delete(l.Books, bookId)
            return nil
        }
    }
    return errors.New("no such book")
}

func (l *Library) BorrowBook(bookId int, memberId int) error {
    book, exists := l.Books[bookId]
    if !exists || book.Status == "Borrowed" {
        return errors.New("book is not available")
    }

    member, exists := l.Members[memberId]
    if !exists {
        return errors.New("member not found. pls register first")
    }

    book.Status = "Borrowed"
    l.Books[bookId] = book

    member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.Members[memberId] = member

    return nil
}

func (l *Library) ReturnBook(bookId int, memberId int) error {
    book, exists := l.Books[bookId]
    if !exists || book.Status == "Available" {
        return errors.New("book is not borrowed")
    }

    member, exists := l.Members[memberId]
    if !exists {
        return errors.New("member not found")
    }

    book.Status = "Available"
    l.Books[bookId] = book

    for i, b := range member.BorrowedBooks {
        if b.ID == bookId {
            member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
            break
        }
    }

    l.Members[memberId] = member

    return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
    var availableBooks []models.Book
    for _, book := range l.Books {
        if book.Status == "Available" {
            availableBooks = append(availableBooks, book)
        }
    }
    return availableBooks
}

func (l *Library) ListBorrowedBooks(memberId int) []models.Book {
    member, exists := l.Members[memberId]
    if (!exists) {
        return nil
    }
    return member.BorrowedBooks
}
