package handler

import (
	"net/http"
	"strconv"
	"url-shortener/dto"
	"url-shortener/helper"
	"url-shortener/service"

	"github.com/gin-gonic/gin"
)

type urlHandler struct {
	urlService service.UrlService
}

type UrlHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewUrlHandler(service service.UrlService) UrlHandler {
	return &urlHandler{
		urlService: service,
	}
}

func (h *urlHandler) Create(c *gin.Context) {
	urlDto := &dto.UrlRequestDTO{}
	if err := c.ShouldBind(&urlDto); err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}

	response, err := h.urlService.Create(urlDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, helper.CreateSuccessResponse("Success create data", response))
}

func (h *urlHandler) Update(c *gin.Context) {
	urlDto := &dto.UrlRequestDTO{}

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.CreateErrorResponse("Data tidak ditemukan", "Not Found"))
		return
	}

	if err := c.ShouldBind(&urlDto); err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}

	response, err := h.urlService.Update(id, urlDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.CreateSuccessResponse("Data berhasil di ubah", response))

}

func (h *urlHandler) Delete(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.CreateErrorResponse("Data tidak ditemukan", "Not Found"))
		return
	}

	if err := h.urlService.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, helper.CreateSuccessResponse("Data berhasil dihapus", nil))
}
