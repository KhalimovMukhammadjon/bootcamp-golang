package read

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func HandleRequest() {
	url := "https://jsonplaceholder.typicode.com/posts" // Replace with your URL
    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()

    var data User
    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(data)

	// r, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	// if err != nil {
	// 	fmt.Println("Cannot get from URL", err)
	// }
	// defer r.Body.Close()

	// data, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Println("Error reading json data:", err)
	// }

	// var person User
	// err = json.Unmarshal(data, &person)
	// if err != nil {
	// 	fmt.Println("Error unmarshalling json data:", err)
	// }
	// fmt.Printf("%# v\n", pretty.Formatter(person))
}
