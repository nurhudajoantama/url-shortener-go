package service

import (
	"errors"
	"fmt"
	"url-shortener/dto"
	"url-shortener/entity"
	"url-shortener/repository"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type urlService struct {
	urlRepository repository.UrlRepository
}

type UrlService interface {
	GetLongUrl(shortUrl string) (string, error)
	Create(urlDTO *dto.UrlRequestDTO) (*dto.UrlResponseDTO, error)
	Update(id uint64, urlDTO *dto.UrlRequestDTO) (*dto.UrlResponseDTO, error)
	Delete(id uint64) error
}

func NewUrlService(ur repository.UrlRepository) UrlService {
	return &urlService{
		urlRepository: ur,
	}
}

func (s *urlService) GetLongUrl(shortUrl string) (string, error) {
	url := &entity.Url{}
	if _, err := s.urlRepository.FindByShortUrl(shortUrl, url); err != nil {
		return "", err
	}
	return url.LongUrl, nil
}

func (s *urlService) Create(urlDTO *dto.UrlRequestDTO) (*dto.UrlResponseDTO, error) {
	url := &entity.Url{}

	if _, err := s.urlRepository.FindByShortUrl(urlDTO.ShortUrl, url); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("url telah terdaftar")
	}

	if err := smapping.FillStruct(url, smapping.MapFields(urlDTO)); err != nil {
		return nil, err
	}
	if _, err := s.urlRepository.Create(url); err != nil {
		return nil, err
	}
	response := &dto.UrlResponseDTO{}
	if err := smapping.FillStruct(response, smapping.MapFields(url)); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *urlService) Update(id uint64, urlDTO *dto.UrlRequestDTO) (*dto.UrlResponseDTO, error) {
	url := &entity.Url{}
	if _, err := s.urlRepository.FindById(id, url); err != nil {
		return nil, err
	}
	if err := smapping.FillStruct(url, smapping.MapFields(urlDTO)); err != nil {
		return nil, err
	}
	if _, err := s.urlRepository.Update(url); err != nil {
		return nil, err
	}
	response := &dto.UrlResponseDTO{}
	if err := smapping.FillStruct(response, smapping.MapFields(url)); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *urlService) Delete(id uint64) error {
	url := &entity.Url{}
	if _, err := s.urlRepository.FindById(id, url); err != nil {
		return err
	}
	if err := s.urlRepository.Delete(url); err != nil {
		return err
	}
	return fmt.Errorf("not implemented")
}
