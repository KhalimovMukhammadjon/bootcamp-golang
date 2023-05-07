package models

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// )

// func true() {
// 	// Create a PostgreSQL database and table.
// 	db, err := sql.Open("postgres", "user=postgres dbname=mydb")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	// Create a Golang struct that matches the schema of your PostgreSQL table.
// 	type Location struct {
// 		Long float64 `json:"long"`
// 		Lat  float64 `json:"lat"`
// 	}

// 	rows, err := db.Query("SELECT id, name, location FROM locations")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	// Use the Scan() method to scan the results of the SQL query into your Golang struct.
// 	for rows.Next() {
// 		var location Location
// 		err = rows.Scan(&location.ID, &location.Name, &location.Location)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(location)
// 	}

// 	// Use the json.Unmarshal() function to convert the value of the location column to a Golang struct.
// 	jsonString, err := json.Marshal(location.Location)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var location2 Location
// 	err = json.Unmarshal(jsonString, &location2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(location2)
// }

/*
// Use the db.Query() method to run a SQL query against your PostgreSQL database.
rows, err := db.Query("SELECT id, name, location FROM locations")
if err != nil {
panic(err)
}
defer rows.Close()

// Use the Scan() method to scan the results of the SQL query into your Golang struct.
for rows.Next() {
var location Location
err = rows.Scan(&location.ID, &location.Name, &location.Location)
if err != nil {
panic(err)
}
fmt.Println(location)
}

// Use the json.Unmarshal() function to convert the value of the location column to a Golang struct.
jsonString, err := json.Marshal(location.Location)
if err != nil {
panic(err)
}
var location2 Location
err = json.Unmarshal(jsonString, &location2)
if err != nil {
panic(err)
}
fmt.Println(location2)

*/
