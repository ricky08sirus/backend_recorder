package main

import (
    "log"
    "net/http"
    "go-client/api"
)

func main() {
    router := api.Router()

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

