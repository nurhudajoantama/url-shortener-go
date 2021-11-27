package handler

import (
	"net/http"
	"url-shortener/dto"
	"url-shortener/helper"
	"url-shortener/service"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type authHandler struct {
	authService service.AuthService
	jwtService  service.JwtService
}

type AuthHandler interface {
	Login(*gin.Context)
	Register(*gin.Context)
}

func NewAuthHandler(as service.AuthService, js service.JwtService) AuthHandler {
	return &authHandler{
		authService: as,
		jwtService:  js,
	}
}

func (h *authHandler) Login(c *gin.Context) {
	login := &dto.LoginDTO{}
	if err := c.Bind(login); err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}
	user, err := h.authService.FindByUsername(login.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("Username tidak ditemukan", err.Error()))
		return
	}

	if res := (h.authService.VerifyPassword(user.Password, login.Password)); !res {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("Pssword salah", "wrong password"))
		return
	}

	token := h.jwtService.GenerateToken(user.Username)

	response := &dto.UserResponseDTO{}
	smapping.FillStruct(response, smapping.MapFields(user))
	response.AccessToken = token

	c.JSON(http.StatusOK, helper.CreateSuccessResponse("success", response))
}

func (h *authHandler) Register(c *gin.Context) {
	register := &dto.UserRequestDTO{}
	if err := c.Bind(register); err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}

	user, err := h.authService.Register(register)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}

	token := h.jwtService.GenerateToken(user.Username)

	response := &dto.UserResponseDTO{}
	smapping.FillStruct(response, smapping.MapFields(user))
	response.AccessToken = token

	c.JSON(http.StatusOK, helper.CreateSuccessResponse("success", response))
}
