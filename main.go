package main

import (
	"fmt"
	"net/http"
)

func main() {
	ConfigLoad()
	cfg := AllConfig()
	ConnectDb(cfg.Db)
	AssignRoutes()
	fmt.Println("Server is starting......")
	http.ListenAndServe(":3000", nil)
}
