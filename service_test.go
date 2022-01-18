package restro

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartTestService() {
	testService, _ := r.New("test_service")
	router := mux.NewRouter()
	router.HandleFunc("/api/v1", getHandler).Methods(http.MethodGet)
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
