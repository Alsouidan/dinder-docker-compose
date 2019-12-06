package Models

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dog struct {
	ID             primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Breed          string               `json:"Breed" bson:"Breed"`
	Gender         string               `json:"Gender" bson:"Gender"`
	DateOfBirth    string               `json:"DateOfBirth" "bson:"DateOfBirth"`
	Name           string               `json:"Name" bson:"Name"`
	Colour         string               `json:"Colour" bson:"Colour"`
	Weight         float64              `json: "Weight" bson: "Weight"`
	Matched_IDs    []primitive.ObjectID `json:"Matched_IDs" bson:"Matched_IDs"`
	Matched_by_IDs []primitive.ObjectID `json:"Matched_by_IDs" bson:"Matched_by_IDs"`
	Rejected_IDs   []primitive.ObjectID `json:"Rejected_IDs" bson:"Rejected_IDs"`
	image_id       string               `json:"image_id" bson:"image_id"`
	Owner_id       primitive.ObjectID   `json:"Owner_id" bson:"Owner_id"`
}

// type Dogs []Dog
func FindDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	ou := vars["id"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id1, _ := url.PathUnescape(ou)
	collection := DB.Database("Dinder").Collection("Dogs")
	var dog Dog
	id2, _ := primitive.ObjectIDFromHex(id1)
	fmt.Println(id2)
	err := collection.FindOne(ctx, bson.M{"_id": id2}).Decode(&dog)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(dog)
}
func AllDogs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting dogd")
	w.Header().Set("content-type", "application/json")
	var dogs []Dog
	// DB.Database("Dinder").Collection("Dogs".find)
	collection := DB.Database("Dinder").Collection("Dogs")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var dog Dog
		cursor.Decode(&dog)
		dogs = append(dogs, dog)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(dogs)

}
func PostDogs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Posting")
	w.Header().Set("content-type", "application/json")
	var dog Dog
	_ = json.NewDecoder(r.Body).Decode(&dog)
	collection := DB.Database("Dinder").Collection("Dogs")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, dog)

	fmt.Println(dog)
	json.NewEncoder(w).Encode(result)
}
func UpdateDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var dog Dog
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	ou := vars["id"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	id1, _ := url.PathUnescape(ou)
	_ = json.NewDecoder(r.Body).Decode(&dog)
	collection := DB.Database("Dinder").Collection("Dogs")
	id2, _ := primitive.ObjectIDFromHex(id1)

	result, _ := collection.UpdateOne(ctx, bson.M{"_id": id2}, bson.M{"$set": dog})

	fmt.Println(dog)
	json.NewEncoder(w).Encode(result)
}
func DeleteDog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	ou := vars["id"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	id1, _ := url.PathUnescape(ou)

	collection := DB.Database("Dinder").Collection("Dogs")
	id2, _ := primitive.ObjectIDFromHex(id1)
	result, _ := collection.DeleteOne(ctx, bson.M{"_id": id2})
	json.NewEncoder(w).Encode(result)
}

func ApproveDog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Approving")
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	ou := vars["id"]
	ou1 := vars["id1"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id11, _ := url.PathUnescape(ou1)
	id1, _ := url.PathUnescape(ou)

	collection := DB.Database("Dinder").Collection("Dogs")
	id2, _ := primitive.ObjectIDFromHex(id1)
	id22, _ := primitive.ObjectIDFromHex(id11)
	fmt.Println(id2)
	fmt.Println(id22)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": id2}, bson.M{"$push": bson.M{"Matched_IDs": id22}})
	_, err1 := collection.UpdateOne(ctx, bson.M{"_id": id22}, bson.M{"$push": bson.M{"Matched_by_IDs": id2}})

	if err != nil || err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func RejectDog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Rejecting")
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	ou := vars["id"]
	ou1 := vars["id1"]

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id1, _ := url.PathUnescape(ou)
	id11, _ := url.PathUnescape(ou1)
	collection := DB.Database("Dinder").Collection("Dogs")
	id2, _ := primitive.ObjectIDFromHex(id1)
	id22, _ := primitive.ObjectIDFromHex(id11)
	fmt.Println(id2)
	fmt.Println(id22)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": id2}, bson.M{"$push": bson.M{"Rejected_IDs": id22}})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetReccomendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var dog Dog
	var dogs []Dog
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_ = json.NewDecoder(r.Body).Decode(&dog)
	collection := DB.Database("Dinder").Collection("Dogs")
	var gender string
	if dog.Gender == "Male" {
		gender = "Female"
	} else {
		gender = "Male"
	}
	fmt.Println(gender)
	fmt.Println(dog.Breed)
	cursor, err := collection.Find(ctx, bson.M{"Breed": dog.Breed, "Gender": gender,"_id":bson.M{"$nin":dog.Matched_IDs}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var dog Dog
		cursor.Decode(&dog)
		dogs = append(dogs, dog)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(dogs)

}
func Intersection(a, b []primitive.ObjectID) (c []primitive.ObjectID) {
	m := make(map[primitive.ObjectID]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
func GetMatches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var dogIDs []primitive.ObjectID
	var dogs []Dog
	vars := mux.Vars(r)
	ou := vars["id"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id1, _ := url.PathUnescape(ou)
	collection := DB.Database("Dinder").Collection("Dogs")
	var dog Dog
	id2, _ := primitive.ObjectIDFromHex(id1)
	fmt.Println(id2)
	err4 := collection.FindOne(ctx, bson.M{"_id": id2}).Decode(&dog)
	if err4 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err4.Error() + `" }`))
		return
	}
	dogIDs = Intersection(dog.Matched_IDs, dog.Matched_by_IDs)
	cursor, err := collection.Find(ctx, bson.M{"_id": bson.M{"$in": dogIDs}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var dog Dog
		cursor.Decode(&dog)
		dogs = append(dogs, dog)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(dogs)

}
