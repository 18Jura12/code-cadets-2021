package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	idValidator IdValidator
	statusValidator StatusValidator
	betService BetService
}

func NewController(idValidator IdValidator, statusValidator StatusValidator, betService BetService) *Controller {
	return &Controller{
		idValidator: idValidator,
		statusValidator: statusValidator,
		betService: betService,
	}
}

func (b *Controller) FetchBetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

