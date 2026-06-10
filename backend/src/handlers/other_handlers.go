package handlers

import (
	"airport-vip-lounge/src/dto"
	"airport-vip-lounge/src/models"
	"airport-vip-lounge/src/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AppointmentRecordHandler struct {
	service *services.AppointmentRecordService
}

func NewAppointmentRecordHandler() *AppointmentRecordHandler {
	return &AppointmentRecordHandler{
		service: services.NewAppointmentRecordService(),
	}
}

func (h *AppointmentRecordHandler) Create(c *fiber.Ctx) error {
	var record models.AppointmentRecord
	if err := c.BodyParser(&record); err != nil {
		return c.Status(400).JSON(dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	err := h.service.Create(&record)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(201).JSON(dto.APIResponse{
		Success: true,
		Data:    record,
		Message: "Appointment record created successfully",
	})
}

func (h *AppointmentRecordHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	record, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(dto.APIResponse{Success: false, Error: "Not found"})
	}
	return c.JSON(dto.APIResponse{Success: true, Data: record})
}

func (h *AppointmentRecordHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	records, totalCount, err := h.service.GetAll(page, pageSize)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}

	totalPages := (totalCount + pageSize - 1) / pageSize
	meta := &dto.Meta{Page: page, PageSize: pageSize, TotalCount: totalCount, TotalPages: totalPages}
	return c.JSON(dto.APIResponse{Success: true, Data: records, Meta: meta})
}

func (h *AppointmentRecordHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var record models.AppointmentRecord
	if err := c.BodyParser(&record); err != nil {
		return c.Status(400).JSON(dto.APIResponse{Success: false, Error: "Invalid request body"})
	}
	record.ID = id

	err := h.service.Update(&record)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}
	return c.JSON(dto.APIResponse{Success: true, Data: record, Message: "Updated successfully"})
}

func (h *AppointmentRecordHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := h.service.Delete(id)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}
	return c.JSON(dto.APIResponse{Success: true, Message: "Deleted successfully"})
}

func (h *AppointmentRecordHandler) TransitionStatus(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req dto.StatusTransitionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(dto.APIResponse{Success: false, Error: "Invalid request body"})
	}

	record, err := h.service.TransitionStatus(id, req)
	if err != nil {
		return c.Status(400).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}
	return c.JSON(dto.APIResponse{Success: true, Data: record, Message: "Status transitioned"})
}

type FlightTimeSlotHandler struct {
	service *services.FlightTimeSlotService
}

func NewFlightTimeSlotHandler() *FlightTimeSlotHandler {
	return &FlightTimeSlotHandler{
		service: services.NewFlightTimeSlotService(),
	}
}

func (h *FlightTimeSlotHandler) Create(c *fiber.Ctx) error {
	var slot models.FlightTimeSlot
	if err := c.BodyParser(&slot); err != nil {
		return c.Status(400).JSON(dto.APIResponse{Success: false, Error: "Invalid request body"})
	}

	err := h.service.Create(&slot)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}
	return c.Status(201).JSON(dto.APIResponse{Success: true, Data: slot, Message: "Created successfully"})
}

func (h *FlightTimeSlotHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	slot, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(dto.APIResponse{Success: false, Error: "Not found"})
	}
	return c.JSON(dto.APIResponse{Success: true, Data: slot})
}

func (h *FlightTimeSlotHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	slots, totalCount, err := h.service.GetAll(page, pageSize)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}

	totalPages := (totalCount + pageSize - 1) / pageSize
	meta := &dto.Meta{Page: page, PageSize: pageSize, TotalCount: totalCount, TotalPages: totalPages}
	return c.JSON(dto.APIResponse{Success: true, Data: slots, Meta: meta})
}

func (h *FlightTimeSlotHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var slot models.FlightTimeSlot
	if err := c.BodyParser(&slot); err != nil {
		return c.Status(400).JSON(dto.APIResponse{Success: false, Error: "Invalid request body"})
	}
	slot.ID = id

	err := h.service.Update(&slot)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}
	return c.JSON(dto.APIResponse{Success: true, Data: slot, Message: "Updated successfully"})
}

func (h *FlightTimeSlotHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := h.service.Delete(id)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}
	return c.JSON(dto.APIResponse{Success: true, Message: "Deleted successfully"})
}

func (h *FlightTimeSlotHandler) GetByAirportCode(c *fiber.Ctx) error {
	airportCode := c.Params("airport_code")
	slots, err := h.service.GetByAirportCode(airportCode)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}
	return c.JSON(dto.APIResponse{Success: true, Data: slots})
}

type CompanionHandler struct {
	service *services.CompanionService
}

func NewCompanionHandler() *CompanionHandler {
	return &CompanionHandler{
		service: services.NewCompanionService(),
	}
}

func (h *CompanionHandler) Create(c *fiber.Ctx) error {
	var companion models.Companion
	if err := c.BodyParser(&companion); err != nil {
		return c.Status(400).JSON(dto.APIResponse{Success: false, Error: "Invalid request body"})
	}

	err := h.service.Create(&companion)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}
	return c.Status(201).JSON(dto.APIResponse{Success: true, Data: companion, Message: "Created successfully"})
}

func (h *CompanionHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	companion, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(dto.APIResponse{Success: false, Error: "Not found"})
	}
	return c.JSON(dto.APIResponse{Success: true, Data: companion})
}

func (h *CompanionHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	companions, totalCount, err := h.service.GetAll(page, pageSize)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}

	totalPages := (totalCount + pageSize - 1) / pageSize
	meta := &dto.Meta{Page: page, PageSize: pageSize, TotalCount: totalCount, TotalPages: totalPages}
	return c.JSON(dto.APIResponse{Success: true, Data: companions, Meta: meta})
}

func (h *CompanionHandler) BatchCreate(c *fiber.Ctx) error {
	var req struct {
		Records []models.Companion `json:"records"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(dto.APIResponse{Success: false, Error: "Invalid request body"})
	}

	err := h.service.BatchCreate(req.Records)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}
	return c.Status(201).JSON(dto.APIResponse{Success: true, Message: "Batch created successfully"})
}

func (h *CompanionHandler) GetByAppointmentID(c *fiber.Ctx) error {
	appointmentID, _ := strconv.Atoi(c.Params("appointment_id"))
	companions, err := h.service.GetByAppointmentRecordID(appointmentID)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{Success: false, Error: err.Error()})
	}
	return c.JSON(dto.APIResponse{Success: true, Data: companions})
}
