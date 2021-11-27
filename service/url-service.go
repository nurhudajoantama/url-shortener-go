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
	GetLongUrl(string) (string, error)
	IsDuplicateUrl(string) bool
	Create(*dto.UrlRequestDTO) (*entity.Url, error)
	Update(uint64, *dto.UrlRequestDTO) (*entity.Url, error)
	Delete(uint64) error
}

func NewUrlService(ur repository.UrlRepository) UrlService {
	return &urlService{
		urlRepository: ur,
	}
}

func (s *urlService) GetLongUrl(shortUrl string) (string, error) {
	url := &entity.Url{}
	if err := s.urlRepository.FindByShortUrl(shortUrl, url); err != nil {
		return "", err
	}
	return url.LongUrl, nil
}

func (s *urlService) IsDuplicateUrl(shortUrl string) bool {
	url := &entity.Url{}
	err := s.urlRepository.FindByShortUrl(shortUrl, url)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (s *urlService) Create(urlDTO *dto.UrlRequestDTO) (*entity.Url, error) {
	url := &entity.Url{}

	if err := smapping.FillStruct(url, smapping.MapFields(urlDTO)); err != nil {
		return nil, err
	}
	if err := s.urlRepository.Create(url); err != nil {
		return nil, err
	}
	return url, nil
}

func (s *urlService) Update(id uint64, urlDTO *dto.UrlRequestDTO) (*entity.Url, error) {
	url := &entity.Url{}
	if err := s.urlRepository.FindById(id, url); err != nil {
		return nil, err
	}
	if err := smapping.FillStruct(url, smapping.MapFields(urlDTO)); err != nil {
		return nil, err
	}
	if err := s.urlRepository.Update(url); err != nil {
		return nil, err
	}
	return url, nil
}

func (s *urlService) Delete(id uint64) error {
	url := &entity.Url{}
	if err := s.urlRepository.FindById(id, url); err != nil {
		return err
	}
	if err := s.urlRepository.Delete(url); err != nil {
		return err
	}
	return fmt.Errorf("not implemented")
}
