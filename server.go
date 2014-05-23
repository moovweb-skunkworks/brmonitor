package main

import (
  "fmt"
  "net/http"
  "brmonitor"
  "io/ioutil"
  "path/filepath"
  "strings"
)

func main() {
  datBathroom := &brmonitor.Bathroom{}
  datBathroom.Stalls = make([]brmonitor.Stall, 0, 2)
  datBathroom.AddStall()
  datBathroom.AddStall()

  sc := &ServerContext{}

  // fmt.Println("Hello, world")
  http.HandleFunc("/",        sc.handleIndex)
  http.HandleFunc("/stalls/", datBathroom.HandleStalls)
  http.HandleFunc("/stalls",  datBathroom.HandleStalls)

  fmt.Println("Starting Server...")
  http.ListenAndServe(":80", nil)
}

type ServerContext struct {
  HTMLCache string
}

func (sc *ServerContext) handleIndex(writer http.ResponseWriter, request *http.Request) {
  // write html
  if len(request.URL.Path) > 1 {
    // you're looking for a file.
    url := strings.Split(request.URL.Path, "/")
    fileName := url[1]
    if strings.HasSuffix(fileName, ".js") {
      writer.Header().Set("Content-Type" , "application/javascript")
    } else if strings.HasSuffix(fileName, ".css") {
      writer.Header().Set("Content-Type" , "text/css")
    }
    resp, err := ioutil.ReadFile(filepath.Join("assets", fileName))
    if err != nil {
      panic(err.Error())
    } 
    fmt.Fprintf(writer, "%s", string(resp))

  } else  {
    if len(sc.HTMLCache) <= 0 {
      fileName := "application.html"
      htmlResp, err := ioutil.ReadFile(filepath.Join("assets", fileName))
      if err != nil {
        panic(err.Error())
      } 
      sc.HTMLCache = string(htmlResp)      
    }
    fmt.Fprintf(writer, "%s", sc.HTMLCache)
  } 
}

// func sendHTML()
