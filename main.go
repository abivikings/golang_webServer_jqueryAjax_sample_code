package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

var people []Person

type Person struct {
    ID string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
    Address *Address `json:"address,omitempty"`
}

type Address struct {
    Country string `json:"country,omitempty"`
    City string `json:"city,omitempty"`
}

func DefaultHandler(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprint(w, "Home")
}


func PeopleHandler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(people)
}

func PersonHandler(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Access-Control-Allow-Origin", "*")
     params := mux.Vars(r)
     for _,item := range people{
         if item.ID == params["id"]{
            json.NewEncoder(w).Encode(item)
         }
     }
}


func main()  {
    route := mux.NewRouter()
    people = append(people, Person{ID : "1", Name : "sojib", Address : &Address{Country : "bangladesh", City : "Dhaka"}})
    people = append(people, Person{ID : "2", Name : "hasan", Address : &Address{Country : "china", City : "C"}})
    people = append(people, Person{ID : "3", Name : "kamal", Address : &Address{Country : "japan", City : "J"}})
    route.HandleFunc("/", DefaultHandler).Methods("Get")
    route.HandleFunc("/people", PeopleHandler).Methods("Get")
    route.HandleFunc("/people/{id}", PersonHandler).Methods("Get")
    log.Fatal(http.ListenAndServe(":9000", route))
}
