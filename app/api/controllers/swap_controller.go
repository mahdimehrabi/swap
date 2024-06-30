package controller

import (
	"bbdk/app/api/dto"
	"bbdk/app/api/response"
	"bbdk/domain/entity"
	transRepo "bbdk/domain/repository/transaction"
	"bbdk/domain/repository/user"
	"bbdk/domain/service"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type SwapController struct {
	tranService service.TransactionService
}

func NewSwapController(tranService service.TransactionService) *SwapController {
	return &SwapController{tranService: tranService}
}

func (uc *SwapController) CreateTransaction(c *gin.Context) {
	tranReq := dto.TransactionReq{}
	if err := c.ShouldBindJSON(&tranReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user := &entity.User{
		Model: gorm.Model{
			ID: tranReq.UserID,
		},
	}
	srcCoin := &entity.Coin{
		Model: gorm.Model{
			ID: tranReq.SrcCoinID,
		},
	}

	destCoin := &entity.Coin{
		Model: gorm.Model{
			ID: tranReq.DestCoinID,
		},
	}

	trans, err := uc.tranService.CreateTransaction(context.Background(), user, srcCoin, destCoin, tranReq.SrcCoinAmount)
	if err != nil {
		response.InternalServerError(c)
		return
	}
	tranRes := new(dto.Transaction)
	tranRes.FromEntity(trans)
	response.Response(c, tranRes, http.StatusCreated, "")
}

func (uc *SwapController) CommitTransaction(c *gin.Context) {
	id := c.Param("id")
	uuID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	transaction := entity.NewTransaction(0, 0, 0)
	transaction.ID = uuID

	if err := uc.tranService.CommitTransaction(context.Background(), transaction); err != nil {
		if errors.Is(err, user.ErrNotEnoughBalance) {
			response.Response(c, gin.H{}, http.StatusBadRequest, "not enough balance")
			return
		} else if errors.Is(err, transRepo.ErrNotFound) {
			response.NotFound(c)
			return
		}
		response.InternalServerError(c)
	}
	response.Response(c, gin.H{}, http.StatusOK, "transaction completed successfully")
}
