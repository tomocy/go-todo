package usecase

import (
	"testing"
	"time"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/memory"
)

func TestGetTasks(t *testing.T) {
	repo := new(memory.TaskRepo)

	ts := []struct {
		userID  todo.UserID
		name    string
		dueDate time.Time
	}{}

	createUsecase := createTask{
		repo: repo,
	}
	for _, t := range ts {
		createUsecase.do(t.userID, t.name, t.dueDate)
	}

	u := getTasks{
		repo: repo,
	}
	tasks, err := u.do()
	if err != nil {
		t.Errorf("should have got tasks: %s", err)
		return
	}

	if len(tasks) != len(ts) {
		t.Errorf("should have returned the got tasks: %s", reportUnexpected("len of tasks", len(tasks), len(ts)))
		return
	}
}

func TestCreateTask(t *testing.T) {
	repo := new(memory.TaskRepo)
	u := createTask{
		repo: repo,
	}

	userID, name, dueDate := todo.UserID("user id"), "task", time.Time{}

	task, err := u.do(userID, name, dueDate)
	if err != nil {
		t.Errorf("should have created task: %s", err)
		return
	}

	if err := assertTask(task, userID, name, dueDate); err != nil {
		t.Errorf("should have returned the created task: %s", err)
		return
	}
}

func TestPostponeTask(t *testing.T) {
	repo := new(memory.TaskRepo)

	createUsecase := createTask{
		repo: repo,
	}
	task, _ := createUsecase.do("user id", "name", time.Time{})

	u := postponeTask{
		repo: repo,
	}
	task, err := u.do(task.ID())
	if err != nil {
		t.Errorf("should have postponed task: %s", err)
		return
	}

	if task.PostponedTimes() != 1 {
		t.Errorf("should have returned the postponed task: %s", err)
		return
	}
}

func assertTask(t *todo.Task, userID todo.UserID, name string, dueDate time.Time) error {
	if t.UserID() != userID {
		return reportUnexpected("user id", t.UserID(), userID)
	}
	if t.Name() != name {
		return reportUnexpected("name", t.Name(), name)
	}
	if t.DueDate() != dueDate {
		return reportUnexpected("due date", t.DueDate(), dueDate)
	}

	return nil
}
