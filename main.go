package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
) 

// Structure to hold info on albums
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}


var albums = [] album{

	{ID: "1", Title:"Blue Train", Artist: "John", Price: 45.00},
	{ID: "2", Title:"Gardenn", Artist: "Clifford", Price: 89.00},
	{ID: "3", Title:"Red Velvet", Artist: "Alstyn", Price: 33.00},
	{ID: "5", Title:"Sarah Vaugh", Artist: "Mulligan", Price: 21.00},

}

func getAlbums(c *gin.Context) {
	// Return the http status 200 along with the entire album array as JSON
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums (c *gin.Context){
	var newAlbum album;
	if err:=c.BindJSON(&newAlbum) ; err !=nil {
		return
	}
	// Add the recieved JSON album item to the existing array 
	albums = append(albums, newAlbum)
	// return a http status created along with the new appended album
	c.IndentedJSON(http.StatusCreated, newAlbum)
	
}

// function to get the albums by ID passed in the path parameter
func getAlbumByID(c *gin.Context){
	id := c.Param("id")

	//Iterating over the albums array to find the album based on the matching ID
	for _, item:= range albums{
		if item.ID == id {
			// Return a status OK/200 along with the matched album details
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	// Else, return a message not found in the header
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})
}

// Entry point of the program
func main(){
	// Initializing the http gin middleware
	router := gin.Default()

	// Registering http calls with the gin http router
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	// Allow the http router to listen on a local host address on port 8080
	router.Run("localhost:8080")
}