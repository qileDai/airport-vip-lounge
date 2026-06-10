package services

import (
	"airport-vip-lounge/src/dto"
	"airport-vip-lounge/src/models"
	"airport-vip-lounge/src/repository"
	"time"
)

type AppointmentRecordService struct {
	repo       *repository.AppointmentRecordRepository
	statusRepo *repository.StatusTransitionRecordRepository
	auditRepo  *repository.AuditLogRepository
}

func NewAppointmentRecordService() *AppointmentRecordService {
	return &AppointmentRecordService{
		repo:       repository.NewAppointmentRecordRepository(),
		statusRepo: repository.NewStatusTransitionRecordRepository(),
		auditRepo:  repository.NewAuditLogRepository(),
	}
}

func (s *AppointmentRecordService) Create(record *models.AppointmentRecord) error {
	record.Status = "draft"
	err := s.repo.Create(record)
	if err != nil {
		return err
	}
	s.createAuditLog("CREATE", "appointment_record", record.ID, "Created appointment record")
	return nil
}

func (s *AppointmentRecordService) GetByID(id int) (*models.AppointmentRecord, error) {
	return s.repo.GetByID(id)
}

func (s *AppointmentRecordService) GetAll(page, pageSize int) ([]models.AppointmentRecord, int, error) {
	return s.repo.GetAll(page, pageSize)
}

func (s *AppointmentRecordService) Update(record *models.AppointmentRecord) error {
	err := s.repo.Update(record)
	if err != nil {
		return err
	}
	s.createAuditLog("UPDATE", "appointment_record", record.ID, "Updated appointment record")
	return nil
}

func (s *AppointmentRecordService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	s.createAuditLog("DELETE", "appointment_record", id, "Deleted appointment record")
	return nil
}

func (s *AppointmentRecordService) TransitionStatus(id int, req dto.StatusTransitionRequest) (*models.AppointmentRecord, error) {
	record, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	validTransitions := map[string][]string{
		"draft":           {"pending_review", "cancelled"},
		"pending_review":   {"confirmed", "rejected", "needs_info"},
		"confirmed":        {"checked_in", "no_show", "cancelled"},
		"checked_in":       {"completed"},
		"needs_info":       {"pending_review", "rejected", "cancelled"},
		"no_show":          {"archived"},
		"completed":        {"archived"},
		"cancelled":        {"archived"},
		"rejected":         {"archived"},
		"archived":         {},
	}

	allowedNext, exists := validTransitions[record.Status]
	if !exists {
		return nil, errors.New("invalid current status")
	}

	isValid := false
	for _, status := range allowedNext {
		if status == req.ToStatus {
			isValid = true
			break
		}
	}

	if !isValid {
		return nil, errors.New("invalid status transition from " + record.Status + " to " + req.ToStatus)
	}

	if (req.ToStatus == "rejected" || req.ToStatus == "cancelled") && req.Reason == "" {
		return nil, errors.New("reason is required for this transition")
	}

	oldStatus := record.Status
	record.Status = req.ToStatus

	err = s.repo.Update(record)
	if err != nil {
		return nil, err
	}

	statusRecord := &models.StatusTransitionRecord{
		EntityType: "appointment_record",
		EntityID:   id,
		FromStatus: oldStatus,
		ToStatus:   req.ToStatus,
		Action:     req.Action,
		Reason:     req.Reason,
		Operator:   "system",
	}
	s.statusRepo.Create(statusRecord)

	s.createAuditLog("STATUS_TRANSITION", "appointment_record", id, "Transitioned from "+oldStatus+" to "+req.ToStatus)
	return record, nil
}

func (s *AppointmentRecordService) createAuditLog(action string, entityType string, entityID int, details string) {
	log := &models.AuditLog{
		Action:    action,
		EntityType: entityType,
		EntityID:  entityID,
		UserID:    "system",
		Details:   details,
	}
	s.auditRepo.Create(log)
}

type FlightTimeSlotService struct {
	repo      *repository.FlightTimeSlotRepository
	auditRepo *repository.AuditLogRepository
}

func NewFlightTimeSlotService() *FlightTimeSlotService {
	return &FlightTimeSlotService{
		repo:      repository.NewFlightTimeSlotRepository(),
		auditRepo: repository.NewAuditLogRepository(),
	}
}

func (s *FlightTimeSlotService) Create(slot *models.FlightTimeSlot) error {
	slot.Status = "active"
	err := s.repo.Create(slot)
	if err != nil {
		return err
	}
	s.createAuditLog("CREATE", "flight_time_slot", slot.ID, "Created flight time slot")
	return nil
}

func (s *FlightTimeSlotService) GetByID(id int) (*models.FlightTimeSlot, error) {
	return s.repo.GetByID(id)
}

func (s *FlightTimeSlotService) GetAll(page, pageSize int) ([]models.FlightTimeSlot, int, error) {
	return s.repo.GetAll(page, pageSize)
}

func (s *FlightTimeSlotService) Update(slot *models.FlightTimeSlot) error {
	err := s.repo.Update(slot)
	if err != nil {
		return err
	}
	s.createAuditLog("UPDATE", "flight_time_slot", slot.ID, "Updated flight time slot")
	return nil
}

func (s *FlightTimeSlotService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	s.createAuditLog("DELETE", "flight_time_slot", id, "Deleted flight time slot")
	return nil
}

func (s *FlightTimeSlotService) GetByAirportCode(airportCode string) ([]models.FlightTimeSlot, error) {
	return s.repo.GetByAirportCode(airportCode)
}

func (s *FlightTimeSlotService) createAuditLog(action string, entityType string, entityID int, details string) {
	log := &models.AuditLog{
		Action:    action,
		EntityType: entityType,
		EntityID:  entityID,
		UserID:    "system",
		Details:   details,
	}
	s.auditRepo.Create(log)
}
