package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/memory"
)

func TestGetTasks(t *testing.T) {
	taskRepo := new(memory.TaskRepo)
	sessRepo := memory.NewSessionRepo()

	sess, _ := todo.NewSession(todo.SessionID("session id"), todo.UserID("user id"))
	sessRepo.Push(context.Background(), sess)

	ts := []struct {
		userID  todo.UserID
		name    string
		dueDate time.Time
	}{}

	createUsecase := createTask{
		taskRepo: taskRepo,
		sessRepo: sessRepo,
	}
	for _, t := range ts {
		createUsecase.do(t.name, t.dueDate)
	}

	u := getTasks{
		repo: taskRepo,
	}
	tasks, err := u.Do()
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
	taskRepo := new(memory.TaskRepo)
	sessRepo := memory.NewSessionRepo()
	u := createTask{
		taskRepo: taskRepo,
		sessRepo: sessRepo,
	}

	userID, name, dueDate := todo.UserID("user id"), "task", time.Time{}

	sess, _ := todo.NewSession(todo.SessionID("session id"), userID)
	sessRepo.Push(context.Background(), sess)

	task, err := u.do(name, dueDate)
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
	taskRepo := new(memory.TaskRepo)
	sessRepo := memory.NewSessionRepo()

	createUsecase := createTask{
		taskRepo: taskRepo,
		sessRepo: sessRepo,
	}

	userID, name, dueDate := todo.UserID("user id"), "task", time.Time{}

	sess, _ := todo.NewSession(todo.SessionID("session id"), userID)
	sessRepo.Push(context.Background(), sess)

	task, _ := createUsecase.do(name, dueDate)

	u := postponeTask{
		repo: taskRepo,
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
