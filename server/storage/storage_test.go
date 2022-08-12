package storage

import (
	"context"
	"testing"
	"time"

	"github.com/keyslapperdev/task-manager-mono/server/models"
	"github.com/stretchr/testify/assert"
)

func Test_CreateTask(t *testing.T) {
	ctx := context.Background()
	db := NewDBStorer(true)

	task := models.Task{Title: "task1"}
	wantedTask := task
	wantedTask.ID = 1
	wantedTask.CreatedAt = time.Now().Local()
	wantedTask.UpdatedAt = time.Now().Local()
	wantedTask.StatusID = models.StatusOpen.ID

	gotTask := db.CreateTask(ctx, task)

	// nanosecnds make the values mismatched. formatting them to
	// seconds works better.
	gotCreatedAt := gotTask.CreatedAt.Local().Format(time.RFC3339)
	gotUpdatedAt := gotTask.CreatedAt.Local().Format(time.RFC3339)
	wantedCreatedAt := wantedTask.CreatedAt.Local().Format(time.RFC3339)
	wantedUpdatedAt := wantedTask.UpdatedAt.Local().Format(time.RFC3339)

	assert.Equal(t, wantedTask.ID, gotTask.ID)
	assert.Equal(t, wantedTask.StatusID, gotTask.StatusID)
	assert.Equal(t, wantedCreatedAt, gotCreatedAt)
	assert.Equal(t, wantedUpdatedAt, gotUpdatedAt)
}

func Test_GetTasks(t *testing.T) {
	ctx := context.Background()
	db := NewDBStorer(true)

	tasks := []models.Task{
		{Title: "task1"},
		{
			Title: "task2",
			Comments: []models.Comment{{
				Message: "comment1",
			}},
		},
	}

	db.CreateTask(ctx, tasks[0])
	db.CreateTask(ctx, tasks[1])

	gotTasks := db.GetTasks(ctx)

	assert.Equal(t, 2, len(gotTasks))
	assert.NotNil(t, gotTasks[1].Comments)
	assert.Equal(t,
		tasks[1].Comments[0].Message,
		gotTasks[1].Comments[0].Message,
	)
}

func Test_GetTaskByID(t *testing.T) {
	ctx := context.Background()
	db := NewDBStorer(true)

	task := models.Task{Title: "task1", ID: 1}

	db.CreateTask(ctx, task)

	gotTask := db.GetTaskByID(ctx, task.ID)

	assert.Equal(t, task.Title, gotTask.Title)
}

func Test_UpdateTask(t *testing.T) {
	ctx := context.Background()
	db := NewDBStorer(true)

	task := models.Task{Title: "task1", ID: 1}

	db.CreateTask(ctx, task)

	task.Comments = []models.Comment{{Message: "words"}}
	task.StatusID = models.StatusInProgress.ID

	time.Sleep(2 * time.Second)
	wantedUpdatedAt := time.Now().Local().Format(time.RFC3339)
	gotTask := db.UpdateTask(ctx, task)
	gotUpdatedAt := gotTask.UpdatedAt.Format(time.RFC3339)

	assert.NotNil(t, task.Comments)
	assert.Equal(t, task.StatusID, gotTask.StatusID)
	assert.Equal(t, task.Comments[0].Message, gotTask.Comments[0].Message)
	assert.Equal(t, wantedUpdatedAt, gotUpdatedAt)
}

func Test_CloseTask(t *testing.T) {
	ctx := context.Background()
	db := NewDBStorer(true)

	task := models.Task{Title: "task1"}
	task = db.CreateTask(ctx, task)

	gotTask := db.CloseTask(ctx, task)
	wantedClosedAt := time.Now().Local().Format(time.RFC3339)
	gotClosedAt := gotTask.ClosedAt.Time.Format(time.RFC3339)

	assert.Equal(t, wantedClosedAt, gotClosedAt)
	assert.Equal(t, models.StatusClosed.ID, gotTask.StatusID)
}

func Test_DeleteTask(t *testing.T) {
	ctx := context.Background()
	db := NewDBStorer(true)

	tasks := []models.Task{
		{Title: "task1"},
		{
			Title: "task2",
			Comments: []models.Comment{{
				Message: "comment1",
			}},
		},
	}

	t1 := db.CreateTask(ctx, tasks[0])
	t2 := db.CreateTask(ctx, tasks[1])

	db.DeleteTask(ctx, t2)

	gotTasks := db.GetTasks(ctx)

	assert.Equal(t, 1, len(gotTasks))
	assert.Equal(t, t1.Title, gotTasks[0].Title)
}
