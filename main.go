package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	configFile := os.Getenv("CONF_FILE")
	f, err := ioutil.ReadFile("config/" + configFile + ".json")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f))

	var configData struct {
		App  string `json:"app"`
		Name string `json:"Name"`
		Port int    `json:"port"`
		DB   struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"db"`
	}

	err = json.Unmarshal(f, &configData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Project Name", configData.Name)

	router := http.NewServeMux()
	router.HandleFunc("/api-demo", func(w http.ResponseWriter, r *http.Request) {
		// 使用配置檔．第一式：在API加上
		w.Write([]byte("目前專案名稱是" + configData.Name))
	})

	// 使用配置檔．第二式：決定伺服器的Port
	addr := fmt.Sprintf(":%d", configData.Port)
	log.Println("開始建立伺服器 ", addr)
	log.Fatal(http.ListenAndServe(addr, router))

	// 使用配置檔．第三式：決定DB連線的帳密或IP
	_ = fmt.Sprintf("connect database: username=%s password=%s",
		configData.DB.Username,
		configData.DB.Password,
	)
}
