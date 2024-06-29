package controller

import (
	"bbdk/app/api/dto"
	"bbdk/app/api/response"
	userRepo "bbdk/domain/repository/user"
	"bbdk/domain/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Response(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	user := req.ToEntity()
	if err := uc.service.CreateUser(user); err != nil {
		if errors.Is(err, userRepo.ErrAlreadyExist) {
			response.Response(c, nil, http.StatusBadRequest, err.Error())
			return
		}
		response.InternalServerError(c)
		return
	}

	response.Response(c, nil, http.StatusCreated, "")
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uc.service.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			response.NotFound(c)
		} else {
			response.InternalServerError(c)
		}
		return
	}
	userResponse := &dto.UserResponse{}
	userResponse.FromEntity(user)
	response.Response(c, userResponse, http.StatusOK, "")
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Response(c, nil, http.StatusBadRequest, "invalid id")
		return
	}

	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Response(c, nil, http.StatusBadRequest, err.Error())
		return
	}
	user := req.ToEntity()
	user.ID = uint(id)

	if err := uc.service.UpdateUser(user); err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			response.NotFound(c)
		} else if errors.Is(err, userRepo.ErrAlreadyExist) {
			response.Response(c, nil, http.StatusBadRequest, err.Error())
			return
		} else {
			response.InternalServerError(c)
		}
		return
	}

	response.Response(c, nil, http.StatusOK, "")
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Response(c, nil, http.StatusBadRequest, "invalid id")
		return
	}

	if err := uc.service.DeleteUser(uint(id)); err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			response.NotFound(c)
		} else {
			response.InternalServerError(c)
		}
		return
	}

	response.Response(c, nil, http.StatusOK, "")
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	users, err := uc.service.GetAllUsers(page, pageSize)
	if err != nil {
		response.InternalServerError(c)
		return
	}

	lnUsers := len(users)
	responses := make([]dto.UserResponse, lnUsers)
	for i := 0; i < lnUsers; i++ {
		userResponse := dto.UserResponse{}
		userResponse.FromEntity(users[i])
		responses[i] = userResponse
	}

	response.Response(c, responses, http.StatusOK, "")
}
