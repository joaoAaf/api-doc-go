// Para criação de uma api:
// Definição de um router, que vai receber a requisição e encaminhar para um controler;
// Definição de uma função para cada controler, que vai processar a requisição e retornar a resposta;

package main

import (
	"net/http"

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

var albums = []Album{
	{Id: 0, Title: "Exemple Title", Artist: "Exemple Artist", Price: 0.00},
}

// "gin.Context" carrega os detalhes da solicitação,
// pode validar e serializar o JSON de uma solicitação e
// retornar uma resporta em JSON tambem.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run(":8080")
}
