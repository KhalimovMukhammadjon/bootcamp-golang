package db

import (
	"Muhammadjon/bootcamp/13-lesson/library/storage/repo"
	//"Muhammadjon/bootcamp/13-lesson/library/cmd"
	"fmt"
)

type myLibrary repo.Library

func (r *myLibrary) GetBooks() []repo.Book {
	return r.Books
}

func (r *myLibrary) AddBook(book repo.Book) {
	var id int

	if len(r.Books) == 0 {
		id = 1
	} else {
		id = len(r.Books) + 1
		fmt.Println("Read id", id)
	}

	r.Books = append(r.Books, book)
	fmt.Println("Book added successfully!")
}

func (r *myLibrary) RemoveBook(book repo.Library) {
	for i, b := range r.Books {
		if b.ID == book.ID {
			r.Books = append(r.Books[:i], r.Books[i+1:]...)
			break
		}
	}
}

func (r *myLibrary) UpdateBook(book repo.Library) {
	for i, b := range r.Books {
		if b.ID == book.ID {
			r.Books[i] = book
			fmt.Println("Successfully updated")
			break
		}
	}
}

func (lib *myLibrary) GetBookById(book repo.Book) repo.Book {
	for _, b := range lib.Books {
		if b.ID == book.ID {
			return b
		}
	}
	return repo.Book{}
}

func (lib *myLibrary) GetBookByCategory(book repo.Book) []repo.Book {
	var books []repo.Book
	for _, b := range lib.Books {
		if b.Category == book.Category {
			books = append(books, b)
		}
	}
	return books
}
