package main

import (
	"fmt"
	"reflect"
	"github.com/go-resty/resty/v2"
	"github.com/buger/jsonparser"
)

func main() {
	client := resty.New()
	
	resp, err := client.R().Get("https://pokeapi.co/api/v2/pokemon")
	results, _, _, _ := jsonparser.Get(resp.Body(), "results")
	fmt.Println(reflect.TypeOf(results))
	fmt.Println(err)
} 