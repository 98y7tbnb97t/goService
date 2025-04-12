package userHandlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

// TestCreateUser проверяет корректную обработку POST /users запроса.
func TestCreateUser(t *testing.T) {
	// Создаем экземпляр Echo для тестирования
	e := echo.New()

	// Формируем тестовое тело запроса
	testUser := map[string]interface{}{
		"email":    "user@example.com",
		"password": "yourpassword",
	}
	jsonBody, err := json.Marshal(testUser)
	if err != nil {
		t.Fatalf("Ошибка маршалинга JSON: %v", err)
	}

	// Формируем HTTP запрос с нужными заголовками
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Для корректного тестирования можно замокать поведение функций из пакета services.
	// Например, можно переопределить services.CreateUser, если он позволяет такую замену.
	// Здесь мы предполагаем, что services.CreateUser успешно создает пользователя.
	// Для демонстрации теста можно либо использовать настоящий код, либо подменить зависимости.

	// Вызываем обработчик createUser
	if err := createUser(c); err != nil {
		t.Fatalf("Ошибка при вызове createUser: %v", err)
	}

	// Проверяем статус ответа
	if rec.Code != http.StatusCreated {
		t.Errorf("Ожидался статус %d, получен %d", http.StatusCreated, rec.Code)
	}

	// Предполагаем, что в ответе вернется созданный пользователь в виде JSON.
	var responseUser map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &responseUser); err != nil {
		t.Fatalf("Ошибка при разборе JSON ответа: %v", err)
	}

	if responseUser["email"] != testUser["email"] {
		t.Errorf("Ожидался email %q, получен %q", testUser["email"], responseUser["email"])
	}
}
