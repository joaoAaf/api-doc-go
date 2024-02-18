package router

import (
	"api-doc-go/controler"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/albums", controler.GetAlbums)
	router.GET("/albums/:id", controler.GetAlbumById)
	router.PUT("/albums/:id", controler.PutAlbumById)
	router.POST("/albums", controler.PostAlbum)
	router.DELETE("/albums/:id", controler.DeleteAlbumById)
	router.Run(":8080")
}
