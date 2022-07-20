package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type DataStatus struct {
	Status `json:"Status"`
}

type Status struct {
	Water int `json:"Water"`
	Wind  int `json:"Wind"`
}

type DataWeather struct {
	DataStatus
	WaterStatus string
	WindStatus  string
}

const PORT = ":4001"

func main() {
	http.HandleFunc("/", getRandomStatus)

	log.Println("server running at port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getRandomStatus(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	dataWeater := DataWeather{}

	//Writing struct type to a JSON file
	statusValue := DataStatus{}
	statusValue.Status.Water = randomNum(100)
	statusValue.Status.Wind = randomNum(100)
	content, err := json.Marshal(statusValue)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("data.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//Reading into struct type from a JSON file
	byteValue, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	var status DataStatus
	json.Unmarshal(byteValue, &status)

	dataWeater.DataStatus = status
	dataWeater.WaterStatus = func(waterStatus int) (result string) {
		if waterStatus <= 5 {
			result = "Aman"
		} else if waterStatus >= 6 && waterStatus <= 8 {
			result = "Siaga"
		} else {
			result = "Bahaya"
		}
		return
	}(status.Water)
	dataWeater.WindStatus = func(windStatus int) (result string) {
		if windStatus <= 6 {
			result = "Aman"
		} else if windStatus >= 7 && windStatus <= 15 {
			result = "Siaga"
		} else {
			result = "Bahaya"
		}
		return
	}(status.Wind)

	fmt.Printf("%+v <<<<\n", dataWeater)

	tpl, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Println("[ERROR]", r.Method, r.URL.Path, "error :", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(rw, dataWeater)
}

func randomNum(num int) (result int) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	result = r.Intn(num)
	return
}
