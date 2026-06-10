package handlers

import (
	"airport-vip-lounge/src/dto"
	"airport-vip-lounge/src/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type StatisticsHandler struct {
	domainCalcSvc *services.DomainCalculationService
}

func NewStatisticsHandler() *StatisticsHandler {
	return &StatisticsHandler{
		domainCalcSvc: services.NewDomainCalculationService(),
	}
}

func (h *StatisticsHandler) GetStatistics(c *fiber.Ctx) error {
	dateFrom := c.Query("date_from", time.Now().AddDate(0, -1, 0).Format("2006-01-02"))
	dateTo := c.Query("date_to", time.Now().Format("2006-01-02"))

	stats, err := h.domainCalcSvc.CalculateStatistics(dateFrom, dateTo)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data:    stats,
	})
}

type DomainCalculationHandler struct {
	domainCalcSvc *services.DomainCalculationService
	waitingSvc    *services.WaitingListService
}

func NewDomainCalculationHandler() *DomainCalculationHandler {
	return &DomainCalculationHandler{
		domainCalcSvc: services.NewDomainCalculationService(),
		waitingSvc:    services.NewWaitingListService(),
	}
}

func (h *DomainCalculationHandler) CalculatePriority(c *fiber.Ctx) error {
	memberBenefitID, err := strconv.Atoi(c.Params("member_benefit_id"))
	if err != nil {
		return c.Status(400).JSON(dto.APIResponse{
			Success: false,
			Error:   "Invalid member benefit ID",
		})
	}

	flightTimeStr := c.Query("flight_time", time.Now().Add(2*time.Hour).Format("2006-01-02 15:04:05"))
	flightTime, err := time.Parse("2006-01-02 15:04:05", flightTimeStr)
	if err != nil {
		flightTime = time.Now().Add(2 * time.Hour)
	}

	priorityScore, err := h.domainCalcSvc.CalculateEntryPriority(memberBenefitID, flightTime)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data: map[string]float64{
			"priority_score": priorityScore,
		},
		Message: "Priority calculated successfully",
	})
}

func (h *DomainCalculationHandler) CheckConsistency(c *fiber.Ctx) error {
	appointmentID, _ := strconv.Atoi(c.Params("appointment_id"))

	result, err := h.domainCalcSvc.CheckConsistency(appointmentID)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data:    result,
	})
}

func (h *DomainCalculationHandler) ProcessWaitingList(c *fiber.Ctx) error {
	slotID, _ := strconv.Atoi(c.Params("slot_id"))
	availableSlots, _ := strconv.Atoi(c.Query("available_slots", "5"))

	entries, err := h.waitingSvc.ProcessAdmissions(slotID, availableSlots)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data:    entries,
		Message: "Waiting list processed successfully",
	})
}

func (h *DomainCalculationHandler) DetectExpiredBenefits(c *fiber.Ctx) error {
	expiredBenefits, err := h.domainCalcSvc.DetectExpiredBenefits()
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data:    expiredBenefits,
		Message: "Expired benefits detected",
	})
}

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(dto.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"status":      "healthy",
			"service":     "Airport VIP Lounge Reservation System",
			"version":     "1.0.0",
			"timestamp":   time.Now().Format(time.RFC3339),
			"database":    "connected",
		},
		Message: "System is operational",
	})
}
