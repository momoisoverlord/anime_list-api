package animes

import (
	"github.com/momoisoverlord/anime_list-api/utils/errors"
)

// Core Domain
type Anime struct {
	Id            int64  `json:"id"`
	Title         string `json:"title"`
	Cover         string `json:"cover"`
	Genre         string `json:"genre"`
	Description   string `json:"description"`
	AnimeType     string `json:"anime_type"`
	Studio        string `json:"studio"`
	ReleaseDate   string `json:"release_date"`
	Status        string `json:"status"`
	Language      string `json:"language"`
	AnimeUrl      string `json:"anime_url"`
	EpisodeUrl    string `json:"eps_url"`
	EpisodeNumber string `json:"eps_number"`
	DownloadLink  string `json:"download_link"`
	DateCreated   string `json:"date_created"`
}

func (anime *Anime) Validate() *errors.RestErr {
	//anime.Title = strings.ToLower(anime.Title)
	if anime.Title == "" {
		return errors.NewBadRequestError("invalid title")
	}
	if anime.EpisodeUrl == "" {
		return errors.NewBadRequestError("invalid episode url")
	}
	if anime.DownloadLink == "" {
		return errors.NewBadRequestError("invalid download url")
	}
	return nil
}
