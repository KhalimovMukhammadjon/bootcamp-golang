package main

import (
	"fmt"
	"math/rand"

	//"sync"

	//"github.com/bxcodec/faker/v4"
	faker "github.com/icrowley/fake"
	//"Muhammadjon/bootcamp/28-lesson/"
)

type Users struct {
	name    string
	balance int
	product []Product
}

type Product struct {
	name  string
	price int
}

func product() []Product {
	var products []Product
	for i := 0; i < rand.Intn(20); i++ {
		products = append(products, Product{
			name:  faker.ProductName(),
			price: rand.Intn(100) * 1000,
		})
	}
	return products
}

func user() []Users {
	var users []Users
	for i := 0; i < 100; i++ {
		var prod = product()
		users = append(users, Users{
			name:    faker.FirstName(),
			balance: rand.Intn(100 * 1000),
			product: prod,
		})
	}
	return users
}

func cashier1(users Users, ch chan int) {
	//defer wg.Done()
	fmt.Println("--------------")
	sum := 0
	fmt.Println(users.name, users.product, users.balance)
	for _, v := range users.product {
		sum += v.price
	}
	ch <- sum
}

func cashier2(users Users, ch chan int) {
	sum := 0
	fmt.Println(users.name, users.product, users.balance)
	for _, v := range users.product {
		sum += v.price
	}
	ch <- sum
}

func cashier3(users Users, ch chan int) {
	sum := 0
	fmt.Println(users.name, users.product, users.balance)
	for _, v := range users.product {
		sum += v.price
	}
	ch <- sum
}

func main() {
	//var wg = sync.WaitGroup{}
	//var users Users
	users := user()
	ch1 := make(chan int)
	// ch2 := make(chan int)
	// ch3 := make(chan int)
	for i := 0; i < len(users)-3; i += 3 {
		go cashier1(users[i], ch1)
		u1 := <-ch1
		go cashier1(users[i], ch1)
		// u2 := <-ch1

		//wg.Wait()

		// go cashier2(users[i+1], ch2)
		// u2 := <-ch2
		// //wg.Add(1)
		// //wg.Wait()

		// go cashier3(users[i+2], ch3)
		// u3 := <-ch3
		//wg.Add(1)
		// wg.Add(3)
		// wg.Wait()
		fmt.Println(u1)
	}
}

// var users []Users

// for i := 0; i < 100; i++ {

// 	var prod = product()
// 	users = append(users, Users{
// 		name:    faker.FirstName(),
// 		balance: rand.Intn(100 * 1000),
// 		product: prod,
// 	})
// }
// fmt.Println("Users", users)
//user()
