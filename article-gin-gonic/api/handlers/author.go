package handlers

import (
	"bootcamp/article/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// type AuthorRepoI interface {
// 	Create(entity models.Person) (err error)
// 	GetList(query models.Query) (resp []models.Person, err error)
// 	GetByID(ID int) (resp models.Person, err error)
// 	Update(entity models.Person) (effectedRowsNum int, err error)
// 	Delete(ID string) (effectedRowsNum int, err error)
// }

func (h *Handler) Create(c *gin.Context) {
	var entity models.PersonCreateModel
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity)

	err = h.strg.Author().Create(entity)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, models.SuccessResponse{
		Message: "article has been created",
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

	resp, err := h.strg.Author().GetList(models.Query{Offset: offset, Limit: limit, Search: c.Query("search")})

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}

func (h *Handler) GetMostPublisher(c *gin.Context) {
	//id := c.Param("id")

	resp, err := h.strg.Author().GetMostPublisher()

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}

func (h *Handler) GetMinPublisher(c *gin.Context) {
	//id := c.Param("id")

	resp, err := h.strg.Author().GetMinPublisher()

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}
