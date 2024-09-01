package models

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}
type Domain struct {
	DomainURL  string
	VisitCount int
}
