package handlers

import (
	"Muhammadjon/bootcamp/todo/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTodo(c *gin.Context) {
	var entity models.CreateTodo
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity, "this is entity")
	err = h.strg.Todo().Create(entity)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "Todo has been created",
		Data:    "ok",
	})
}
