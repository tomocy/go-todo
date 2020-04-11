package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/tomocy/go-todo"
)

func NewGetTasks(repo todo.TaskRepo) *getTasks {
	return &getTasks{
		repo: repo,
	}
}

type getTasks struct {
	repo todo.TaskRepo
}

func (u *getTasks) Do() ([]*todo.Task, error) {
	ctx := context.TODO()

	tasks, err := u.repo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	return tasks, nil
}

func NewCreateTask(repo todo.TaskRepo) *createTask {
	return &createTask{
		repo: repo,
	}
}

type createTask struct {
	repo todo.TaskRepo
}

func (u *createTask) do(userID todo.UserID, name string, dueDate time.Time) (*todo.Task, error) {
	ctx := context.Background()

	id, err := u.repo.NextID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate task id: %w", err)
	}
	task, err := todo.NewTask(id, userID, name, dueDate)
	if err != nil {
		return nil, fmt.Errorf("failed to generate task: %w", err)
	}

	if err := u.repo.Save(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to save task: %w", err)
	}

	return task, nil
}

type postponeTask struct {
	repo todo.TaskRepo
}

func (u *postponeTask) do(id todo.TaskID) (*todo.Task, error) {
	ctx := context.TODO()

	task, err := u.repo.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find task: %w", err)
	}

	if err := task.Postpone(); err != nil {
		return nil, fmt.Errorf("failed to postpone task: %w", err)
	}

	if err := u.repo.Save(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to postpone task: %w", err)
	}

	return task, nil
}
