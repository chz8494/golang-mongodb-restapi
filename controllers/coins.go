package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"github.com/chz8494/golang-mongodb-restapi/config"
	"github.com/chz8494/golang-mongodb-restapi/models"
	"gopkg.in/go-playground/validator.v9"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/matryer/respond.v1"
	"strconv"
)

// connect db
var database = config.ConnectDB()

// CreateCoin : This is create coin method
func CreateCoin(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	var coin models.Coin

	// we decode our body request params
	json.NewDecoder(request.Body).Decode(&coin)

	//validation for empty fields

	validate := validator.New()

	err := validate.Struct(coin)

	if err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}
	//set time out
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//cancel to prevent memory leakage
	defer cancel()

	// insert our book model.
	collection := database.Collection("BTC|WOW")
	result, err := collection.InsertOne(ctx, coin)

	if err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}

	data := map[string]interface{}{"data": result, "message": "Success", "status": http.StatusOK}
	respond.With(response, request, http.StatusOK, data)

}

//GetCoin : Get coin by id
func GetCoin(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	//request params
	params := mux.Vars(request)
	var coin models.Coin

	//set time out
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//cancel to prevent memory leakage
	collection := database.Collection(params["coin"])
	defer cancel()
	//query the model
	cursor, err := collection.Find(ctx, bson.M{})
	var coins []bson.M
	err = cursor.All(ctx, &coins)
	//handle error
	if err != nil {

		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}
	//handle success
	data := map[string]interface{}{"data": coins, "message": "Success", "status": http.StatusOK}
	respond.With(response, request, http.StatusOK, data)
}

func GetCoin_Timestamp(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	//request params
	params := mux.Vars(request)
	t, err := strconv.ParseInt(params["timestamp"], 10, 64)
	if err == nil {
	    fmt.Printf("%d of type %T", t, t)
	}
	/*convert id to usable mongodb object id
	id, errorID := primitive.ObjectIDFromHex(params["Timestamp"])
	if errorID != nil {
		data := map[string]interface{}{"data": nil, "message": errorID.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}
	*/
	var coin models.Coin

	//set time out
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//cancel to prevent memory leakage
	collection := database.Collection(params["coin"])
	defer cancel()
	//query the model
	err := collection.FindOne(ctx, models.Coin{Timestamp: t}).Decode(&coin)

	//handle error
	if err != nil {
		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}
	// handle success data
	data := map[string]interface{}{"data": coin, "message": "Success", "status": http.StatusOK}
	respond.With(response, request, http.StatusOK, data)
}

//GetCoins : Get all coins
func GetCoins(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")


	//set time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//cancel to prevent memory leakage
	defer cancel()

	//query data
	cursor, err := database.ListCollections(ctx, bson.D{})
	var collections []bson.M
	err = cursor.All(ctx, &collections)
	//handle error
	if err != nil {

		data := map[string]interface{}{"data": nil, "message": err.Error(), "status": http.StatusInternalServerError}
		respond.With(response, request, http.StatusInternalServerError, data)
		return
	}
	//handle success
	data := map[string]interface{}{"data": collections, "message": "Success", "status": http.StatusOK}
	respond.With(response, request, http.StatusOK, data)
}
