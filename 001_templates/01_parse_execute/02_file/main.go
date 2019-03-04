package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	defer f.Close()

	err = tpl.Execute(f, nil)
	if err != nil {
		log.Panic("Error writing file:", err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
