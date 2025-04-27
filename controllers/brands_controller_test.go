// controllers/brands_controller_test.go
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

func TestCreateBrand(t *testing.T) {
	InitTestDB()
	SeedTestData()
	router := SetupTestRouter()

	brand := models.Brand{Name: "Chevrolet", Country: "USA", Description: "American cars", LogoURL: "url"}
	body, err := json.Marshal(brand)
	assert.NoError(t, err)

	req, _ := http.NewRequest("POST", "/brands", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdBrand models.Brand
	err = json.Unmarshal(w.Body.Bytes(), &createdBrand)
	assert.NoError(t, err)
	assert.Equal(t, brand.Name, createdBrand.Name)
	assert.Equal(t, brand.Country, createdBrand.Country)
}

func TestGetBrands(t *testing.T) {
	InitTestDB()
	SeedTestData()
	router := SetupTestRouter()

	req, _ := http.NewRequest("GET", "/brands", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var brands []models.Brand
	err := json.Unmarshal(w.Body.Bytes(), &brands)
	assert.NoError(t, err)
	assert.Len(t, brands, 3) // Убедитесь, что три бренда в базе
}

func TestGetBrandByID(t *testing.T) {
	InitTestDB()
	SeedTestData()
	router := SetupTestRouter()

	brand := models.Brand{Name: "BMW", Country: "Germany", Description: "Luxury cars", LogoURL: "url"}
	database.DB.Create(&brand)

	req, _ := http.NewRequest("GET", "/brands/"+strconv.Itoa(int(brand.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetchedBrand models.Brand
	err := json.Unmarshal(w.Body.Bytes(), &fetchedBrand)
	assert.NoError(t, err)
	assert.Equal(t, brand.Name, fetchedBrand.Name)
}

func TestUpdateBrand(t *testing.T) {
	InitTestDB()
	SeedTestData()
	router := SetupTestRouter()

	brand := models.Brand{Name: "Chevrolet", Country: "USA", Description: "American cars", LogoURL: "url"}
	database.DB.Create(&brand)

	brand.Name = "Chevrolet Updated"
	brand.Description = "Updated description"
	body, err := json.Marshal(brand)
	assert.NoError(t, err)

	req, _ := http.NewRequest("PUT", "/brands/"+strconv.Itoa(int(brand.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedBrand models.Brand
	err = json.Unmarshal(w.Body.Bytes(), &updatedBrand)
	assert.NoError(t, err)
	assert.Equal(t, "Chevrolet Updated", updatedBrand.Name)
	assert.Equal(t, "Updated description", updatedBrand.Description)
}

func TestDeleteBrand(t *testing.T) {
	InitTestDB()
	SeedTestData()
	router := SetupTestRouter()

	brand := models.Brand{Name: "Chevrolet", Country: "USA", Description: "American cars", LogoURL: "url"}
	database.DB.Create(&brand)

	req, _ := http.NewRequest("DELETE", "/brands/"+strconv.Itoa(int(brand.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Brand deleted successfully", response["message"])
}
