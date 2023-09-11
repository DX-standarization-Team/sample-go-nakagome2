package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// user page
type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func getUsers() []*User {
	db, err := sql.Open("mysql", "testuser:secret@tcp(db:30306)/sample_database")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//Execute query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	var users []*User
	for results.Next(){
		var u User
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}
		users = append(users, &u)
	}

	return users
	
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcom")
	fmt.Printf("Hit the homepage endpoint")
}

func usersPage(w http.ResponseWriter, r *http.Request){
	users := getUsers()
	fmt.Printf("Hit the usersPage endpoint")

	json.NewEncoder(w).Encode(users)
}

func main(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe("8080", nil))
}