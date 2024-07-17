package main

import (
    "log"
    "net/http"
    "mytodoapp/src/routes"
    "mytodoapp/src/utils"
    "mytodoapp/src/middlewares"

    "github.com/gorilla/mux"
)

func main() {
    utils.InitDB()
    defer utils.DB.Close()

    router := mux.NewRouter()
    routes.RegisterTaskRoutes(router)

    router.Use(middlewares.Protect)

    log.Println("Server running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
