package controller

import (
	"context"
	"fmt"
	"log"
	"mongoDB/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "<*your mongo connection string*>"
const dbName = "netflix"
const colName = "watchlist"

// very IMportant
var collection *mongo.Collection

// connect with mongoDB
func init() {
	//clint option
	clientOption := options.Client().ApplyURI(connectionString)

	//connection to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection successfull to MongoDB")
	collection = client.Database(dbName).Collection(colName)

	//collection instance is ready
	fmt.Println("Collectoin instsnce is ready")
}

//defining MongoDB helper methods they will only take the data and going to add that to the DB

//inserting one record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in the DB with ID:", inserted.InsertedID)

}

// Update one record
func updateOnemovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id} //we can also used bson,D here which is used to get the data in ordere format
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	//this will help us to find how many value are going to update
	//in this case we will get only one value
	fmt.Println("Modified count: ", result.ModifiedCount)
}

// deleting one value
func deletingOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with count :", deletecount)

}

//deelting all movies at once

func deleteAllMovies() int64 {
	//filter:= bson.M{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("No. of movies deleted:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//getting all the data from the database

func getAllMovies() []primitive.M {
	//cursor is an object which loops around the object to get the data
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

// controlers we are going to use
func GetAllMoviesController(c *gin.Context) {
	allMoives := getAllMovies()
	c.JSON(http.StatusOK, gin.H{"result": allMoives})
}

func CreateMovie(c *gin.Context) {
	var movie model.Netflix
	err := c.BindJSON(&movie)
	if err != nil {
		log.Fatal(err)
	}
	insertOneMovie(movie)
	c.JSON(http.StatusOK, gin.H{"result": movie})
}

func MarkAsWatched(c *gin.Context) {
	param := c.Param("id")
	updateOnemovie(param)
	c.JSON(http.StatusOK, gin.H{"result": param})
}

func DeleteOneMovie(c *gin.Context) {
	param := c.Param("id")
	deletingOneMovie(param)
	c.JSON(http.StatusOK, gin.H{"result": param})
}

func DeleteAllMovies(c *gin.Context) {
	count := deleteAllMovies()
	c.JSON(http.StatusOK, gin.H{"result": count})
}
