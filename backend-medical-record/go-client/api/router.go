package api

import (
    "github.com/gorilla/mux"
)

func Router() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/add-record", AddRecordHandler).Methods("POST")

    return router
}

