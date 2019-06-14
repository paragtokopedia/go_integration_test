package main


import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var databaseInstance *sql.DB

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	w.Write([]byte(`{"alive": true,"version":1}`))
}

func main() {
	log.Println("Starting Server at 80 ")
	InitDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", r))
}

func InitDatabase() {
	var err error
	databaseInstance, err = sql.Open("mysql", "root@localhost/test")
	if err !=nil{
		log.Fatal(err)
	}
}

func GetUserByID(ctx context.Context, id int64, db *sql.DB) ([]User, error) {
	results, err := db.QueryContext(ctx,"select * from user where id = ?",id)

	if err !=nil{
		return nil, err
	}
	var users []User
	for results.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&user.ID, &user.Name)
		if err != nil {
			return users, err
		}
		// and then print out the tag's Name attribute
		log.Printf(user.Name)
		users = append(users,user)
	}

	return users, err
}