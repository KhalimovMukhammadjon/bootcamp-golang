package model

type Basket struct {
	OrderID int      `json:"id"`
	Foods   []string `json:"food"`
	Sum     int      `json:"sum"`
	Balance int      `json:"balance"`
}
