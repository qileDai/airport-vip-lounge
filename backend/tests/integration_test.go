package tests

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lunch/src/handlers"
	"airport-vip-lunch/src/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func setupTestApp() *fiber.App {
	app := fiber.New(fiber.Config{AppName: "Test App"})
	database.InitDB()
	database.RunMigrations()
	return app
}

func TestHealthCheck(t *testing.T) {
	app := setupTestApp()
	app.Get("/health", handlers.HealthCheck)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	if response["success"] != true {
		t.Errorf("Expected success=true, got %v", response["success"])
	}
}

func TestCreateMemberBenefit(t *testing.T) {
	app := setupTestApp()
	handler := handlers.NewMemberBenefitHandler()

	api := app.Group("/api/v1")
	api.Post("/member-benefits", handler.Create)

	payload := map[string]interface{}{
		"code":           "TEST-MB-001",
		"name":           "测试会员权益",
		"member_level":   "gold",
		"remaining_quota": 5,
		"owner":          "测试员",
		"airport_code":   "PEK",
	}

	bodyBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/member-benefits", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	responseBody, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(responseBody, &response)

	if response["success"] != true {
		t.Errorf("Expected success=true, got %v", response["success"])
	}
}

func TestGetAllMemberBenefits(t *testing.T) {
	app := setupTestApp()
	handler := handlers.NewMemberBenefitHandler()

	api := app.Group("/api/v1")
	api.Get("/member-benefits", handler.GetAll)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/member-benefits?page=1&page_size=10", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestStatusTransition(t *testing.T) {
	app := setupTestApp()
	handler := handlers.NewMemberBenefitHandler()

	api := app.Group("/api/v1")
	api.Post("/member-benefits", handler.Create)
	api.Post("/member-benefits/:id/transition", handler.TransitionStatus)

	createPayload := map[string]interface{}{
		"code":           "TEST-MB-002",
		"name":           "状态流转测试权益",
		"member_level":   "platinum",
		"remaining_quota": 10,
		"owner":          "测试员",
		"airport_code":   "PVG",
	}

	bodyBytes, _ := json.Marshal(createPayload)
	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/member-benefits", bytes.NewBuffer(bodyBytes))
	createReq.Header.Set("Content-Type", "application/json")

	createResp, _ := app.Test(createReq, -1)
	createBody, _ := io.ReadAll(createResp.Body)
	var createResponse map[string]interface{}
	json.Unmarshal(createBody, &createResponse)

	data := createResponse["data"].(map[string]interface{})
	id := int(data["id"].(float64))

	transitionPayload := map[string]interface{}{
		"to_status": "pending_review",
		"action":    "submit_for_review",
		"reason":    "",
	}

	transitionBytes, _ := json.Marshal(transitionPayload)
	transitionReq := httptest.NewRequest(http.MethodPost, "/api/v1/member-benefits/"+string(rune(id))+"/transition", bytes.NewBuffer(transitionBytes))
	transitionReq.Header.Set("Content-Type", "application/json")

	transitionResp, err := app.Test(transitionReq, -1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if transitionResp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", transitionResp.StatusCode)
	}
}

func TestInvalidStatusTransition(t *testing.T) {
	app := setupTestApp()
	handler := handlers.NewMemberBenefitHandler()

	api := app.Group("/api/v1")
	api.Post("/member-benefits/:id/transition", handler.TransitionStatus)

	transitionPayload := map[string]interface{}{
		"to_status": "invalid_status",
		"action":    "test",
		"reason":    "",
	}

	transitionBytes, _ := json.Marshal(transitionPayload)
	transitionReq := httptest.NewRequest(http.MethodPost, "/api/v1/member-benefits/99999/transition", bytes.NewBuffer(transitionBytes))
	transitionReq.Header.Set("Content-Type", "application/json")

	transitionResp, _ := app.Test(transitionReq, -1)

	if transitionResp.StatusCode == http.StatusOK {
		t.Error("Expected error status for invalid transition")
	}
}

func TestSeedDataReset(t *testing.T) {
	app := setupTestApp()
	api := app.Group("/api/v1")
	api.Post("/seed/reset", handlers.ResetSeedData)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/seed/reset", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	json.Unmarshal(body, &response)

	if response["success"] != true {
		t.Errorf("Expected success=true, got %v", response["success"])
	}
}

func TestDomainCalculationPriority(t *testing.T) {
	calcService := services.NewDomainCalculationService()

	testTime, _ := time.Parse("2006-01-02 15:04:05", "2026-06-01 10:00:00")
	score, err := calcService.CalculateEntryPriority(1, testTime)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if score <= 0 || score > 100 {
		t.Errorf("Expected score between 0 and 100, got %f", score)
	}
}

func TestStatisticsEndpoint(t *testing.T) {
	app := setupTestApp()
	handler := handlers.NewStatisticsHandler()

	api := app.Group("/api/v1")
	api.Get("/statistics", handler.GetStatistics)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/statistics?date_from=2026-01-01&date_to=2026-12-31", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
