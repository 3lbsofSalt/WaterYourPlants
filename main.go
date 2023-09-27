package main

import (
	"fmt"
	"log"
	"net/http"
  "html/template"
  "time"
)

var lastWatered time.Time;

func handleIndex(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("index.html"));
  tmpl.Execute(w, nil);
}

func isWatered(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    lastWatered = time.Now();
  }

  water_next := lastWatered.Add(time.Hour * 24);
  watered := !time.Now().After(water_next);

  var tmpl *template.Template;
  if watered {
    tmpl = template.Must(template.ParseFiles("watered.html"));
    data := water_next.String();
    tmpl.Execute(w, data);
  } else {
    tmpl = template.Must(template.ParseFiles("dry.html"));
    tmpl.Execute(w, nil);
  }

}


func main() {
  fmt.Println("listening...");

  lastWatered = time.Now();
  mux := http.NewServeMux();

  fs := http.FileServer(http.Dir("static"));

  mux.Handle("/static/", http.StripPrefix("/static/", fs));

  mux.HandleFunc("/", handleIndex);
  mux.HandleFunc("/is-watered", isWatered)

  log.Fatal(http.ListenAndServe(":8000", mux));
}
