package main

import (
	"database/sql"
	"log"
	_ "mysql-master"
	"net/http"
	dbcon "usermanagement/db"
	router "usermanagement/routers"
)

func main() {
	var err error
	//dbcon.Db, err = sql.Open("mysql", "golive:ShdP6ScS9E4Dr2vS@tcp(localhost:3306)/test_db")
	dbcon.Db, err = sql.Open("mysql", "root:07081125@tcp(localhost:3307)/tes_db")

	if err != nil {
		log.Fatal(err.Error())
	}

	for key, val := range router.Routes {
		http.HandleFunc(key, val)
	}

	log.Println("Listen  and serve...")
	http.ListenAndServe(":9000", nil)
}
