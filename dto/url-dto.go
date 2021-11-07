package dto

type UrlResponseDTO struct {
	ID       uint64 `json:"id"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}

type UrlRequestDTO struct {
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}
