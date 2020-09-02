package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir("./public")))

	log.Println("開始建立伺服器")
	log.Fatal(http.ListenAndServe(":8000", router))
}
