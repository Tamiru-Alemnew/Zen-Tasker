package usecases

import (
	"context"
	"errors"

	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
)

type TaskUsecase struct {
	TaskRepository domain.TaskRepository
}

// NewTaskUsecase creates a new instance of taskUsecase
func NewTaskUsecase(taskRepository domain.TaskRepository) domain.TaskUsecase {
	return &TaskUsecase{
		TaskRepository: taskRepository,
	}
}

// GetAll retrieves all tasks from the repository
func (tc *TaskUsecase) GetAll(ctx context.Context) ([]domain.Task, error) {
	tasks, err := tc.TaskRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetByID retrieves a task by its ID from the repository
func (tc *TaskUsecase) GetByID(ctx context.Context, id int) (*domain.Task, error) {
	task, err := tc.TaskRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, errors.New("task not found")
	}
	return task, nil
}

// Create creates a new task and stores it in the repository
func (tc *TaskUsecase) Create(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	err := tc.TaskRepository.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// Update updates an existing task in the repository
func (tc *TaskUsecase) Update(ctx context.Context, id int, task *domain.Task) (*domain.Task, error) {
	// First, find the task to ensure it exists
	existingTask, err := tc.TaskRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingTask == nil {
		return nil, errors.New("task not found")
	}

	// Perform the update
	err = tc.TaskRepository.Update(ctx, id, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tc *TaskUsecase) Delete(ctx context.Context, id int) error {

	existingTask, err := tc.TaskRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingTask == nil {
		return errors.New("task not found")
	}

	err = tc.TaskRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
