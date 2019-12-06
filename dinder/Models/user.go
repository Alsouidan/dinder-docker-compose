package Models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"net/url"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Gender      string               `json:"Gender" bson:"Gender"`
	DateOfBirth string               `json:"DateOfBirth" "bson:"DateOfBirth"`
	Name        string               `json:"Name" bson:"Name"`
	Email       string               `json:"Email" bson:"Email"`
	Password    string               `json:"Password" bson:"Password"`
	DogArray    []primitive.ObjectID `json:"DogArray" bson:"DogArray"`
}
type DogID struct {
	ID string `json:"_id" bson:"_id"`
}
type Claims struct {
	ID string `json:"ID"`

	Email    string `json:"Email"`
	Password string `json:"Password" bson:"Password"`

	jwt.StandardClaims
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	ou := vars["id"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	id1, _ := url.PathUnescape(ou)

	collection := DB.Database("Dinder").Collection("Users")
	var user User
	id2, _ := primitive.ObjectIDFromHex(id1)
	fmt.Println(id2)
	err := collection.FindOne(ctx, bson.M{"_id": id2}).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(user)
}
func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting Users")
	w.Header().Set("content-type", "application/json")
	var users []User
	// DB.Database("Dinder").Collection("Dogs".find)
	collection := DB.Database("Dinder").Collection("Users")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(users)

}
func PostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Posting")
	w.Header().Set("content-type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	collection := DB.Database("Dinder").Collection("Users")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, user)

	fmt.Println(user)
	json.NewEncoder(w).Encode(result)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user User
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	ou := vars["id"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	id1, _ := url.PathUnescape(ou)
	_ = json.NewDecoder(r.Body).Decode(&user)
	collection := DB.Database("Dinder").Collection("Users")
	id2, _ := primitive.ObjectIDFromHex(id1)

	result, _ := collection.UpdateOne(ctx, bson.M{"_id": id2}, bson.M{"$set": user})

	fmt.Println(user)
	json.NewEncoder(w).Encode(result)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	ou := vars["id"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	id1, _ := url.PathUnescape(ou)

	collection := DB.Database("Dinder").Collection("Users")
	id2, _ := primitive.ObjectIDFromHex(id1)
	result, _ := collection.DeleteOne(ctx, bson.M{"_id": id2})
	json.NewEncoder(w).Encode(result)
}
func AddDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var dogID DogID
	vars := mux.Vars(r)
	ou := vars["id"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id1, _ := url.PathUnescape(ou)
	_ = json.NewDecoder(r.Body).Decode(&dogID)
	id := dogID.ID
	collection := DB.Database("Dinder").Collection("Users")
	id2, _ := primitive.ObjectIDFromHex(id1)
	id3, _ := primitive.ObjectIDFromHex(id)
	result, _ := collection.UpdateOne(ctx, bson.M{"_id": id2}, bson.M{"$push": bson.M{"DogArray": id3}})
	json.NewEncoder(w).Encode(result)
}
func GetDogs(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var dogIDs []primitive.ObjectID
	var dogs []Dog
	vars := mux.Vars(r)
	ou := vars["id"]
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	id1, _ := url.PathUnescape(ou)
	collection := DB.Database("Dinder").Collection("Users")
	var user User
	id2, _ := primitive.ObjectIDFromHex(id1)
	fmt.Println(id2)
	err4 := collection.FindOne(ctx, bson.M{"_id": id2}).Decode(&user)
	if err4 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err4.Error() + `" }`))
		return
	}
	dogCollection := DB.Database("Dinder").Collection("Dogs")
	dogIDs=user.DogArray
	cursor, err := dogCollection.Find(ctx, bson.M{"_id":bson.M{"$in":dogIDs}})
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
func Login(w http.ResponseWriter, r *http.Request) {
	var jwtKey = []byte("my_secret_key")

	w.Header().Set("content-type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := DB.Database("Dinder").Collection("Users")
	err := collection.FindOne(ctx, bson.M{"Password": user.Password, "Email": user.Email}).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	expirationTime := time.Now().Add(30 * time.Minute)
	fmt.Println(user)
	fmt.Println(user.Password)
	claims := &Claims{
		Email: user.Email,
		ID:    user.ID.Hex(),
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	w.Write([]byte(`{ "token": "` + tokenString + `" }`))
}

// func FindDogs(w http.ResponseWriter, r *http.Request)
// {
// 	collection := DB.Database("Dinder").Collection("Users")

// 	cursor, err := collection.Find(ctx, bson.M{DogIDs:bson.M{"$in":"DogArray"}})
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var user User
// 		//cursor.Decode(&User)
// 		//dogs = append(dogs, dog)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	json.NewEncoder(w).Encode(dogs)

// }
