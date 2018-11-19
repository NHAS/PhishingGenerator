// PhishingGenerator project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var page []byte
var err error

func inputFormIndexHandler(writer http.ResponseWriter, read *http.Request) {
	page, err = ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(page)
}

func generateSpec(writer http.ResponseWriter, read *http.Request) {
	read.ParseForm()
	fmt.Print(read.Form, "\n")

}

func main() {

	page, err = ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", inputFormIndexHandler)
	http.HandleFunc("/gen", generateSpec)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
