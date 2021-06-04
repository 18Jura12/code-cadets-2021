package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	idValidator     IdValidator
	statusValidator StatusValidator
	betService      BetService
}

func NewController(idValidator IdValidator, statusValidator StatusValidator, betService BetService) *Controller {
	return &Controller{
		idValidator:     idValidator,
		statusValidator: statusValidator,
		betService:      betService,
	}
}

func (b *Controller) GetBetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Params.ByName("id")
		if id == "" {
			ctx.String(http.StatusBadRequest, "id not valid")
			return
		}

		if !b.idValidator.IdIsValid(id) {
			ctx.String(http.StatusBadRequest, "get request not valid")
			return
		}

		resultingBet, err := b.betService.GetBetById(ctx, id)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		ctx.JSON(http.StatusOK, resultingBet)
	}
}

func (b *Controller) GetBetsByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Params.ByName("id")
		if userId == "" {
			ctx.String(http.StatusBadRequest, "user id not valid")
			return
		}

		if !b.idValidator.IdIsValid(userId) {
			ctx.String(http.StatusBadRequest, "get request not valid")
			return
		}

		resultingBets, err := b.betService.GetBetsByUserId(ctx, userId)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		ctx.JSON(http.StatusOK, resultingBets)
	}
}

func (b *Controller) GetBetsByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Query("status")
		if status == "" {
			ctx.String(http.StatusBadRequest, "status not valid")
			return
		}

		if !b.statusValidator.StatusIsValid(status) {
			ctx.String(http.StatusBadRequest, "get request not valid")
			return
		}

		resultingBets, err := b.betService.GetBetsByStatus(ctx, status)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		ctx.JSON(http.StatusOK, resultingBets)
	}
}
