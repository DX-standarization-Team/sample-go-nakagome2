package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// user page
type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func getUsers() []*User {

	// jst, err := time.LoadLocation("Asia/Tokyo")
	// if err != nil {
	// 	// エラーハンドリング
	// 	panic(err)
	// }
	// c := mysql.Config{
	// 	DBName:    "sample_database",
	// 	User:      "testuser",
	// 	Passwd:    "secret",
	// 	Addr:      "db:3306",
	// 	Net:       "tcp",
	// 	ParseTime: true,
	// 	Collation: "utf8mb4_unicode_ci",
	// 	Loc:       jst,
	// }
	// db, err := sql.Open("mysql", c.FormatDSN())
	
	// Connect DB
	db, err := sql.Open("mysql", "testuser:secret@tcp(db:3306)/sample_database")
	if err != nil {
		panic(err)
	}
	
	defer db.Close()

	// Execute query
	fmt.Printf("Execute query")
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

func userPage(w http.ResponseWriter, r *http.Request){
	users := getUsers()
	fmt.Printf("Hit the userPage endpoint")

	json.NewEncoder(w).Encode(users)
}

func main(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}