// Para criação de uma api:
// Definição de um router, que vai receber a requisição e encaminhar para um controler;
// Definição de uma função para cada controler, que vai processar a requisição e retornar a resposta;

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Album struct {
	// A notação após a declaração da variavel,
	// serializa dos dados da struct para o formato de JSON
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type AlbumDTO struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var lastId int
var albums = []Album{
	{Id: addId(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: addId(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: addId(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func addId() int {
	lastId += 1
	return lastId
}

func (a *Album) ConvertAlbum(albumDto AlbumDTO) {
	a.Id = addId()
	a.Title = albumDto.Title
	a.Artist = albumDto.Artist
	a.Price = albumDto.Price
}

// "gin.Context" carrega os detalhes da solicitação,
// pode validar e serializar o JSON de uma solicitação e
// retornar uma resporta em JSON tambem.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	// o "c.Param" retorna o parametro
	// enviado pelo usuario por meio da url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	for _, a := range albums {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var albumDto AlbumDTO
	var album Album
	// "c.BindJSON" recebe o JSON do corpo da requisição
	// e converte para newAlbum
	if err := c.BindJSON(&albumDto); err != nil {
		return
	}
	album.ConvertAlbum(albumDto)
	albums = append(albums, album)
	c.IndentedJSON(http.StatusCreated, album)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.Run(":8080")
}
