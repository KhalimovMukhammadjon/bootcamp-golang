package handlers

import (
	"Muhammadjon/bootcamp/stadium/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateActor godoc
// Tags actor
// @ID create-actor-handler
// @Router /actor [POST]
// @Summary Create Actor
// @Description Create Actor By Given Info
// @Param data body models.CreateActor true "Actor Body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SuccessResponse{data=string}
// @Failure default {object} models.DefaultError
func (h *Handler) CreateStadium(c *gin.Context) {
	var entity models.CreateStadium
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity)

	err = h.strg.Stadium().Create(entity)

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

func (h *Handler) GetStadiumList(c *gin.Context) {
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

	resp, err := h.strg.Stadium().GetStadiumList(models.Query{Offset: offset, Limit: limit, Search: c.Query("search")})

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}
