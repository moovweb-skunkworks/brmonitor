package brmonitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Bathroom struct {
	Stalls []Stall `json:"stalls"`
}

type FromJson struct {
	Status bool `json:"status"`
}

// {
//   "status": "false"
// }

func (bathroom Bathroom) GetStalls() string {
	data, err := json.Marshal(bathroom)

	if err != nil {
		panic(err.Error())
	}

	return string(data)
}

func (bathroom *Bathroom) AddStall() {
	// write to db the stall.
	newStall := &Stall{}
	newStall.Status = true

	newStall.Id = len(bathroom.Stalls) + 1

	fmt.Println(newStall.ToJSON())

	bathroom.Stalls = append(bathroom.Stalls, *newStall)
}

func (bathroom Bathroom) HandleStalls(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		writer.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(writer, "%s", bathroom.GetStalls())
	} else if request.Method == "PUT" {
		currPath := request.URL.Path

		pathEls := strings.Split(currPath, "/")
		if len(pathEls) == 3 {
			pathId, err := strconv.Atoi(pathEls[len(pathEls)-1])
			if err != nil {
				thatsShitty(err.Error(), writer, request)
			} else {
				// update dat thingieeeeeeeeeeeeeeeeeee
				// fmt.Fprintf(writer, "%d", pathId)
				stall, err := bathroom.GetStall(pathId)
				if err != nil {
					thatsShitty(err.Error(), writer, request)
				} else {
					body, err := ioutil.ReadAll(request.Body)
					if err != nil {
						thatsShitty(err.Error(), writer, request)
					} else {
						// ody = string(body)
						put := &FromJson{}
						err := json.Unmarshal(body, put)

						if err != nil {
							thatsShitty(err.Error(), writer, request)

						} else {
							stall.SetStatus(put.Status)
							writer.Header().Set("Content-Type", "application/json")
							fmt.Fprintf(writer, "%s", stall.ToJSON())
						}
					}
				}
			}
		} else {
			thatsShitty("URL Format Incorrect.", writer, request)
		}
		// fmt.Fprintf(writer, "%s", currPath)
	}
}

func (bathroom Bathroom) GetStall(stallId int) (*Stall, error) {
	if len(bathroom.Stalls) > stallId-1 && stallId > 0 {
		return &bathroom.Stalls[stallId-1], nil
	} else {
		return nil, fmt.Errorf("No stall with that ID %d", stallId)
	}
}

func thatsShitty(msg string, writer http.ResponseWriter, request *http.Request) {
	http.Error(writer, "Page Not Found "+msg, 404)
}

// POST /stalls/2/update
// bathroom.GetStall(2)
// update it
// reply with new stall.
// {
// stall: {
// status: false
// }
// }

// stall[status]=false

// 1: Stall w/ id = 1
// 2:

// Bathroom.FindById(4)
