package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"api.go/pkg/database"
	"github.com/gorilla/mux"
)

var dbConn *sql.DB

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Mobile    string `json:"mobileNumber"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func init() {
	db := database.New("localhost", "root", "Amarnarh99@", "3306", "table")
	dbConn = db.Connect()
	errPing := dbConn.Ping()
	if errPing != nil {
		log.Println(errPing)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomePageHandler).Methods(http.MethodGet)
	router.HandleFunc("/SignUp", SignUpHandler).Methods(http.MethodPost)
	fmt.Println("Server at 1119")
	log.Fatal(http.ListenAndServe(":1119", router))
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
		log.Println(errJson)
	}
	stmt, _ := dbConn.Prepare("INSERT INTO employee(firstName,lastName,email,mobileNumber,password) values(?,?,?,?,?)")
	result, _ := stmt.Exec(&user.FirstName, &user.LastName, &user.Email, &user.Mobile, &user.Password)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, result)
}
