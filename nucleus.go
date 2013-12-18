package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "log"
)

func Join(w http.ResponseWriter, req *http.Request){
  msg := string("Hi Hi.. Welcome\n")
  log.Println(msg)
}

func main(){
  router := mux.NewRouter()
  router.HandleFunc("/join", Join)
  http.Handle("/", router)
  http.ListenAndServe(":4567", nil)
}
