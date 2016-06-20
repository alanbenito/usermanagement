package  main

import (
	"log"
	"net/http"
	router "usermanagement/routers"
)

func main() {
	for key, val := range router.Routes {
		http.HandleFunc(key, val)
	}

	log.Println("Listen  and serve...")
	http.ListenAndServe(":9000",nil)
}