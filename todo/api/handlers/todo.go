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

func (h *Handler) GetList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	resp, err := h.strg.Todo().GetList(models.Query{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}

func (h *Handler) GetListById(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.strg.Todo().GetListById(id)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}

func (h *Handler) DeleteList(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.strg.Todo().DeleteList(id)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}
