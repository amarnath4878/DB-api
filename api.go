package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"pkg/database"

	"github.com/gorilla/mux"
)

var dbConn *sql.DB

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Mobile    int64  `json:"mobileNumber"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func init() {
	db := database.New("localhost", "root", "Shravan@123#", "8808", "gouser")
	dbConn = db.Connect()
	errPing := dbConn.Ping()
	if errPing != nil {
		log.Println(errPing)
	}
}

func main() {
	router := mux.NewRouter()

	// Route handles & endpoints

	router.HandleFunc("/", HomePageHandler).Methods(http.MethodGet)

	// Create a movie
	router.HandleFunc("/Signup", SignUpHandler).Methods(http.MethodPost)

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	SearchKey := r.URL.Query().Get("q")
	fmt.Fprint(w, SearchKey)
}
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user *User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	errJson := json.Unmarshal(body, &user)
	if errJson != nil {
		log.Fatalln(errJson)
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, user)
}