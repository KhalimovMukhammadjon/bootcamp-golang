package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BookData interface {
	GetBooks() []Library
	AddBook(Book)
	RemoveBookBook(Book)
	UpdateBook(Book)
	GetBookByld(Book)
	GetBookByCategory(Book)
}

type Library struct {
	Books []Book
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Status   string `json:"status"`
	Price    int    `json:"price"`
	Period   int    `json:"period"`
	Category string `json:"category"`
	Page     int    `json:"page"`
}

var id, year, price, period, page int
var title, author, status, category string

func (lib *Library) GetBooks() []Book {
	// for _, v := range lib.Books {
	// 	fmt.Println("b", v)
	// }
	return lib.Books
}

func (lib *Library) AddBook(book Book) {
	// for _, v := range lib.Books {
	// 	fmt.Println("New idd", v)
	// 	if v.ID == 0 {
	// 		fmt.Println("None")
	// 		id = 1
	// 	} else if v.ID != 0 {
	// 		fmt.Println("Created")
	// 		id = v.ID + 1
	// 	}
	// }

	//second version
	// if len(lib.Books) == 0 {
	// 	id = 1
	// } else if len(lib.Books) != 0 {
	// 	id = len(lib.Books) + 1
	// }
	lib.Books = append(lib.Books, book)
	fmt.Println("Book added successfully!")
}

func (lib *Library) RemoveBook(book Book) {
	for i, b := range lib.Books {
		if b.ID == book.ID {
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			break
		}
	}
}

func (lib *Library) UpdateBook(book Book) {
	for i, b := range lib.Books {
		if b.ID == book.ID {
			lib.Books[i] = book
			fmt.Println("Successfully updated")
			break
		}
	}
}

func (lib *Library) GetBookById(book Book) Book {
	for _, b := range lib.Books {
		if b.ID == book.ID {
			return b
		}
	}
	return Book{}
}

func (lib *Library) GetBookByCategory(book Book) []Book {
	var books []Book
	for _, b := range lib.Books {
		if b.Category == book.Category {
			books = append(books, b)
		}
	}
	return books
}

func main() {
	data, err := ioutil.ReadFile("mybook.json")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var books Library
	err = json.Unmarshal(data, &books)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// GET books from storage
	library := &Library{Books: books.Books}
	fmt.Println("Read", library.GetBooks())

	fmt.Println("Enter Id:")
	fmt.Scanln(&id)

	fmt.Println("Enter title")
	fmt.Scanln(&title)

	fmt.Println("Enter author")
	fmt.Scanln(&author)

	fmt.Println("Enter year:")
	fmt.Scanln(&year)

	fmt.Println("Enter price:")
	fmt.Scanln(&price)

	fmt.Println("Enter period:")
	fmt.Scanln(&period)

	fmt.Println("Enter category:")
	fmt.Scanln(&category)

	fmt.Println("Enter page:")
	fmt.Scanln(&page)

	// added new book to our storage
	// newBook := Book{
	// 	ID:       id,
	// 	Title:    title,
	// 	Author:   author,
	// 	Year:     year,
	// 	Status:   status,
	// 	Price:    price,
	// 	Period:   period,
	// 	Category: category,
	// 	Page:     page,
	// }
	// library.AddBook(newBook)

	// update book
	library.UpdateBook(Book{
		ID:     id,
		Title:  "Atomic Habits",
		Author: "John Doe",
		Price:  18900,
	})
	//library.UpdateBook()

	// getBookByCategory := Book{
	// 	Category: category,
	// }
	// fmt.Println("Book's category name is:", library.GetBookByCategory(getBookByCategory))

	data, err = json.Marshal(library)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	err = ioutil.WriteFile("mybook.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	// getBookByID := Book{
	// 	ID: id,
	// }
	// fmt.Println("This book is", library.GetBookById(getBookByID))
	fmt.Println("New books", library.GetBooks())
}
