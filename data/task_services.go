package data

import (
	"errors"
	"github.com/Tamiru-Alemnew/task-manager/models"
)

var tasks = make(map[int]models.Task)
var nextID = 1

func GetAllTasks () []models.Task {
	var taskList []models.Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return taskList
}


func GetTaskByID(id int) (models.Task , error) {
	task, exists := tasks[id]

	if !exists {
        return models.Task{}, errors.New("task not found")
    }
    return task, nil
}


func CreateTask(task models.Task) models.Task {
    task.ID = nextID
    nextID++
    tasks[task.ID] = task
    return task
}


func UpdateTask(id int, newTask models.Task) (models.Task, error) {
    _, exists := tasks[id]
    if !exists {
        return models.Task{}, errors.New("task not found")
    }
    newTask.ID = id
    tasks[id] = newTask
    return newTask, nil
}

func DeleteTask(id int) error {
	_, exists := tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}