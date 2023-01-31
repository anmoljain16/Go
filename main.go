package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type hehe struct {
	ID       string `json:"id"`
	Item     string `json:"item"`
	Complete bool   `json:"complete"`
}

var hehes = []hehe{
	{ID: "1", Item: "Fuck you", Complete: true},
	{ID: "2", Item: "Fuck you two time", Complete: true},
	{ID: "3", Item: "Fuck you three times", Complete: true},
}

func gethehes(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, hehes)
}

func addhehehe(context *gin.Context) {
	var newhehe hehe
	if err := context.BindJSON(&newhehe); err != nil {
		return
	}

	hehes = append(hehes, newhehe)

	context.IndentedJSON(http.StatusCreated, newhehe)
}

func getTodobyid(idjson string) (*hehe, error) {
	for i, t := range hehes {
		if t.ID == idjson {
			return &hehes[i], nil
		}
	}
	return nil, errors.New("hehe not found")
}

func gethehe(context *gin.Context) {
	id := context.Param("id")
	hehe, err := getTodobyid(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, hehe)
}
func getToggle(context *gin.Context) {
	id := context.Param("id")
	hehe, err := getTodobyid(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
		return
	}

	hehe.Complete = !hehe.Complete

	context.IndentedJSON(http.StatusOK, hehe)
}

func main() {

	router := gin.Default()
	router.GET("/hehe", gethehes)
	router.GET("/hehe/:id", gethehe)
	router.PATCH("/hehe/:id", getToggle)
	router.POST("/hehe", addhehehe)
	router.Run("localhost:6969")

}
