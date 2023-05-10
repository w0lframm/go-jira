package main

import (
	"GoJira/pkg/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestDownloadAllProjects(t *testing.T) {
	// Создаем фейковый HTTP запрос
	req, err := http.NewRequest("POST", "/updateAll", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем фейковый HTTP ResponseWriter
	rr := httptest.NewRecorder()

	// Создаем фейковый маршрутизатор и регистрируем обработчик
	router := mux.NewRouter()
	router.HandleFunc("/updateAll", controller.DownloadAllProjects).Methods("POST")

	// Выполняем запрос
	router.ServeHTTP(rr, req)

	// Проверяем статус код
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}
}

func TestDownloadProject(t *testing.T) {
	// Создаем фейковый HTTP запрос
	req, err := http.NewRequest("POST", "/updateProject", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем фейковый HTTP ResponseWriter
	rr := httptest.NewRecorder()

	// Создаем фейковый маршрутизатор и регистрируем обработчик
	router := mux.NewRouter()
	router.HandleFunc("/updateProject", controller.DownloadProject).Methods("POST")

	// Выполняем запрос
	router.ServeHTTP(rr, req)

	// Проверяем статус код
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}
}

func TestGetProjects(t *testing.T) {
	// Создаем фейковый HTTP запрос
	req, err := http.NewRequest("GET", "/getProjects", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем фейковый HTTP ResponseWriter
	rr := httptest.NewRecorder()

	// Создаем фейковый маршрутизатор и регистрируем обработчик
	router := mux.NewRouter()
	router.HandleFunc("/getProjects", controller.GetProjects).Methods("GET")

	// Выполняем запрос
	router.ServeHTTP(rr, req)

	// Проверяем статус код
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}
}
