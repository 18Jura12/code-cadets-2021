package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/api/controllers/models"
	"net/http"
)

type Controller struct {
	betValidator BetValidator
	betService BetService
}

func NewController(betValidator BetValidator, betService BetService) *Controller {
	return &Controller{
		betValidator: betValidator,
		betService:   betService,
	}
}

func (c *Controller) CreateBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var betRequestDto models.BetRequestDto
		err := ctx.ShouldBindWith(&betRequestDto, binding.JSON)
		if err != nil {
			ctx.String(http.StatusBadRequest, "bet request is not valid:" + err.Error())
			return
		}

		var valid bool
		valid, err = c.betValidator.IsBetValid(betRequestDto)
		if !valid {
			ctx.String(http.StatusBadRequest, "bet request is not valid: " + err.Error())
			return
		}

		err = c.betService.CreateBet(betRequestDto)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed." + err.Error())
			return
		}

		ctx.Status(http.StatusOK)
	}
}

