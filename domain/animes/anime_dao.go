package animes

import (
	"fmt"
	"github.com/momoisoverlord/anime_list-api/datasources/mysql/animes_db"
	"github.com/momoisoverlord/anime_list-api/utils/errors"
	"github.com/momoisoverlord/anime_list-api/utils/mysql_utils"
)

const (
	//indexUniqueEpisodeUrl  = "animes.eps_url_UNIQUE"
	//indexUniqueDownloadUrl = "animes.download_link_UNIQUE"
	//errorNoRows            = "no rows in result set"
	queryInsertAnime      = "INSERT INTO animes(title, cover, genre, description, anime_type, studio, release_date, status, language, anime_url, eps_url, eps_number, download_link, date_created) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	queryGetAnime         = "SELECT id, title, cover, genre, description, anime_type, studio, release_date, status, language, anime_url, eps_url, eps_number, download_link, date_created FROM animes WHERE id=?;"
	queryUpdateAnime      = "UPDATE animes SET title=?, cover=?, genre=?, description=?, anime_type=?, studio=?, release_date=?, status=?, language=?, anime_url=?, eps_url=?, eps_number=?, download_link=? WHERE id=?;"
	queryDeleteAnime      = "DELETE FROM animes WHERE id=?;"
	queryFindAnimeByTitle = "SELECT id, title, cover, genre, description, anime_type, studio, release_date, status, language, anime_url, eps_url, eps_number, download_link, date_created FROM animes WHERE title=?;"
)

func (anime *Anime) Get() *errors.RestErr {
	stmt, err := animes_db.Client.Prepare(queryGetAnime)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(anime.Id)
	if getErr := result.Scan(&anime.Id, &anime.Title, &anime.Cover, &anime.Genre, &anime.Description, &anime.AnimeType, &anime.Studio, &anime.ReleaseDate, &anime.Status, &anime.Language, &anime.AnimeUrl, &anime.EpisodeUrl, &anime.EpisodeNumber, &anime.DownloadLink, &anime.DateCreated); getErr != nil {
		//if strings.Contains(getErr.Error(), errorNoRows) {
		//	return errors.NewNotFoundError(fmt.Sprintf("anime with the id: %d not found", anime.Id))
		//}
		//return errors.NewInternalServerError(fmt.Sprintf("error when trying to get anime with the id: %d. Reason: %s", anime.Id, getErr.Error()))
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (anime *Anime) Save() *errors.RestErr {
	stmt, err := animes_db.Client.Prepare(queryInsertAnime)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(
		anime.Title,
		anime.Cover,
		anime.Genre,
		anime.Description,
		anime.AnimeType,
		anime.Studio,
		anime.ReleaseDate,
		anime.Status,
		anime.Language,
		anime.AnimeUrl,
		anime.EpisodeUrl,
		anime.EpisodeNumber,
		anime.DownloadLink,
		anime.DateCreated,
	)
	if saveErr != nil {
		//sqlErr, ok := saveErr.(*mysql.MySQLError)
		//if !ok {
		//	return errors.NewInternalServerError(fmt.Sprintf("error when trying to save anime: %s", saveErr.Error()))
		//}
		//switch sqlErr.Number {
		//case 1062:
		//	if strings.Contains(sqlErr.Message, indexUniqueEpisodeUrl) {
		//		return errors.NewBadRequestError(fmt.Sprintf("Episode URL: %s already exists", anime.EpisodeUrl))
		//	}
		//	if strings.Contains(sqlErr.Message, indexUniqueDownloadUrl) {
		//		return errors.NewBadRequestError(fmt.Sprintf("Download URL: %s already exists", anime.DownloadLink))
		//	}
		//}
		//return errors.NewInternalServerError(fmt.Sprintf("error when trying to save anime: %s", saveErr.Error()))

		return mysql_utils.ParseError(saveErr)
	}

	animeId, err := insertResult.LastInsertId()
	if err != nil {
		//return errors.NewInternalServerError(fmt.Sprintf("error when trying to save anime: %s", err.Error()))
		return mysql_utils.ParseError(err)
	}
	anime.Id = animeId
	return nil
}

func (anime *Anime) Update() *errors.RestErr {
	stmt, err := animes_db.Client.Prepare(queryUpdateAnime)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(anime.Title, anime.Cover, anime.Genre, anime.Description, anime.AnimeType, anime.Studio, anime.ReleaseDate, anime.Status, anime.Language, anime.AnimeUrl, anime.EpisodeUrl, anime.EpisodeNumber, anime.DownloadLink, anime.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (anime *Anime) Delete() *errors.RestErr {
	stmt, err := animes_db.Client.Prepare(queryDeleteAnime)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(anime.Id); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (anime *Anime) FindByTitle(title string) ([]Anime, *errors.RestErr) {
	stmt, err := animes_db.Client.Prepare(queryFindAnimeByTitle)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(title)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]Anime, 0)
	for rows.Next() {
		var anime Anime
		if err := rows.Scan(&anime.Id, &anime.Title, &anime.Cover, &anime.Genre, &anime.Description, &anime.AnimeType, &anime.Studio, &anime.ReleaseDate, &anime.Status, &anime.Language, &anime.AnimeUrl, &anime.EpisodeUrl, &anime.EpisodeNumber, &anime.DownloadLink, &anime.DateCreated); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, anime)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no anime matching the title: %s", title))
	}
	return results, nil
}
