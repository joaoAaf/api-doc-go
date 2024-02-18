// Para criação de uma api:
// Definição de um router, que vai receber a requisição e encaminhar para um controler;
// Definição de uma função para cada controler, que vai processar a requisição e retornar a resposta;

package main

import (
	"api-doc-go/entity"
	"api-doc-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "gin.Context" carrega os detalhes da solicitação,
// pode validar e serializar o JSON de uma solicitação e
// retornar uma resporta em JSON tambem.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.ViewAlbums())
}

func getAlbumById(c *gin.Context) {
	// o "c.Param" retorna o parametro
	// enviado pelo usuario por meio da url
	album, _ := service.VerifyId(c.Param("id"))
	if album != nil {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func putAlbumById(c *gin.Context) {
	var albumDto entity.AlbumDTO
	album, _ := service.VerifyId(c.Param("id"))
	if err := c.BindJSON(&albumDto); err != nil {
		return
	}
	if album != nil {
		service.UpdateAlbum(album, albumDto)
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var albumDto entity.AlbumDTO
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

func deleteAlbumById(c *gin.Context) {
	_, index := service.VerifyId(c.Param("id"))
	if index >= 0 {
		service.DeleteAlbum(index)
		c.IndentedJSON(http.StatusNoContent, "")
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.PUT("/albums/:id", putAlbumById)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.Run(":8080")
}
