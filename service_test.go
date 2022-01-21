package restro

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartTestService() {
	testService, _ := r.New("test_service")
	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	router.HandleFunc("", getHandler).Methods(http.MethodGet)
	router.HandleFunc("", postHandler).Methods(http.MethodPost)
	router.HandleFunc("", putHandler).Methods(http.MethodPut)
	router.HandleFunc("", deleteHandler).Methods(http.MethodDelete)
	srv := &http.Server{
		Handler: router,
		Addr:    testService.URL.Host,
	}
	
	log.Fatal(srv.ListenAndServe())
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"running"}`)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var (
		body *requestModel
		err  error
	)
	// vars := mux.Vars(r)
	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"your number":%d}`, body.Number)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	var (
		body *requestModel
		err  error
	)
	// vars := mux.Vars(r)
	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error":"%s"}`, err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"your number":%d}`, body.Number)
}

type requestModel struct {
	// title  string
	Number int `json:"number"`
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"running"}`)
}
