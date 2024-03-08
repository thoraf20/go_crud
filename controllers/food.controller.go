package controller

import (
	"context"
	"fmt"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var foodCollection *mongo.Collection = database.Client.Database("restaurant").Collection("food")

var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		perPage, err := strconv.Atoi(c.Query(("perPage")))
		if err != nil || perPage < 1 {
			perPage = 10
		}

		page, err := strconv.Atoi(c.Query(("page")))
		if err != nil || page < 1 {
			perPage = 1
		}

		startIndex := (page - 1) * perPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{Name: "$match", Value: bson.D{{}}}}
		groupStage := bson.D{{Name: "$group", Value: bson.D{{Name: "_id", Value: bson.D{{Name: "_id", Value: "null"}}}, {Name: "total_count", Value: bson.D{{Name: "sum", Value: "1"}}}, {Name: "data", Value: bson.D{{Name: "$push", Value: "$$ROOT"}}}}}}
		projectStage := bson.D{
			{
				Name: "$project", Value: bson.D{
					{Name: "_id", Value: 0},
					{Name: "total_count", Value: 1},
					{Name: "food_items", Value: bson.D{{Name: "$slice", Value: []interface{}{"data", startIndex, perPage}}}},
				},
			},
		}

		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage
		})

		defer cancel()
		if err != nil{
			c.JSON(http.internalServerError, gin.H{"error": "error occured while listing foods"})
		}

		var allFoods []bson.M
		if err = result.All(ctx, &allFoods); err != nil{
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allFoods[0])
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		validationErr := validate.Struct(&food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()

		if err != nil {
			msg := fmt.Sprintf("menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*&food.Price, 2)
		food.Price = *&num

		result, insertError := foodCollection.InsertOne(ctx, food)
		if insertError != nil {
			msg := fmt.Sprintf("food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
		defer cancel()
	}
}

func round(num float64) int {
}

func toFixed(num float64, precision int) float64 {
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
