package main

import (
    "net/http"

    "github.com/gin-gonic/gin"

	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"

    "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)


func main() {
	tracer.Start()
	defer tracer.Stop() 
	
    router := gin.Default()

	router.Use(gintrace.Middleware("gowith"))

    router.GET("/albums", getAlbums)

    router.Run("0.0.0.0:8080")


}

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

