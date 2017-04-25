package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	var user JsonBean
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(400)
	}

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
	if username == "laosiji" && password == "666" {
		result.Data = 100
		result.Message = "login success"
	} else {
		result.Data = 101
		result.Message = "err passwd or name"
	}
	byts, _ := json.Marshal(result)
	fmt.Fprintf(w, string(byts))
}

func main() {
	server := &http.Server{
		Addr:         ":5555",
		WriteTimeout: 2 * time.Second,
		//Handler:      http.HandlerFunc(myHandler{}),
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)
	mux.HandleFunc("/login", UserLogin)

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	server.Handler = mux
	log.Print("start server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("server closed under request")
		} else {
			log.Fatal("server closed unexpected")
		}

	}

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,its my mux!"))
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)
	w.Write([]byte("see u my friend!"))
}
