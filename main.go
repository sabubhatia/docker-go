package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var hnPid string;

func server() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", frontend)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	return mux
}

func frontend(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling frontend()")
	tpl, err := template.ParseFiles("./public/index.htm")
	if err != nil {
		log.Fatal(err)
	}

	if len(tpl.Templates()) < 1 {
		log.Fatal("No templates found")
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "text/html")
	err = tpl.Execute(w, hnPid)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if _,  ret := os.LookupEnv("HOSTNAME"); !ret {
		log.Fatal("HOSTNAME variable is not present in the env")
	}
	hnPid = fmt.Sprintf("%s_%d", os.Getenv("HOSTNAME"), os.Getpid())

	log.Println(hnPid, " listening on port 8081...")
	if err := http.ListenAndServe(":8081", server()); err != nil {
		log.Fatal(err)
	}
}
