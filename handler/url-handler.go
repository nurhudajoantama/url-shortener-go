package handler

import (
	"net/http"
	"strconv"
	"url-shortener/dto"
	"url-shortener/helper"
	"url-shortener/service"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type urlHandler struct {
	urlService service.UrlService
}

type UrlHandler interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func NewUrlHandler(service service.UrlService) UrlHandler {
	return &urlHandler{
		urlService: service,
	}
}

func (h *urlHandler) Create(c *gin.Context) {
	urlDto := &dto.UrlRequestDTO{}
	if err := c.ShouldBind(urlDto); err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}

	if duplicate := h.urlService.IsDuplicateUrl(urlDto.ShortUrl); duplicate {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", "Short Url telah terdaftar"))
		return
	}

	url, err := h.urlService.Create(urlDto)

	response := &dto.UrlResponseDTO{}
	smapping.FillStruct(response, smapping.MapFields(url))

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, helper.CreateSuccessResponse("Success create data", response))
}

func (h *urlHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.CreateErrorResponse("Data tidak ditemukan", "Not Found"))
		return
	}

	urlDto := &dto.UrlRequestDTO{}
	if err := c.ShouldBind(&urlDto); err != nil {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", err.Error()))
		return
	}

	if duplicate := h.urlService.IsDuplicateUrl(urlDto.ShortUrl); duplicate {
		c.JSON(http.StatusBadRequest, helper.CreateErrorResponse("error", "Short Url telah terdaftar"))
		return
	}

	url, err := h.urlService.Update(id, urlDto)

	response := &dto.UrlResponseDTO{}
	smapping.FillStruct(response, smapping.MapFields(url))

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
