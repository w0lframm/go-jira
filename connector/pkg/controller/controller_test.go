package controller

import (
	"GoJira/pkg/connector"
	"GoJira/pkg/controller"
	"encoding/json"
	_ "encoding/json"
	"net/http"
	"net/http/httptest"
	_ "net/url"
	_ "strconv"
	_ "strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
)

func TestDownloadAllProjects_Success(t *testing.T) {
	connector.DownloadProjects = func() error {
		return nil
	}

	req, err := http.NewRequest("POST", "/updateAll", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	controller.DownloadAllProjects(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	assert.Equal(t, "ok", rr.Body.String())
}

func TestDownloadAllProjects_Failure(t *testing.T) {

	connector.DownloadProjects = func() error {
		return connector.ErrDownloadFailed
	}

	req, err := http.NewRequest("POST", "/updateAll", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	controller.DownloadAllProjects(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	assert.Equal(t, "bad request", rr.Body.String())
}

func TestDownloadProject_Success(t *testing.T) {
	connector.DownloadProject = func(key string) error {
		assert.Equal(t, "project_key", key)
		return nil
	}

	req, err := http.NewRequest("POST", "/updateProject?key=project_key", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	controller.DownloadProject(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	assert.Equal(t, "ok", rr.Body.String())
}

func TestDownloadProject_Failure(t *testing.T) {
	connector.DownloadProject = func(key string) error {
		return connector.ErrDownloadFailed
	}

	req, err := http.NewRequest("POST", "/updateProject?key=invalid_key", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	controller.DownloadProject(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	assert.Equal(t, "bad request", rr.Body.String())
}

func TestGetProjects(t *testing.T) {
	// Создаем фейковый HTTP запрос
	req, err := http.NewRequest("GET", "/getProjects?limit=10&page=1&search=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем фейковый HTTP ResponseWriter
	rr := httptest.NewRecorder()

	// Создаем фейковый маршрутизатор и регистрируем обработчик
	router := mux.NewRouter()
	router.HandleFunc("/getProjects", controller.GetProjects).Methods("GET")

	// Мокируем функцию GetProjects
	connector.GetProjects = func() ([]structure.Project, error) {
		return []structure.Project{
			{ID: 1, Name: "Test Project 1"},
			{ID: 2, Name: "Test Project 2"},
			{ID: 3, Name: "Test Project 3"},
		}, nil
	}

	// Выполняем запрос
	router.ServeHTTP(rr, req)

	// Проверяем статус код
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// Проверяем содержимое ответа
	var response structure.RestProjects
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	// Проверяем количество проектов
	expectedCount := 3
	if len(response.Projects) != expectedCount {
		t.Errorf("expected %d projects, but got %d", expectedCount, len(response.Projects))
	}

	// Проверяем значения проектов
	expectedProjects := []structure.Project{
		{ID: 1, Name: "Test Project 1"},
		{ID: 2, Name: "Test Project 2"},
		{ID: 3, Name: "Test Project 3"},
	}
	for i, expected := range expectedProjects {
		if response.Projects[i].ID != expected.ID || response.Projects[i].Name != expected.Name {
			t.Errorf("expected project %+v, but got %+v", expected, response.Projects[i])
		}
	}
}
