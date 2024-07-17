package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "mytodoapp/src/services"
    "github.com/gorilla/mux"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Description string `json:"description"`
        Responsible string `json:"responsible"`
        Email       string `json:"email"`
    }
    json.NewDecoder(r.Body).Decode(&request)
    err := services.AddTask(request.Description, request.Responsible, request.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func GetPendingTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := services.GetPendingTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(tasks)
}

func GetCompletedTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := services.GetCompletedTasks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(tasks)
}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }
    err = services.CompleteTask(uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func RevertTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }
    err = services.RevertTask(uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }
    err = services.DeleteTask(uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
