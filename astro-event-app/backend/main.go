package main

import (
    "log"
    "net/http"
    "astro-event-app/backend/config"
    "astro-event-app/backend/routes"
)

func main() {
    config.ConnectDatabase()
    router := routes.SetupRoutes()
    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
