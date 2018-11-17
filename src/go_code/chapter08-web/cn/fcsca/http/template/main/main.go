package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tmpl, err := template.ParseFiles("/main.html")
		if err != nil {
			fmt.Fprintf(writer, "ParseFiles:%v", err)
			return
		}

		/*
			os.Stdout-----------响应输出到控制台
			writer------------响应输出到页面
		*/
		err = tmpl.Execute(writer, map[string]interface{}{
			"Request": request,
			"Score":   99,
		})
		if err != nil {
			fmt.Fprintf(writer, "Execute:%v", err)
			return
		}
	})
	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
