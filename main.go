package main

type Album struct {
	// A notação após a declaração da variavel,
	// serializa dos dados da struct para o formato de JSON
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
