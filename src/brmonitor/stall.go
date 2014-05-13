package brmonitor

import(
  "fmt"
  "time"
  "net/http"
  "encoding/json"
)


type Stall struct {
  Id int `json:"id"`
  Status bool `json:"status"`
  LastOpened int64 `json:"last_closed"`
  LastClosed int64 `json:"last_opened"`
}

func (theStall Stall) GetStatus() (bool) {
  return theStall.Status
}

func (theStall *Stall) SetStatus(status bool) {
  // theStall.status = status
  oldStatus := &theStall.Status
  // if status == true -> set LastOpened to the time
  // if status !== true -> set LastClosed to the time
  // if status == theStall.status WHAT THE FUCK
  if status != *oldStatus {
    // cool, shit is awesome
    if status == true {
      theStall.LastOpened = time.Now().Unix()
    } else {
      theStall.LastClosed = time.Now().Unix()
    }
    *oldStatus = status
  }

}

func (theStall Stall) GetLastOpened() (string) {
  return time.Unix(theStall.LastOpened, 0).Format(time.RFC1123)
}
func (theStall Stall) GetLastClosed() (string) { 
  return time.Unix(theStall.LastClosed, 0).Format(time.RFC1123)
}
func (theStall *Stall) ToJSON() (string) {
  data, err := json.Marshal(theStall)

  if err != nil {
    panic(err.Error())
  }
  return string(data)
}

func (theStall Stall) HandleJSON(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, theStall.ToJSON())
}