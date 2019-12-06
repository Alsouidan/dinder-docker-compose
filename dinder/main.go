package main

import (
	"dinder/Models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
    "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Configuration struct {
	Port     string `json: "Port"`
	Dbname   string `json: "dbname"`
	Username string `json: "Username"`
	Password string `json: "Password"`
}

var Config Configuration

func LoadConfiguration(file string) Configuration {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&Config)
	return Config
}
func main() {
	var mongoURI string 
	Config = LoadConfiguration("config.json")
if (Configuration{})!=Config{
	mongoURI = "mongodb+srv://" + Config.Username + ":" + Config.Password + "@" + Config.Dbname +
	"-uqqlp.mongodb.net/test?retryWrites=true&w=majority"

	}else {
		mongoURI = "mongodb+srv://" + os.Getenv("username") + ":" + os.Getenv("password") + "@" + os.Getenv("dbname")+
	"-uqqlp.mongodb.net/test?retryWrites=true&w=majority"
	}
	// "&replicaSet=dinder-shard-00-01-uqqlp.mongodb.net"
	// mongoURI:="mongodb://database:28017/?compressors=disabled&gssapiServiceName=mongodb"
	fmt.Println("connection string is:", mongoURI)
	Models.ExampleNewClient()
	Models.InitDB(mongoURI)
	handler()
}

func handler() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/dogs", Models.AllDogs).Methods("GET")
	myRouter.HandleFunc("/dogs/{id}", Models.FindDog).Methods("GET")
	myRouter.HandleFunc("/dogs", Models.PostDogs).Methods("Post")
	myRouter.HandleFunc("/dogs/{id}", Models.DeleteDog).Methods("Delete")
	myRouter.HandleFunc("/dogs/{id}", Models.UpdateDog).Methods("Put")
	myRouter.HandleFunc("/dogs/approve/{id}/{id1}", Models.ApproveDog).Methods("Put")
	myRouter.HandleFunc("/dogs/reject/{id}/{id1}", Models.RejectDog).Methods("Put")
	myRouter.HandleFunc("/dogs/getRec", Models.GetReccomendations).Methods("Post")
	myRouter.HandleFunc("/dogs/getMatches/{id}", Models.GetMatches).Methods("Get")
	myRouter.HandleFunc("/users", Models.AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", Models.FindUser).Methods("GET")
	myRouter.HandleFunc("/users", Models.PostUsers).Methods("Post")
	myRouter.HandleFunc("/users/{id}", Models.DeleteUser).Methods("Delete")
	myRouter.HandleFunc("/users/{id}", Models.UpdateUser).Methods("Put")
	myRouter.HandleFunc("/users/addDog/{id}", Models.AddDog).Methods("Put")
	myRouter.HandleFunc("/users/login", Models.Login).Methods("Post")
	myRouter.HandleFunc("/users/getDogs/{id}", Models.GetDogs).Methods("Get")
if (Configuration{})==Config{
    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myRouter)))


}else{
    log.Fatal(http.ListenAndServe(":"+Config.Port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myRouter)))
}
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage endpoint hit")
	enableCors(&w)
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
