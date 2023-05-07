package models

// type Product struct {
// 	ID         string  `json:"id"`
// 	Name       string  `json:"name"`
// 	Price      float32 `json:"price"`
// 	UserID string  `json:"user_id"`
// 	CustomeID string `json:"customer_id"`
// 	CourierID string `json:"courier_id"`
// 	ProductID string `json:"product_id"`

// }

// type CreateProduct struct {
// 	Name       string  `json:"name"`
// 	Price      float32 `json:"price"`
// 	CategoryID string  `json:"category_id"`
// }

// type UpdateProduct struct {
// 	Name       string  `json:"name"`
// 	Price      float32 `json:"price"`
// 	CategoryID string  `json:"category_id"`
// }



// CREATE TABLE "orders" (
//     "id" UUID NOT NULL PRIMARY KEY,
//     "name" VARCHAR NOT NULL,
//     "price" NUMERIC NOT NULL,
//     "user_id" UUID NOT NULL REFERENCES "users"("id"),
//     "customer_id" UUID NOT NULL REFERENCES "customers"("id"),
//     "courier_id" UUID NOT NULL REFERENCES "couriers"("id"),
//     "product_id" UUID NOT NULL REFERENCES "products"("id"),
//     "quantity" NUMERIC NOT NULl DEFAULT 1
// );

