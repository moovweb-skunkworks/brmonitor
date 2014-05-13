package main

import (
  "fmt"
  "net/http"
  "brmonitor"
)

func main() {
  datStall := &brmonitor.Stall{}
  datStall.SetStatus(true)
  datStall.SetStatus(false)

  http.HandleFunc("/"      , Potato)
  http.HandleFunc("/stalls", datStall.HandleJSON)
  // http.HandleFunc("/stalls", theBathroom.HandleJSON)
  http.HandleFunc("/queue" , Potato)

  fmt.Println("Starting Server...")
  http.ListenAndServe(":3000", nil)
}

func Potato(writer http.ResponseWriter, request *http.Request) {
  fmt.Println("Hello, world")
  datStall := &brmonitor.Stall{}
  datStall.SetStatus(true)
  datStall.SetStatus(false)
  fmt.Fprintf(writer, "Hello, World. Let's poop safely.\n %#v", datStall.GetLastOpened())
}
