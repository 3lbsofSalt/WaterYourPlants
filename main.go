package main

import (
	"fmt"
	"log"
	"net/http"
  "html/template"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("index.html"));
  tmpl.Execute(w, nil);
}


func main() {
  fmt.Println("listening...");

  mux := http.NewServeMux();

  fs := http.FileServer(http.Dir("static"));
  mux.Handle("/static/", http.StripPrefix("/static/", fs));
  mux.HandleFunc("/", handleIndex);

  log.Fatal(http.ListenAndServe(":8000", mux));
}
