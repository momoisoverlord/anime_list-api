package app

import (
	"github.com/momoisoverlord/anime_list-api/controllers/animes"
	"github.com/momoisoverlord/anime_list-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/animes", animes.Create)
	router.GET("/animes/:anime_id", animes.Get)
	router.PUT("/animes/:anime_id", animes.Update)
	router.PATCH("/animes/:anime_id", animes.Update)
	router.DELETE("/animes/:anime_id", animes.Delete)
	router.GET("/internal/animes/search", animes.Search)
}
