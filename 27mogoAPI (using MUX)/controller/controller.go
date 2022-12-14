package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoDB/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "<*Your mongo connection string*>"
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
func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMoives := getAllMovies()
	json.NewEncoder(w).Encode(allMoives)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Tyoe", "application/json")
	w.Header().Set("Allow-Control-Allow-Mehtods", "PUT")

	param := mux.Vars(r)
	updateOnemovie(param["id"])
	json.NewEncoder(w).Encode(param["id"])

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	param := mux.Vars(r)
	deletingOneMovie(param["id"])
	json.NewEncoder(w).Encode(param["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatoin/json")
	w.Header().Set("Allow-Control-Allow-Mehtods", "DELETE")
	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
