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

func handleIsWatered(w http.ResponseWriter, r *http.Request) {
  watered := false;
  // Do I need to initialize tmpl here?
  var tmpl *template.Template
  if watered {
    tmpl = template.Must(template.ParseFiles("watered.html"));
  } else {
    tmpl = template.Must(template.ParseFiles("dry.html"));
  }

  tmpl.Execute(w, nil);
}


func main() {
  fmt.Println("listening...");

  mux := http.NewServeMux();

  fs := http.FileServer(http.Dir("static"));
  mux.Handle("/static/", http.StripPrefix("/static/", fs));
  mux.HandleFunc("/", handleIndex);
  mux.HandleFunc("/is-watered", handleIsWatered)

  log.Fatal(http.ListenAndServe(":8000", mux));
}
