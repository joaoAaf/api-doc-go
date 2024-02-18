package controler

import (
	"api-doc-go/dto"
	"api-doc-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "gin.Context" carrega os detalhes da solicitação,
// pode validar e serializar o JSON de uma solicitação e
// retornar uma resporta em JSON tambem.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.ViewAlbums())
}

func GetAlbumById(c *gin.Context) {
	// o "c.Param" retorna o parametro
	// enviado pelo usuario por meio da url
	album, _ := service.VerifyId(c.Param("id"))
	if album != nil {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PutAlbumById(c *gin.Context) {
	var albumDto dto.AlbumDTO
	album, _ := service.VerifyId(c.Param("id"))
	if err := c.BindJSON(&albumDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if album != nil {
		service.UpdateAlbum(album, albumDto)
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbum(c *gin.Context) {
	var albumDto dto.AlbumDTO
	// "c.BindJSON" recebe o JSON do corpo da requisição
	// e converte para newAlbum
	if err := c.BindJSON(&albumDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	album := service.ConvertAlbum(albumDto)
	service.InsertAlbum(album)
	c.IndentedJSON(http.StatusCreated, album)
}

func DeleteAlbumById(c *gin.Context) {
	_, index := service.VerifyId(c.Param("id"))
	if index >= 0 {
		service.DeleteAlbum(index)
		c.IndentedJSON(http.StatusNoContent, "")
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
