package main

import (
	m "Muhammadjon/bootcamp/29-lesson/store/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Phone struct
// type Phone struct {
// 	ID     string
// 	Name   string
// 	Price  float64
// 	Color  string
// 	Memory string
// }

// Phones slice
var phones []m.Product

func GetProduct(c *gin.Context) {
	var products []m.Product
	for _, product := range m.Product {
		if product.Type == productType {
			products = append(products, product)
		}
	}
	return products
	c.JSON(http.StatusOK, phones)
	// phones = []m.Product{
	// 	m.Product{ID: 1,
	// 		Name:        "iPhone 13 Pro",
	// 		Description: "",
	// 		Price:       999.00,
	// 		Attributes: map[string][]string{

	// 		},
	// 			{"color":  "Silver",}
	// 			"memory": "256GB",
	// 		},
	// 	},
	// }
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	for _, phone := range phones {
		if phone.ID == id {
			c.JSON(http.StatusOK, phone)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Phone not found"})
}

func CreateProduct(c *gin.Context) {
	//var phone m.Product
	// c.BindJSON(&phone)
	// phones = append(phones, phone)
	// c.JSON(http.StatusCreated, phone)

	// phone := NewPhone(
	// 	"iPhone 13 Pro",
	// 	"The latest iPhone from Apple",
	// 	999,
	// 	map[string][]string{
	// 		"memory": []string{},
	// 		"color":  []string{},
	// 	})

	var phone m.Product

	ph := NewPhone("iPhone 13 Pro",
		"The latest iPhone from Apple",
		999,
		map[string][]string{
			"memory": []string{},
			"color":  []string{},
		})
	fmt.Println(ph)

	phones = append(phones, phone)
	c.BindJSON(&phone)
	c.JSON(http.StatusCreated, phone)
}

func NewPhone(name string, description string, price float32, attributes map[string][]string) *m.Product {
	return &m.Product{
		ID:          "",
		Name:        name,
		Description: description,
		Price:       price,
		Attributes:  attributes,
	}
}

func main() {
	// Initialize Gin engine
	r := gin.Default()

	// Get all phones
	r.GET("/phones", GetProduct)
	r.GET("/phones/:id", GetProductById)
	r.POST("/phones", CreateProduct)

	// Update phone
	r.PUT("/phones/:id", func(c *gin.Context) {
		id := c.Param("id")
		var phone m.Product
		c.BindJSON(&phone)
		for i, existingPhone := range phones {
			if existingPhone.ID == id {
				phones[i] = phone
				c.JSON(http.StatusOK, phone)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Phone not found"})
	})

	// Delete phone
	r.DELETE("/phones/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, phone := range phones {
			if phone.ID == id {
				phones = append(phones[:i], phones[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Phone deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Phone not found"})
	})

	// Listen on port 8080
	r.Run(":8080")
}
