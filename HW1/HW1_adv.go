package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

type Result struct {
	CurrentObservation struct {
		TempC		float64		`json:"temp_c"`
		IconURL		string		`json:"icon_url"`
	} `json:"current_observation"`
}

func main() {
	res, err := http.Get("https://gist.githubusercontent.com/PichuChen/6ce430b5474b037b9a4dcafb719f9db1/raw/f6d9485db787f3b860da41c523e6745ee8b3fd53/NSYSU.json")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", robots)
	
	var result Result
	err = json.Unmarshal([]byte(robots), &result)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%f\n%s\n", result.CurrentObservation.TempC, result.CurrentObservation.IconURL);
}
