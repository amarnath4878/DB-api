package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
func main(){
	router := mux.NewRouter()

	router.HandleFunc("/",HomePageHandler)
	http.ListenAndServe(":1119", router)
}
func HomePageHandler(w http.ResponseWriter, req *http.Request){
    fmt.Fprint(w, "amarnath")
}
