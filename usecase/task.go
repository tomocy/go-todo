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

func NewCreateTask(taskRepo todo.TaskRepo, sessRepo todo.SessionRepo) *createTask {
	return &createTask{
		taskRepo: taskRepo,
		sessRepo: sessRepo,
	}
}

type createTask struct {
	taskRepo todo.TaskRepo
	sessRepo todo.SessionRepo
}

func (u *createTask) do(name string, dueDate time.Time) (*todo.Task, error) {
	ctx := context.Background()

	sess, err := u.sessRepo.Pull(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to pull session: %w", err)
	}

	id, err := u.taskRepo.NextID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate task id: %w", err)
	}
	task, err := todo.NewTask(id, sess.UserID(), name, dueDate)
	if err != nil {
		return nil, fmt.Errorf("failed to generate task: %w", err)
	}

	if err := u.taskRepo.Save(ctx, task); err != nil {
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
