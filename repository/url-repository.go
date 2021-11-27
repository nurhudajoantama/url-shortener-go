package repository

import (
	"url-shortener/entity"

	"gorm.io/gorm"
)

type urlRepository struct {
	DB *gorm.DB
}

type UrlRepository interface {
	FindByShortUrl(string, *entity.Url) error
	FindById(uint64, *entity.Url) error
	Create(*entity.Url) error
	Update(*entity.Url) error
	Delete(*entity.Url) error
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{
		DB: db,
	}
}

func (r *urlRepository) FindByShortUrl(shortUrl string, u *entity.Url) error {
	res := r.DB.Where("short_url = ?", shortUrl).First(&u)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r *urlRepository) FindById(id uint64, u *entity.Url) error {
	res := r.DB.First(&u, id)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r *urlRepository) Create(u *entity.Url) error {
	res := r.DB.Create(&u)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r *urlRepository) Update(u *entity.Url) error {
	res := r.DB.Save(&u)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r *urlRepository) Delete(u *entity.Url) error {
	res := r.DB.Delete(&u)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
