package handler

import (
	"errors"
	"net/http"
	"url-shortener/helper"
	"url-shortener/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type redirectHandler struct {
	urlService service.UrlService
}

type RedirectHandler interface {
	RedirectUrl(c *gin.Context)
}

func NewRedirectHandler(us service.UrlService) RedirectHandler {
	return &redirectHandler{
		urlService: us,
	}
}

func (h *redirectHandler) RedirectUrl(c *gin.Context) {
	reqShortUrl := c.Param("url")

	longUrl, err := h.urlService.GetLongUrl(reqShortUrl)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, helper.CreateErrorResponse("Data not found", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, helper.CreateErrorResponse("Error", err.Error()))
		return
	}

	c.Redirect(http.StatusMovedPermanently, longUrl)

}
