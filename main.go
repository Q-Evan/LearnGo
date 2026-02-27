package main
import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "我是 Go 云原生小服务！\n")
}

func main(){
	http.HandleFunc("/", hello)
	fmt.Println("服务启动在 :8080")
	http.ListenAndServe(":8080", nil)
}