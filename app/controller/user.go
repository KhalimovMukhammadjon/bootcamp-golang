package controller

import (
	"app/models"
	"errors"
	"fmt"
	"strings"
)

func (c *Controller) CreateUser(req *models.CreateUser) (string, error) {
	id, err := c.store.User().Create(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) DeleteUser(req *models.UserPrimaryKey) error {
	err := c.store.User().Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateUser(req *models.UpdateUser, userId string) error {
	err := c.store.User().Update(req, userId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetByIdUser(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User().GetByID(req)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// func (c *Controller) GetByIdUserss(req *models.UserPrimaryKey) (models.ShopCart, error) {
// 	user, err := c.store.User().GetByID(req)
// 	if err != nil {
// 		return models.ShopCart{}, err
// 	}
// 	return user, nil
// }

func (c *Controller) WithdrawCheque(total float64, userId string) error {
	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: userId})
	if err != nil {
		return err
	}
	if user.Balance >= total {
		user.Balance -= total
	} else {
		return errors.New("You don't have enough money")
	}

	err = c.store.User().Update(&models.UpdateUser{
		Balance: user.Balance,
	}, userId)
	if err != nil {
		return err
	}

	err = c.store.ShopCart().UpdateShopCart(userId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) MoneyTransfer(sender string, receiver string, money float64) error {
	send, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: sender})
	if err != nil {
		return err
	}

	receive, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: receiver})
	if err != nil {
		return err
	}

	comMoney := 0.1 * float64(money)
	if money+comMoney > send.Balance {
		return errors.New("Sender doesn't have enough money")
	}
	send.Balance -= money + comMoney
	err = c.store.User().Update(&models.UpdateUser{
		Name:    send.Name,
		Surname: send.Surname,
		Balance: send.Balance,
	}, sender)
	if err != nil {
		return err
	}

	err = c.store.Commission().AddCommission(&models.Commission{
		Balance: comMoney,
	})

	receive.Balance += money
	err = c.store.User().Update(&models.UpdateUser{
		Name:    receive.Name,
		Surname: receive.Surname,
		Balance: receive.Balance,
	}, receiver)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetPurchaseHistory(userID *models.UserPrimaryKey) (string, error) {
	user, err := c.store.User().GetByID(userID)
	if err != nil {
		return "", err
	}

	carts, err := c.store.ShopCart().GetUserShopCart(userID)
	if err != nil {
		return "", err
	}

	type cartInfo struct {
		Id    int
		Name  string
		Price float64
		Count int
		Total float64
		//Time  time.Time
	}
	// type ShopCart struct {
	// 	Id        string `json:"id"`
	// 	ProductId string `json:"productId"`
	// 	UserId    string `json:"userID"`
	// 	Count     int    `json:"count"`
	// 	Status    bool   `json:"status"`
	// }

	var cartInfos []cartInfo
	for _, cart := range carts {
		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: cart.ProductId})
		if err != nil {
			return "", err
		}
		info := cartInfo{
			Name:  product.Name,
			Price: product.Price,
			Count: cart.Count,
			Total: float64(cart.Count) * product.Price,
			// Id: product.Id,
			// ProductId: product.Name,
			// UserId: product.Id,
			// Count: int(product.Price),
			//Time:  cart.Date,
		}
		cartInfos = append(cartInfos, info)
	}

	var output strings.Builder
	fmt.Fprintf(&output, "Client Name: %s\n", user.Name)
	for i, info := range cartInfos {
		fmt.Fprintf(&output, "%d. Name: %s Price: %.0f Count: %d Total: %.0f Time: %v\n", i+1, info.Name, info.Price, info.Count, info.Total)
	}
	return output.String(), nil
}

// func (c *Controller) GetByIdUser(req *models.UserPrimaryKey) (models.User, error) {
// 	user, err := c.store.User().GetByID(req)
// 	if err != nil {
// 		return models.User{}, err
// 	}
// 	return user, nil
// }

func (c *Controller) GetPurchaseHistory1(req *models.User) (models.User, error) {

	// user, err := c.store.ShopCart().GetUserShopCart(&models.UserPrimaryKey{

	// })
	// if err != nil {
	// 	return models.User{}, err
	// }
	// return user, nil
	users, err := c.store.ShopCart().GetUserShopCart(req)
	if err != nil {
		return models.User{}, err
	}
	fmt.Println(users)

	var shopcart []models.ShopCart
	var mId string

	for _, v := range users {
		var cart models.ShopCart

		if v.UserId == mId {
			//cart. = v.Time
			cart.Count = v.Count
			mId = v.ProductId

			for _, k := range shopcart {
				if mId == k.Id {
					cart.ProductId = k.Id
					cart.ProductPrice = k.Price

				}
			}

			result := 0
			for i := 0; i < cart.Count; i++ {
				result += cart.ProductPrice
			}
			cart.sum = result

			shopcart = append(shopcart, cart)

		}
	}
}

// func (c *Client) TotalSpent() float64 {
//     total := 0.0
//     for _, s := range c.Sales {
//         total += float64(s.Quantity) * s.Price
//     }
//     return total
// }

// func ProductSalesByClient(clients []Client, productID int) map[string]int {
//     salesByClient := make(map[string]int)
//     for _, c := range clients {
//         for _, s := range c.Sales {
//             if s.ProductID == productID {
//                 salesByClient[c.Name] += s.Quantity
//             }
//         }
//     }
//     return salesByClient
// }
