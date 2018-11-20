// PhishingGenerator project main.go
package main

import (
	"fmt"
	"html/template"
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

type Data struct {
	EmailAddr    string
	Domain       string
	NumEmployees string
	Complexity   string
	Information  string
}

var phishingExamples *template.Template

func generateSpec(writer http.ResponseWriter, read *http.Request) {

	read.ParseForm()
	//map[staffInfo:[] emailAddr:[] domain:[] employeesNum:[]]

	generation := &Data{read.FormValue("emailAddr"),
		read.FormValue("domain"),
		read.FormValue("employeesNum"),
		read.FormValue("attackComplexity"),
		read.FormValue("notes"),
	}

	err = phishingExamples.Execute(writer, generation)
	if err != nil {
		log.Print("Error encountered: ", err)
		fmt.Fprint(writer, "<h1> An error has occurred </h1>")
	}

	log.Println("emailAddr: ", generation.EmailAddr)
	log.Println("domain: ", generation.Domain)
	log.Println("employeesNum: ", generation.NumEmployees)
	log.Println("atackComplexity: ", generation.Complexity)
	log.Println("notes: ", generation.Information)
}

func main() {

	page, err = ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	phishingExamples, err = template.New("template.html").ParseFiles("template.html")
	if err != nil {
		log.Fatal("Cannot load template")
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", inputFormIndexHandler)
	http.HandleFunc("/gen", generateSpec)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
