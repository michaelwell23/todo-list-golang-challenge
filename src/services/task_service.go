package services

import (
    "mytodoapp/src/models"
    "mytodoapp/src/repositories"
    "errors"
)

func AddTask(description string, responsible string, email string) error {
    task := models.Task{
        Description: description,
        Responsible: responsible,
        Email:       email,
        Completed:   false,
        RevertCount: 0,
    }
    return repositories.AddTask(&task)
}

func GetPendingTasks() ([]models.Task, error) {
    return repositories.GetPendingTasks()
}

func GetCompletedTasks() ([]models.Task, error) {
    return repositories.GetCompletedTasks()
}

func CompleteTask(id uint) error {
    task, err := repositories.GetTaskByID(id)
    if err != nil {
        return err
    }
    task.Completed = true
    return repositories.UpdateTask(task)
}

func RevertTask(id uint) error {
    task, err := repositories.GetTaskByID(id)
    if err != nil {
        return err
    }
    if task.RevertCount >= 3 {
        return errors.New("revert limit reached")
    }
    task.Completed = false
    task.RevertCount++
    return repositories.UpdateTask(task)
}

func DeleteTask(id uint) error {
    task, err := repositories.GetTaskByID(id)
    if err != nil {
        return err
    }
    return repositories.DeleteTask(task)
}
