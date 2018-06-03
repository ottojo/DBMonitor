package main

import (
	"net/http"
	"strconv"
	"time"
	"log"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

func main() {

	for {
		apiResponse := getData()
		fmt.Printf("Speed: %.3f km/h\n", apiResponse.Speed*3.6)
		for _, c := range apiResponse.Cards {
			fmt.Printf("Sim %s: Signalstärke %d, %s, ID %s\n", c.Provider, c.Signal, c.Type, c.Sim)
		}
		time.Sleep(1 * time.Second)
	}
}

func getData() DBApiResponse {
	resp, err := http.Get("https://wifi-bahn.de/schedule.jason?_=" + strconv.Itoa(int(time.Now().Unix())))
	if err != nil {
		log.Fatal(err)
	}

	var apiResponse DBApiResponse

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &apiResponse)

	return apiResponse
}
