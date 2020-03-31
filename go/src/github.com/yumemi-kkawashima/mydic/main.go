package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/yumemi-kkawashima/mydic/controller"
	"github.com/yumemi-kkawashima/mydic/service"
	"github.com/yumemi-kkawashima/mydic/repository"
)

func main() {
	config := repository.DbConfig {
		Dbname:		"dic",
		Host:		"dbsvr",
		Port:		5432,
		User:		"root",
		Password:	"root00",
	}

	repo, err := repository.NewRepository(&config)
	if err != nil {
		panic(err)
	}
	service := service.DicService { Repo: repo }
	wc := controller.WordController { Service: service }

	router := mux.NewRouter()
	router.HandleFunc("/hello", hello)
	sub := router.PathPrefix("/word/").Subrouter()
	sub.HandleFunc("/histories", wc.GetHistories).Methods("GET")
	sub.HandleFunc("/", wc.AddWord).Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}/example", wc.GetExamples).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+}/example", wc.AddExample).Methods("POST")
	sub.HandleFunc("/{id:[0-9]+}", todo).Methods("GET")
	sub.HandleFunc("/{id:[0-9]+}", todo).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":80", nil)
}

func todo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: %s\n", r.Method)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
