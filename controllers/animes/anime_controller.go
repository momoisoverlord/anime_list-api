package animes

import (
	"github.com/gin-gonic/gin"
	"github.com/momoisoverlord/anime_list-api/domain/animes"
	"github.com/momoisoverlord/anime_list-api/services"
	"github.com/momoisoverlord/anime_list-api/utils/errors"
	"net/http"
	"strconv"
)

func getAnimeId(animeIdParam string) (int64, *errors.RestErr) {
	animeId, animeErr := strconv.ParseInt(animeIdParam, 10, 64)
	if animeErr != nil {
		return 0, errors.NewBadRequestError("anime id should be a number")
	}
	return animeId, nil
}

func Create(c *gin.Context) {
	var anime animes.Anime
	if err := c.ShouldBindJSON(&anime); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateAnime(anime)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	animeId, idErr := getAnimeId(c.Param("anime_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	anime, getErr := services.GetAnime(animeId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, anime)
}

func Update(c *gin.Context) {
	animeId, idErr := getAnimeId(c.Param("anime_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var anime animes.Anime
	if err := c.ShouldBindJSON(&anime); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	anime.Id = animeId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateAnime(isPartial, anime)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	animeId, idErr := getAnimeId(c.Param("anime_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteAnime(animeId); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted successfully"})
}

func Search(c *gin.Context) {
	title := c.Query("title")

	animes, err := services.Search(title)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, animes)
}
