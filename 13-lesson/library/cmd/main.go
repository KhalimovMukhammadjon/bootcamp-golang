package main

// import (
// 	"Muhammadjon/bootcamp/13-lesson/library/storage/db"
// 	"Muhammadjon/bootcamp/13-lesson/library/storage/repo"

// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// )

// var id, year, price, period, page int
// var title, author, status, category string

// // type AllBook repo.BookData

// func main() {
// 	data, err := ioutil.ReadFile("mybook.json")
// 	if err != nil {
// 		fmt.Println("Error reading file", err)
// 		return
// 	}

// 	var books []repo.Library
// 	err = json.Unmarshal(data, &books)
// 	if err != nil {
// 		fmt.Println("Error unmarshaling JSON:", err)
// 		return
// 	}

// 	library := &repo.Library{Books: books}
// 	fmt.Println("Read", library.Books())
// 	// AllBook.GetBooks()

// 	library1 := &repo.Library{Books: books.Books}
// 	fmt.Println("Read", library1.GetBooks())
// 	//repo.BookData.GetBooks()

// 	// fmt.Println("Enter Id:")
// 	// fmt.Scanln(&id)

// 	fmt.Println("Enter title")
// 	fmt.Scanln(&title)

// 	fmt.Println("Enter author")
// 	fmt.Scanln(&author)

// 	fmt.Println("Enter year:")
// 	fmt.Scanln(&year)

// 	fmt.Println("Enter price:")
// 	fmt.Scanln(&price)

// 	fmt.Println("Enter period:")
// 	fmt.Scanln(&period)

// 	fmt.Println("Enter category:")
// 	fmt.Scanln(&category)

// 	fmt.Println("Enter page:")
// 	fmt.Scanln(&page)

// 	//added new book to our storage
// 	newBook := repo.Book{
// 		ID:       id,
// 		Title:    title,
// 		Author:   author,
// 		Year:     year,
// 		Status:   status,
// 		Price:    price,
// 		Period:   period,
// 		Category: category,
// 		Page:     page,
// 	}
// 	library.AddBook(newBook)

// 	library.UpdateBook(repo.Book{
// 		ID:     id,
// 		Title:  "Atomic Habits",
// 		Author: "John Doe",
// 		Price:  18900,
// 	})

// 	data, err = json.Marshal(library)
// 	if err != nil {
// 		fmt.Println("Error writing JSON to file:", err)
// 		return
// 	}

// 	err = ioutil.WriteFile("mybook.json", data, 0644)
// 	if err != nil {
// 		fmt.Println("Error writing JSON to file:", err)
// 		return
// 	}

// 	fmt.Println("New books", library.GetBooks())
// }
