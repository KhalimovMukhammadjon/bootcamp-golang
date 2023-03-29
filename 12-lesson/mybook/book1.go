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
	for _, v := range lib.Books {
		//fmt.Println(v)
		fmt.Println()
		fmt.Println("ID:", v.ID)
		fmt.Println("Title:", v.Title)
		fmt.Println("Author:", v.Author)
		fmt.Println("Year:", v.Year)
		fmt.Println("Status:", v.Status)
		fmt.Println("Price:", v.Price)
		fmt.Println("Period:", v.Period)
		fmt.Println("Category:", v.Category)
		fmt.Println("Page:", v.Page)
		fmt.Println()
		fmt.Println("------------------")

	}
	return lib.Books
}

func (lib *Library) AddBook(book Book) {
	// if len(lib.Books) == 0 {
	// 	id = 1
	// } else {
	// 	id = len(lib.Books) + 1
	// 	fmt.Println("Read id", id)
	// }
	lib.Books = append(lib.Books, book)
	fmt.Println("Book added successfully!")
}

func (lib *Library) RemoveBook(book Book) {
	for i, b := range lib.Books {
		if b.ID == book.ID {
			lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
			fmt.Println("Book removed successfully")
			break
		}
	}
}

func (lib *Library) UpdateBook(book Book) {
	for i, b := range lib.Books {
		if b.ID == book.ID {
			lib.Books[i] = book
			//lib.Books[i].Title = title
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

	var str string
	fmt.Println("1.GetBooks\n2.AddBook\n3.RemoveBook\n4.UpdateBook\n5.GetBookById\n6.GetBookByCategory")
	fmt.Scanln(&str)
	// GetBooks() []Library
	// AddBook(Book)
	// RemoveBookBook(Book)
	// UpdateBook(Book)
	// GetBookByld(Book)
	// GetBookByCategory(Book)

	library := &Library{Books: books.Books}

	switch str {
	case "1":
		fmt.Println("Welcome to our library")
		library.GetBooks()
		//fmt.Println("Read",library.GetBooks())
	case "2":

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

		newBook := Book{
			ID:       len(books.GetBooks()) + 1,
			Title:    title,
			Author:   author,
			Year:     year,
			Status:   status,
			Price:    price,
			Period:   period,
			Category: category,
			Page:     page,
		}
		library.AddBook(newBook)

		fmt.Println("Read", library.GetBooks())
	case "3":
		fmt.Println("Enter Id:")
		fmt.Scanln(&id)

		library.RemoveBook(Book{
			ID: id,
		})

		// removeBook := Book{
		// 	ID: id,
		// }
		// fmt.Println(removeBook)
	case "4":
		library.UpdateBook(Book{
			ID:     id,
			Title:  "Atomic Habits",
			Author: "John Doe",
			Price:  18900,
		})
		//fmt.Println()
		//library.UpdateBook()
	}

	// GET books from storage

	// fmt.Println("Enter Id:")
	// fmt.Scanln(&id)

	//added new book to our storage

	// update book
	// library.UpdateBook(Book{
	// 	ID:     id,
	// 	Title:  "Atomic Habits",
	// 	Author: "John Doe",
	// 	Price:  18900,
	// })
	//library.UpdateBook()

	// get book by id
	getBookByCategory := Book{
		Category: category,
	}
	fmt.Println()
	fmt.Println("Book's category name is:", library.GetBookByCategory(getBookByCategory))

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
	//fmt.Println("New books", library.GetBooks())
}
