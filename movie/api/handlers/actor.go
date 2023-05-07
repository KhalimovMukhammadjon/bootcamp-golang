package handlers

import (
	"Muhammadjon/bootcamp/movie/models"
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
func (h *Handler) CreateActor(c *gin.Context) {
	var entity models.CreateActor
	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	fmt.Println(entity)

	err = h.strg.Actor().Create(entity)

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

// GetList godoc
// @tags actors
// @ID get-all-handler
// @Summary List actor
// @Description get all actors
// @Router /actor [GET]
// @Param offset query int true "offset"
// @Param limit query int true "limit"
// @Param search query string true "search"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Actors
// @Failure default {object} models.DefaultError
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

	resp, err := h.strg.Actor().GetList(models.Query{Offset: offset, Limit: limit, Search: c.Query("search")})

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}

// Update godoc
// @tags actors
// @ID update-actor
// @Summary Update actor
// @Description update actor
// @Router /actor/update [PUT]
// @Param body body models.UpdateActor true "body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Actors
// @Failure default {object} models.DefaultError
func (h *Handler) Update(c *gin.Context) {
	var entity models.UpdateActor

	err := c.BindJSON(&entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	resp, err := h.strg.Actor().Update(entity)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, resp)
}

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.strg.Actor().GetById(id)

	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	c.JSON(200, resp)
}
