package todo

import (
	"testing"
	"time"
)

func TestPostponeTask(t *testing.T) {
	task, _ := NewTask("id", "user id", "name", time.Time{})

	if err := task.Postpone(); err != nil {
		t.Errorf("should have postponed task once: %s", err)
		return
	}
	if err := task.Postpone(); err != nil {
		t.Errorf("should have postponed task twice: %s", err)
		return
	}
	if err := task.Postpone(); err != nil {
		t.Errorf("should have postponed three times: %s", err)
		return
	}

	if err := task.Postpone(); err == nil {
		t.Errorf("should not have postponed over four time")
		return
	}
}
