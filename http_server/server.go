package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "pong")
	})

	handler.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		vars := request.URL.Query()
		fmt.Println(vars)
		fmt.Fprint(writer, "ok")
	})

	handler.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		js, _ := json.Marshal(SUCCESS)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(js)
	})

	handler.HandleFunc("/xml", func(writer http.ResponseWriter, request *http.Request) {
		js, _ := xml.MarshalIndent(SUCCESS, "", "  ")
		writer.Header().Set("Content-Type", "application/xml")
		writer.Write(js)
	})

	_ = http.ListenAndServe(":20000", handler)
}
