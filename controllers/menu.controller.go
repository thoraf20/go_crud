package controller

import (
	"log"
	"context"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"go.mongodb/mongo-driver/mongo"
	"golang-restaurant-management/database"

)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandleFunc() {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.BackGround(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer(cancel)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H({"error": "error occured while listing menu items"}))
		}
		var allMenus = []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandleFunc() {
	return func(c *gin.Context) {
		
	}
}

func CreateMenu() gin.HandleFunc() {
	return func(c *gin.Context) {
		
	}
}

func UpdateMenu() gin.HandleFunc() {
	return func(c *gin.Context) {
		
	}
}