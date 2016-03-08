package task

// PendingTask is a Task that is not completed yet.
type PendingTask struct {
	LinkedTask
}

// Complete the task.
func (t *PendingTask) Complete() error {
	return t.LinkedTask.move(dirPendingTasks, dirCompletedTasks)
}

func pendingTasks() ([]PendingTask, error) {
	linkedTasks, err := tasksIn(dirPendingTasks)
	if err != nil {
		return nil, err
	}
	pendingTasks := make([]PendingTask, 0, len(linkedTasks))
	for _, lt := range linkedTasks {
		t := PendingTask{
			LinkedTask: lt,
		}
		pendingTasks = append(pendingTasks, t)
	}
	return pendingTasks, nil
}

func getPendingTask(id uint32) (*PendingTask, error) {
	lt, err := getLinkedTask(dirPendingTasks, id)
	if err != nil {
		return nil, err
	}
	t := PendingTask{
		LinkedTask: *lt,
	}
	return &t, nil
}
