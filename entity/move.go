package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct{
	Title string
	Year int `json:"released"`
	Color bool `json:"color:omitempty"`
	Actors [] string
}

func print(movies []Movie){
	data,err := json.MarshalIndent(movies,"","		")
	if(err != nil){
		log.Fatalf("JSON marshaling failed: %s",err)
	}
	fmt.Printf("%s\n",data)
}


func main(){
	var movies = [] Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool and Luke", Year: 1943, Color: true,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Bullitt", Year: 1945, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	}


}
