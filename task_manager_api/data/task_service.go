package data 

import (
	"errors"
	"task_manager_api/models"
)

var Tasks []models.Task

func AddTask(newTask models.Task) error {
	if newTask.ID == "" {
		return errors.New("task ID cannot be empty")
	}
	Tasks = append(Tasks, newTask)
	return nil
}

func ViewTasks() []models.Task {
	return Tasks
}

func ViewTaskByID(id string) (models.Task, error) {
	for _, task := range Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func UpdateTask(id string, updatedTask models.Task) error {
	for i, task := range Tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				Tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				Tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				Tasks[i].Status = updatedTask.Status
			}
			if !updatedTask.DueDate.IsZero() {
				Tasks[i].DueDate = updatedTask.DueDate
			}
			return nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id string) error {
	for i, task := range Tasks {
		if task.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
