package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/api-demo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("你的名字是" + r.URL.Query().Get("name")))
	})

	log.Println("開始建立伺服器")
	log.Fatal(http.ListenAndServe(":8000", router))
}
