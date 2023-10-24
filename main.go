package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type error struct { // наша структура ошибки
	Error string `json: "error"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Milligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
	//c.JSON(http.StatusOK, albums)

}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id") // получаем id из url

	//Loop over the list of albums, looking for
	//an album whose ID value matches the parameter

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, error{"not_found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, error{"not_found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Cal BindJSON to bind the receives JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, error{"bad_request"})
		return
	}

	// в этом коде используется контекст bindjson, чтобы спарсить
	// данные из запроса в переменную newAlbum

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum) //201
	//отвечаем 201 и возвращаем созданный абльбом

}

func main() {
	router := gin.Default()            // иниц роутера джин
	router.GET("/albums", getAlbums)   // гет для ассоциации пути с хэндлером
	router.POST("/albums", postAlbums) // опять биндим путь и хэндлер
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8080") // запускаем сервер
}
