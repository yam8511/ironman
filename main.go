package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	/// 編譯程式: go build -o program .
	/// 執行指令: APP=ironman ./program 1 2 3

	// cmd・第一式：取所有環境變數
	fmt.Println("all environments: ", os.Environ())
	// Output: all environments: [APP=ironman GOROOT=/usr/local/go SHELL=/bin/zsh ......忽略]

	// cmd・第二式：取指定的環境變數
	appInEnv := os.Getenv("APP")
	fmt.Println("APP in environments: ", appInEnv)
	// Output: APP in environments: ironman

	// cmd・第三式：取指令參數
	fmt.Println("all args: ", os.Args)
	// Output: all args:  [./program 1 2 3]

	if len(os.Args) < 2 { // 判斷參數除了執行指令本身，有沒有額外參數
		// 如果沒有額外參數，正常結束並提示指令
		fmt.Println("請輸入其中之一的參數： demo, server")
		os.Exit(0)
	}

	// 定義每個指令，要執行的動作
	switch os.Args[1] {
	case "demo":
		fmt.Println("這是一個示範指令")

	case "server":
		port := ":" + os.Getenv("PORT")
		log.Println("開始建立伺服器 ", port)
		log.Fatal(http.ListenAndServe(port, nil))

	default:
		fmt.Println("請輸入其中之一的參數： demo, server")
		os.Exit(1)
	}
}
