package repository

import (
	"gorm.io/gorm"
)

type Url struct {
	ID       uint64
	ShortUrl string
	LongUrl  string
}

type UrlRepository interface {
	FindByShortUrl(shortUrl string, u *Url) (*Url, error)
	FindById(id uint64, u *Url) (*Url, error)
	Create(u *Url) (*Url, error)
	Update(u *Url) (*Url, error)
	Delete(u *Url) error
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &Repository{
		DB: db,
	}
}

func (repository *Repository) FindByShortUrl(shortUrl string, u *Url) (*Url, error) {
	res := repository.DB.Where("short_url = ?", shortUrl).First(&u)
	if err := res.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repository *Repository) FindById(id uint64, u *Url) (*Url, error) {
	res := repository.DB.First(&u, id)
	if err := res.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repository *Repository) Create(u *Url) (*Url, error) {
	res := repository.DB.Create(&u)
	if err := res.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repository *Repository) Update(u *Url) (*Url, error) {
	res := repository.DB.Save(&u)
	if err := res.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repository *Repository) Delete(u *Url) error {
	res := repository.DB.Delete(&u)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
