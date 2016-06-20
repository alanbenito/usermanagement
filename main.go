package  main

import (
	"log"
	"net/http"
	router "usermanagement/routers"
	"database/sql"
	_"mysql-master"
	
)

func main() {
	//create connection string
	db, err := sql.Open("mysql", "root:07081125@tcp(localhost:3307)/tes_db")
	if err != nil{
		log.Fatal("fatal error connection ->  ", err.Error())
		return
	}
	defer db.Close()
	err = db.Ping()
	log.Println("connect to DB succesfully")
	
	for key, val := range router.Routes {
		http.HandleFunc(key, val)
	}

	log.Println("Listen  and serve...")
	http.ListenAndServe(":9000",nil)
}