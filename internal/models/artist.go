package models

type Artist struct {
	ID          string
	Name        string
	Genres      []string
	Image       string
	Popularity  int
	ExternalUrl string
}
