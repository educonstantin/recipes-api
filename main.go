package main

import (
	"net/http"
	"recipes-api/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var recipes []model.Recipe

func init() {
	recipes = make([]model.Recipe, 0)
}

func NewRecipeHandler(c *gin.Context) {
	var recipe model.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.Run()
}
