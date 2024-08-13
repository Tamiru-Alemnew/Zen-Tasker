package usecases

import (
	"context"
	"errors"

	"github.com/Tamiru-Alemnew/task-manager/Domain"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
}

// NewTaskUsecase creates a new instance of taskUsecase
func NewTaskUsecase(taskRepository domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}

// GetAll retrieves all tasks from the repository
func (tc *taskUsecase) GetAll(ctx context.Context) ([]domain.Task, error) {
	tasks, err := tc.taskRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetByID retrieves a task by its ID from the repository
func (tc *taskUsecase) GetByID(ctx context.Context, id int) (*domain.Task, error) {
	task, err := tc.taskRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, errors.New("task not found")
	}
	return task, nil
}

// Create creates a new task and stores it in the repository
func (tc *taskUsecase) Create(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	err := tc.taskRepository.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// Update updates an existing task in the repository
func (tc *taskUsecase) Update(ctx context.Context, id int, task *domain.Task) (*domain.Task, error) {
	// First, find the task to ensure it exists
	existingTask, err := tc.taskRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if existingTask == nil {
		return nil, errors.New("task not found")
	}

	// Perform the update
	err = tc.taskRepository.Update(ctx, id, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tc *taskUsecase) Delete(ctx context.Context, id int) error {

	existingTask, err := tc.taskRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingTask == nil {
		return errors.New("task not found")
	}

	err = tc.taskRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
