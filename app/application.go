package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	_ = router.Run()
	//_ = router.Run(":8080")

}
