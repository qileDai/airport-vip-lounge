package handlers

import (
	"airport-vip-lounge/src/dto"
	"airport-vip-lounge/src/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MemberBenefitHandler struct {
	service *services.MemberBenefitService
}

func NewMemberBenefitHandler() *MemberBenefitHandler {
	return &MemberBenefitHandler{
		service: services.NewMemberBenefitService(),
	}
}

func (h *MemberBenefitHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateMemberBenefitRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	benefit, err := h.service.Create(req)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(201).JSON(dto.APIResponse{
		Success: true,
		Data:    benefit,
		Message: "Member benefit created successfully",
	})
}

func (h *MemberBenefitHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(dto.APIResponse{
			Success: false,
			Error:   "Invalid ID",
		})
	}

	benefit, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(dto.APIResponse{
			Success: false,
			Error:   "Member benefit not found",
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data:    benefit,
	})
}

func (h *MemberBenefitHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	benefits, totalCount, err := h.service.GetAll(page, pageSize)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	totalPages := (totalCount + pageSize - 1) / pageSize
	meta := &dto.Meta{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data:    benefits,
		Meta:    meta,
	})
}

func (h *MemberBenefitHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(dto.APIResponse{
			Success: false,
			Error:   "Invalid ID",
		})
	}

	var req dto.UpdateMemberBenefitRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	benefit, err := h.service.Update(id, req)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data:    benefit,
		Message: "Member benefit updated successfully",
	})
}

func (h *MemberBenefitHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(dto.APIResponse{
			Success: false,
			Error:   "Invalid ID",
		})
	}

	err = h.service.Delete(id)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Message: "Member benefit deleted successfully",
	})
}

func (h *MemberBenefitHandler) TransitionStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(dto.APIResponse{
			Success: false,
			Error:   "Invalid ID",
		})
	}

	var req dto.StatusTransitionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(dto.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	benefit, err := h.service.TransitionStatus(id, req)
	if err != nil {
		return c.Status(400).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data:    benefit,
		Message: "Status transitioned successfully",
	})
}

func (h *MemberBenefitHandler) GetByStatus(c *fiber.Ctx) error {
	status := c.Params("status")

	benefits, err := h.service.GetByStatus(status)
	if err != nil {
		return c.Status(500).JSON(dto.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Data:    benefits,
	})
}
