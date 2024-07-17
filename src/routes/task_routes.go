package routes

import (
    "github.com/gorilla/mux"
    "mytodoapp/src/controllers"
)

func RegisterTaskRoutes(router *mux.Router) {
    router.HandleFunc("/tasks", controllers.AddTask).Methods("POST")
    router.HandleFunc("/tasks/pending", controllers.GetPendingTasks).Methods("GET")
    router.HandleFunc("/tasks/completed", controllers.GetCompletedTasks).Methods("GET")
    router.HandleFunc("/tasks/{id}/complete", controllers.CompleteTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}/revert", controllers.RevertTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
}
