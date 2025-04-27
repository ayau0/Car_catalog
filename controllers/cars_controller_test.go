// controllers/cars_controller_test.go
package controllers

import (
	"bytes"
	"car-catalog-backend/database"
	"car-catalog-backend/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreateCar(t *testing.T) {
	InitTestDB()                // Инициализация базы данных для теста
	SeedTestData()              // Заполнение базы данных тестовыми данными
	router := SetupTestRouter() // Настройка тестового роутера

	// Сначала получаем ID брендов, которые были добавлены в SeedTestData
	var tesla, mercedes, audi models.Brand
	database.DB.First(&tesla, "name = ?", "Tesla")
	database.DB.First(&mercedes, "name = ?", "Mercedes")
	database.DB.First(&audi, "name = ?", "Audi")

	// Подготовка данных для запроса
	car := models.Car{BrandID: tesla.ID, Name: "Tesla Model 3", Year: 2023, Price: 50000}
	body, err := json.Marshal(car)
	assert.NoError(t, err)

	// Отправка POST запроса для создания машины
	req, _ := http.NewRequest("POST", "/cars", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверка, что код ответа 201 (Created)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Декодируем ответ и проверяем поля
	var createdCar models.Car
	err = json.Unmarshal(w.Body.Bytes(), &createdCar)
	assert.NoError(t, err)

	// Проверка, что имя и цена в ответе совпадают с ожидаемыми
	assert.Equal(t, car.Name, createdCar.Name)
	assert.Equal(t, car.Price, createdCar.Price)

	// Дополнительная проверка на ID автомобиля
	assert.NotZero(t, createdCar.ID, "Car ID should not be zero")

	// Проверка, что ID бренда совпадает
	assert.Equal(t, car.BrandID, createdCar.BrandID, "Car brand ID should match the expected")
}

func TestGetCars(t *testing.T) {
	InitTestDB()
	SeedTestData()
	router := SetupTestRouter()

	req, _ := http.NewRequest("GET", "/cars", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var cars []models.Car
	err := json.Unmarshal(w.Body.Bytes(), &cars)
	assert.NoError(t, err)
	assert.Len(t, cars, 3) // Убедитесь, что 3 машины в базе
}

func TestGetCarByID(t *testing.T) {
	InitTestDB()
	SeedTestData()
	router := SetupTestRouter()

	// Получаем первую машину из базы данных
	var car models.Car
	err := database.DB.First(&car, "name = ?", "Tesla Model X").Error
	assert.NoError(t, err)

	// Подготовка запроса
	req, _ := http.NewRequest("GET", "/cars/"+strconv.Itoa(int(car.ID)), nil)
	rr := httptest.NewRecorder()

	// Отправляем запрос
	router.ServeHTTP(rr, req)

	// Проверяем статус код
	assert.Equal(t, http.StatusOK, rr.Code)

	// Распарсим JSON-ответ
	var fetchedCar models.Car
	err = json.Unmarshal(rr.Body.Bytes(), &fetchedCar)
	assert.NoError(t, err)

	// Проверка, что данные совпадают
	assert.Equal(t, car.ID, fetchedCar.ID)
	assert.Equal(t, car.Name, fetchedCar.Name)
	assert.Equal(t, car.Year, fetchedCar.Year)
	assert.Equal(t, car.Price, fetchedCar.Price)
	assert.Equal(t, car.BrandID, fetchedCar.BrandID)
}

func TestUpdateCar(t *testing.T) {
	// Инициализируем базу данных и заполняем её тестовыми данными
	InitTestDB()
	SeedTestData()
	router := SetupTestRouter()

	var brand models.Brand
	err := database.DB.First(&brand, "name = ?", "Tesla").Error
	assert.NoError(t, err)

	// Создаём новую машину
	car := models.Car{
		BrandID: brand.ID,
		Name:    "Tesla Model S",
		Year:    2022,
		Price:   75000,
	}
	err = database.DB.Create(&car).Error
	assert.NoError(t, err, "Ошибка при создании машины")

	// Обновляем данные машины
	car.Name = "Tesla Model S Updated"
	car.Price = 80000
	body, err := json.Marshal(car)
	assert.NoError(t, err)

	// Подготовка запроса для обновления
	req, _ := http.NewRequest("PUT", "/cars/"+strconv.Itoa(int(car.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверяем, что статус код 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Десериализуем обновлённые данные
	var updatedCar models.Car
	err = json.Unmarshal(w.Body.Bytes(), &updatedCar)
	assert.NoError(t, err)

	// Проверяем, что имя и цена обновлены корректно
	assert.Equal(t, "Tesla Model S Updated", updatedCar.Name)
	assert.Equal(t, 80000.0, updatedCar.Price)
}

func TestDeleteCar(t *testing.T) {
	InitTestDB()
	SeedTestData()
	router := SetupTestRouter()

	car := models.Car{BrandID: 1, Name: "Tesla Model S", Year: 2022, Price: 75000}
	database.DB.Create(&car)

	req, _ := http.NewRequest("DELETE", "/cars/"+strconv.Itoa(int(car.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Car deleted successfully", response["message"])
}
