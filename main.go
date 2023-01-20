package main

import (
	"fmt"
	"net/http"

	"demo/src"
)

func main() {
	src.ConfigLoad()
	cfg := src.AllConfig()
	src.ConnectDb(cfg.Db)
	src.AssignRoutes()
	fmt.Println("Server is starting......")
	http.ListenAndServe(":3000", nil)
}
