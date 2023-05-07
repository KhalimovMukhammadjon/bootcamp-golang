package handlers

import (
	"Muhammadjon/bootcamp/movie/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMovie(c *gin.Context) {
	var entity models.CreateMovie
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity, "this is entity")
	err = h.strg.Movie().Create(entity)

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

func (h *Handler) GetMovie(c *gin.Context) {
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

	resp, err := h.strg.Movie().GetMovieList(models.Query{Offset: offset, Limit: limit, Search: c.Query("search")})

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}

// 	c.JSON(200, resp)
// }

// func (h *Handler) GetListById(c *gin.Context) {
// 	id := c.Param("id")
// 	resp, err := h.strg.Todo().GetListById(id)

// 	if err != nil {
// 		c.JSON(400, models.DefaultError{
// 			Message: err.Error(),
// 		})
// 	}

// 	c.JSON(200, resp)
// }

// func (h *Handler) DeleteList(c *gin.Context) {
// 	id := c.Param("id")
// 	resp, err := h.strg.Todo().DeleteList(id)

// 	if err != nil {
// 		c.JSON(400, models.DefaultError{
// 			Message: err.Error(),
// 		})
// 	}

// 	c.JSON(200, resp)
// }

// func (h *Handler) UpdateList(c *gin.Context) {
// 	var entity models.UpdateTodo

// 	entity.Id = c.Param("id")

// 	err := c.BindJSON(&entity)
// 	if err != nil {
// 		c.JSON(400, models.DefaultError{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	resp, err := h.strg.Todo().UpdateList(entity)
// 	if err != nil {
// 		c.JSON(400, models.DefaultError{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, resp)
// }
