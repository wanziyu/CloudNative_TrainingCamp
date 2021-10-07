package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/logger", logHandler)
	http.HandleFunc("/healthz", healthzHandler)
	err := http.ListenAndServe(":30080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func headerHandler(w http.ResponseWriter, r *http.Request) {

	ua := r.UserAgent()

	w.Header().Set("Server", "Nginx-1.12.1")

	w.Header().Set("User-Agent", ua)

	w.Write([]byte("Add request header"))
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	version := runtime.Version()
	w.Header().Set("Version: ", version)
	w.Write([]byte("Golang Version is " + version))
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	//客户端IP-请求方法-返回码
	
	fmt.Print(r.RemoteAddr, "-")
	fmt.Print("-", r.Method)
	fmt.Print("-", http.StatusAccepted)
	w.Write([]byte("This is Server's log"))
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Status Code: 200"))
}
