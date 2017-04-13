package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

)

type JsonBean struct {
	Name    int         `json:"name"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewJsonBean() *JsonBean {
	return &JsonBean{}

}

func sayyes(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	//fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("path", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))

	}
	fmt.Fprintf(w, "hi,Banana Lee")

}
func UserLogin(w http.ResponseWriter, r *http.Request) {

	time.Sleep(time.Second)
	r.ParseForm()
	//var user JsonBean
	//use goin get Body of Request-stream 111
	//if err:=json.Unmarshal(r.Body,&user);err!=nil{
	//	w.Header().Set("Content-Type","application/json;charset=UTF-8")
	//	w.WriteHeader(400)
	//}

	param_userName, found1 := r.Form["userName"]
	param_password, found2 := r.Form["password"]

	if !(found1 && found2) {
		fmt.Fprintf(w, "err 666")
		return
	}

	result := NewJsonBean()
	username := param_userName[0]
	password := param_password[0]
	s := "userName:" + username + ",password:" + password
	fmt.Println(s)
	if username == "laosiji" && password=="666"{
	result.Data = 100
	result.Message = "login success"
	}else{
	result.Data=101
		result.Message="err passwd or name"
	}
	byts,_:=json.Marshal(result)
	fmt.Fprintf(w,string(byts))
}

func main() {
	http.HandleFunc("/login", UserLogin)
	err := http.ListenAndServe(":33888", nil)
	if err != nil {
		log.Fatal("Listenandserve:", err)

	}

}
