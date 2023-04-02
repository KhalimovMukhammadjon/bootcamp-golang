package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsonDb"
	"fmt"
	"log"
)

type Product2 struct {
	name         string
	quantitySold int
}

func main() {
	cfg := config.Load()

	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("error while connecting to database")
	}
	defer jsonDb.CloseDb()

	c := controller.NewController(&cfg, jsonDb)

	Product(c)
	Category(c)

}

func Product(c *controller.Controller) {

	// c.CreateProduct(&models.CreateProduct{
	// 	Name:       "Smartfon vivo V25 8/256 GB",
	// 	Price:      4_860_000,
	// 	CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
	// })

	// product, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: "38292285-4c27-497b-bc5f-dfe418a9f959"})

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	products, err := c.GetAllProduct(
		&models.GetListProductRequest{
			Offset:     0,
			Limit:      1,
			CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
		},
	)
	//tr := c.GetByIdUser() 2d0a65a6-8a5d-4850-9211-38a78cc37595

	if err != nil {
		log.Println(err)
		return
	}

	// user, err := c.GetByIdUser(
	// 	&models.UserPrimaryKey{
	// 		Id: "e4421e6d-cf37-4dd7-a87f-97c91feffaef",
	// 	},
	// )
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println("read", user)

	// 2
	userID := &models.UserPrimaryKey{Id: "0743751d-2ae0-4637-b4d6-23f63d96d4a3"}
	history, err := c.GetPurchaseHistory(userID)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(history)

	for in, product := range products.Products {
		fmt.Println(in+1, product)
	}

	//3
	// products := []Product2{
	// 	{"Product A", 10},
	// 	{"Product B", 20},
	// 	{"Product C", 30},
	// 	{"Product D", 25},
	// }

	// totalQuantitySold := 0
	// for _, product := range products {
	// 	totalQuantitySold += product.quantitySold
	// }

	// fmt.Println("The total quantity of products sold in is:", totalQuantitySold)
}

func Category(c *controller.Controller) {
	// c.CreateCategory(&models.CreateCategory{
	// 	Name:     "Smartfonlar va telefonlar",
	// 	ParentID: "eed2e676-1f17-429f-b75c-899eda296e65",
	// })

	// category, err := c.GetByIdCategory(&models.CategoryPrimaryKey{Id: "eed2e676-1f17-429f-b75c-899eda296e65"})
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	//fmt.Println(category)

}

func User(c *controller.Controller) {
	sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
	receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
	err := c.MoneyTransfer(sender, receiver, 500_000)
	if err != nil {
		log.Println(err)
	}
}
