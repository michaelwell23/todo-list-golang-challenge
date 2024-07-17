package repositories

import (
    "mytodoapp/src/models"
    "mytodoapp/src/utils"
    "errors"
)

func GetTaskByID(id uint) (*models.Task, error) {
    var task models.Task
    if err := utils.DB.First(&task, id).Error; err != nil {
        return nil, errors.New("task not found")
    }
    return &task, nil
}

func AddTask(task *models.Task) error {
    if err := utils.DB.Create(task).Error; err != nil {
        return err
    }
    return nil
}

func GetPendingTasks() ([]models.Task, error) {
    var tasks []models.Task
    if err := utils.DB.Where("completed = ?", false).Find(&tasks).Error; err != nil {
        return nil, err
    }
    return tasks, nil
}

func GetCompletedTasks() ([]models.Task, error) {
    var tasks []models.Task
    if err := utils.DB.Where("completed = ?", true).Find(&tasks).Error; err != nil {
        return nil, err
    }
    return tasks, nil
}

func UpdateTask(task *models.Task) error {
    if err := utils.DB.Save(task).Error; err != nil {
        return err
    }
    return nil
}

func DeleteTask(task *models.Task) error {
    if err := utils.DB.Delete(task).Error; err != nil {
        return err
    }
    return nil
}
