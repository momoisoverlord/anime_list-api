package services

import (
	"github.com/momoisoverlord/anime_list-api/domain/animes"
	"github.com/momoisoverlord/anime_list-api/utils/date_utils"
	"github.com/momoisoverlord/anime_list-api/utils/errors"
)

func GetAnime(animeId int64) (*animes.Anime, *errors.RestErr) {
	result := &animes.Anime{Id: animeId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateAnime(anime animes.Anime) (*animes.Anime, *errors.RestErr) {
	if err := anime.Validate(); err != nil {
		return nil, err
	}

	anime.DateCreated = date_utils.GetNowDBFormat()
	if err := anime.Save(); err != nil {
		return nil, err
	}

	return &anime, nil
}

func UpdateAnime(isPartial bool, anime animes.Anime) (*animes.Anime, *errors.RestErr) {
	current, err := GetAnime(anime.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if anime.Title != "" {
			current.Title = anime.Title
		}

		if anime.Cover != "" {
			current.Cover = anime.Cover
		}

		if anime.Genre != "" {
			current.Genre = anime.Genre
		}

		if anime.Description != "" {
			current.Description = anime.Description
		}

		if anime.AnimeType != "" {
			current.AnimeType = anime.AnimeType
		}

		if anime.Studio != "" {
			current.Studio = anime.Studio
		}

		if anime.ReleaseDate != "" {
			current.ReleaseDate = anime.ReleaseDate
		}

		if anime.Status != "" {
			current.Status = anime.Status
		}

		if anime.Language != "" {
			current.Language = anime.Language
		}

		if anime.AnimeUrl != "" {
			current.AnimeUrl = anime.AnimeUrl
		}

		if anime.EpisodeUrl != "" {
			current.EpisodeUrl = anime.EpisodeUrl
		}

		if anime.EpisodeNumber != "" {
			current.EpisodeNumber = anime.EpisodeNumber
		}

		if anime.DownloadLink != "" {
			current.DownloadLink = anime.DownloadLink
		}
	} else {
		current.Title = anime.Title
		current.Cover = anime.Cover
		current.Genre = anime.Genre
		current.Description = anime.Description
		current.AnimeType = anime.AnimeType
		current.Studio = anime.Studio
		current.ReleaseDate = anime.ReleaseDate
		current.Status = anime.Status
		current.Language = anime.Language
		current.AnimeUrl = anime.AnimeUrl
		current.EpisodeUrl = anime.EpisodeUrl
		current.EpisodeNumber = anime.EpisodeNumber
		current.DownloadLink = anime.DownloadLink
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteAnime(animeId int64) *errors.RestErr {
	anime := &animes.Anime{Id: animeId}
	return anime.Delete()
}

func Search(title string) ([]animes.Anime, *errors.RestErr) {
	dao := &animes.Anime{}
	return dao.FindByTitle(title)
}
