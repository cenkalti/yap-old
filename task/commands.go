package task

import (
	"sort"
	"time"

	"github.com/satori/go.uuid"
)

// Add new task in pending state.
func Add(title string, dueDate, waitDate *time.Time) (id uint16, err error) {
	t := Task{
		UUID:      uuid.NewV4(),
		Title:     title,
		CreatedAt: time.Now(),
		DueDate:   dueDate,
		WaitDate:  waitDate,
	}
	if err = t.write(); err != nil {
		return
	}
	if err = t.link(dirPendingTasks); err != nil {
		return
	}
	return t.ID, nil
}

// ListPending returns all pending tasks.
func ListPending() ([]Task, error) {
	tasks, err := tasksIn(dirPendingTasks)
	if err != nil {
		return nil, err
	}
	sort.Sort(byCreatedAtDesc(tasks))
	return tasks, nil
}

// ListCompleted returns all completed tasks.
func ListCompleted() ([]Task, error) {
	tasks, err := tasksIn(dirCompletedTasks)
	if err != nil {
		return nil, err
	}
	sort.Sort(byCompletedAtDesc(tasks))
	return tasks, nil
}

// Complete pending task.
func Complete(ids []uint16) error {
	tasks := make([]Task, 0, len(ids))
	for _, id := range ids {
		t, err := readLink(dirPendingTasks, id)
		if err != nil {
			return err
		}
		tasks = append(tasks, t)
	}
	for _, t := range tasks {
		err := t.moveLink(dirPendingTasks, dirCompletedTasks)
		if err != nil {
			return err
		}
	}
	return nil
}

// Continue completed task.
func Continue(ids []uint16) error {
	tasks := make([]Task, 0, len(ids))
	for _, id := range ids {
		t, err := readLink(dirCompletedTasks, id)
		if err != nil {
			return err
		}
		tasks = append(tasks, t)
	}
	for _, t := range tasks {
		err := t.moveLink(dirCompletedTasks, dirPendingTasks)
		if err != nil {
			return err
		}
	}
	return nil
}
