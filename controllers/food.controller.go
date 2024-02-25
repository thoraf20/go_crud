package controller

import (
	"fmt"
	"time"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb/mongo-driver/mongo"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"gopkg.in/mgo.v2/bson"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

var validate = validator.New()

func GetFoods() gin.HandlerFunc() {
	return func(c *gin.Context){

	}
}

func GetFood() gin.HandlerFunc() {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.BackGround(), 100*time.Second)
		foodId := c.Param(food_id)
		var food models.Food

		foodCollection.FindOne(ctx, bson.M{"food_id", foodId}).Decoded(&food)
		defer(cancel)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc() {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.BackGround(), 100*time.Second)
		var food models.Food
		var menu models.Menu

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		validationErr := validate.Struct(&food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
			return 
		}

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer(cancel)
		if err != nil {
			msg := fmt.SPrintf("menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{ "error": msg })
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.now()).Format(time.RFC3339)
		food.Updated_at, _ = time.Parse(time.RFC3339, time.now()).Format(time.RFC3339)
		food.ID = primitive.NewObjectID()
		food.food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertError := foodCollection.InsertOne(ctx, food)
		if insertError != nil {
			msg := fmt.SPrintf("food item was not found")
			 c.JSON(http.StatusInternalServerError, gin.H{ "error": msg })
			return
		}
		defer(cancel)
		
		c.JSON(http.StatusOK, result)
	}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {
	
}

func UpdateFood() gin.HandlerFunc() {
	return func(c *gin.Context){

	}
}