package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

var books []Book

func ReadJson() {
	data, err := ioutil.ReadFile("book.json")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	err = json.Unmarshal(data, &books)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
}

func (b *Library) GetBooks() []Book {
	return b.Books
}

func (b *Library) AddBook(book Book) {
	b.Books = append(b.Books, book)
}

// func (lib *Library) AddBook(book Book) {
// 	lib.Books = append(lib.Books, book)
// 	fmt.Println("Book added successfully!")
// }

func main() {
	file, err := os.Open("data.json") // Replace with your JSON file name
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var data Library
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(data)

	// GET books from storage
	library := &Library{Books: books}
	//fmt.Println(library.GetBooks())
	fmt.Println("Enter Title")
	fmt.Scanln(&title)
	fmt.Println("Enter author:")
	fmt.Scanln(&author)


	//Add new book to our storage
	library.AddBook(Book{
		ID:       id,
		Title:    title,
		Author:   author,
		Year:     year,
		Status:   status,
		Price:    price,
		Period:   period,
		Category: category,
		Page:     page,
	})

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

	// for _, v := range books {
	// 	fmt.Println(v)
	// }

	data, err = json.Marshal(library)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	err = ioutil.WriteFile("book.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
	fmt.Println(library.GetBooks())
}
