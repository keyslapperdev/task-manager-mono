package router

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-test/deep"
	"github.com/keyslapperdev/task-manager-mono/server/models"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func Test_AddTask(t *testing.T) {
	t.Run("success returned task", func(t *testing.T) {
		task := models.Task{
			Title: "a title",
			Comments: []models.Comment{{
				Message: "a message",
			}},
		}

		w := httptest.NewRecorder()
		router := SetupRouter(fakeDataMgr{task: task})

		taskJSON, _ := json.Marshal(&task)
		buf := bytes.NewBuffer(taskJSON)

		req, _ := http.NewRequest(http.MethodPost, "/api/task", buf)
		router.ServeHTTP(w, req)

		var gotData ReturnDatum
		json.Unmarshal(w.Body.Bytes(), &gotData)

		assert.Equal(t, http.StatusOK, w.Code, "unsuccessful call")
		errs := deep.Equal(ReturnDatum{Data: task}, gotData)
		if errs != nil {
			t.Fail()
			for _, err := range errs {
				t.Log(err)
			}
		}
	})
}

func Test_GetTasks(t *testing.T) {
	t.Run("success returned task", func(t *testing.T) {
		tasks := []models.Task{
			{
				Title: "testing GetTasks",
				Comments: []models.Comment{{
					Message: "a message",
				}},
			},
			{
				Title: "title 2",
			},
		}

		w := httptest.NewRecorder()
		router := SetupRouter(fakeDataMgr{tasks: tasks})

		req, _ := http.NewRequest(http.MethodGet, "/api/tasks", nil)
		router.ServeHTTP(w, req)

		var gotData ReturnData
		json.Unmarshal(w.Body.Bytes(), &gotData)

		assert.Equal(t, http.StatusOK, w.Code, "unsuccessful call")
		errs := deep.Equal(ReturnData{Data: tasks}, gotData)
		if errs != nil {
			t.Fail()
			for _, err := range errs {
				t.Log(err)
			}
		}
	})
}

func Test_GetTaskByID(t *testing.T) {
	task := models.Task{
		ID:    1,
		Title: "a title",
		Comments: []models.Comment{{
			Message: "a message",
		}},
	}

	w := httptest.NewRecorder()
	router := SetupRouter(fakeDataMgr{task: task})

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/task?id=%d", task.ID),
		nil,
	)
	router.ServeHTTP(w, req)

	var gotData ReturnDatum
	json.Unmarshal(w.Body.Bytes(), &gotData)

	t.Run("success/returned task", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "unsuccessful call")
		errs := deep.Equal(ReturnDatum{Data: task}, gotData)
		if errs != nil {
			t.Fail()
			for _, err := range errs {
				t.Log(err)
			}
		}
	})
}

func Test_UpdateTask(t *testing.T) {
	task := models.Task{
		Title: "a title",
		Comments: []models.Comment{{
			Message: "a message",
		}},
	}

	w := httptest.NewRecorder()
	router := SetupRouter(fakeDataMgr{task: task})

	taskJSON, _ := json.Marshal(&task)
	buf := bytes.NewBuffer(taskJSON)

	req, _ := http.NewRequest(http.MethodPatch, "/api/task", buf)
	router.ServeHTTP(w, req)

	var gotData ReturnDatum
	json.Unmarshal(w.Body.Bytes(), &gotData)

	t.Run("success returned task", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "unsuccessful call")
		errs := deep.Equal(ReturnDatum{Data: task}, gotData)
		if errs != nil {
			t.Fail()
			for _, err := range errs {
				t.Log(err)
			}
		}
	})
}

func Test_CloseTask(t *testing.T) {
	task := models.Task{
		Title: "a title",
		Comments: []models.Comment{{
			Message: "a message",
		}},
	}

	taskJSON, _ := json.Marshal(&task)
	buf := bytes.NewBuffer(taskJSON)

	called := false
	fdm := fakeDataMgr{task: task, called: &called}
	router := SetupRouter(fdm)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/task", buf)
	router.ServeHTTP(w, req)

	var gotData ReturnDatum
	json.Unmarshal(w.Body.Bytes(), &gotData)

	t.Run("success returned task", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "unsuccessful call")
		assert.True(t, called, "correct function not called")
		errs := deep.Equal(ReturnDatum{Data: task}, gotData)
		if errs != nil {
			t.Fail()
			for _, err := range errs {
				t.Log(err)
			}
		}
	})
}

func Test_DeleteTask(t *testing.T) {
	task := models.Task{
		Title: "a title",
		Comments: []models.Comment{{
			Message: "a message",
		}},
	}

	taskJSON, _ := json.Marshal(&task)
	buf := bytes.NewBuffer(taskJSON)

	called := false
	fdm := fakeDataMgr{task: task, called: &called}
	router := SetupRouter(fdm)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/task", buf)
	router.ServeHTTP(w, req)

	var gotData ReturnDatum
	json.Unmarshal(w.Body.Bytes(), &gotData)

	t.Run("success returned task", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, w.Code, "unsuccessful call")
		errs := deep.Equal(ReturnDatum{Data: task}, gotData)
		if errs != nil {
			t.Fail()
			for _, err := range errs {
				t.Log(err)
			}
		}
	})
}

type ReturnDatum struct {
	Data models.Task `json:"data"`
}

type ReturnData struct {
	Data []models.Task `json:"data"`
}

type fakeDataMgr struct {
	tasks  []models.Task
	task   models.Task
	called *bool
}

func (fdm fakeDataMgr) CreateTask(c context.Context, task models.Task) models.Task { return fdm.task }
func (fdm fakeDataMgr) GetTasks(c context.Context) []models.Task                   { return fdm.tasks }
func (fdm fakeDataMgr) GetTaskByID(c context.Context, id uint) models.Task         { return fdm.task }
func (fdm fakeDataMgr) UpdateTask(c context.Context, task models.Task) models.Task { return fdm.task }
func (fdm fakeDataMgr) CloseTask(c context.Context, task models.Task) models.Task {
	*fdm.called = true
	return fdm.task
}
func (fdm fakeDataMgr) DeleteTask(c context.Context, task models.Task) {
	*fdm.called = true
}
