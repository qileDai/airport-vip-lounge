package services

import (
	"airport-vip-lounge/src/dto"
	"airport-vip-lounge/src/models"
	"airport-vip-lounge/src/repository"
	"errors"
	"time"
)

type MemberBenefitService struct {
	repo         *repository.MemberBenefitRepository
	statusRepo   *repository.StatusTransitionRecordRepository
	auditRepo    *repository.AuditLogRepository
}

func NewMemberBenefitService() *MemberBenefitService {
	return &MemberBenefitService{
		repo:       repository.NewMemberBenefitRepository(),
		statusRepo: repository.NewStatusTransitionRecordRepository(),
		auditRepo:  repository.NewAuditLogRepository(),
	}
}

func (s *MemberBenefitService) Create(req dto.CreateMemberBenefitRequest) (*models.MemberBenefit, error) {
	benefit := &models.MemberBenefit{
		Code:           req.Code,
		Name:           req.Name,
		Status:         "draft",
		MemberLevel:    req.MemberLevel,
		RemainingQuota: req.RemainingQuota,
		Owner:          req.Owner,
		BatchID:        req.BatchID,
		AirportCode:    req.AirportCode,
		LoungeID:       req.LoungeID,
		Notes:          req.Notes,
	}

	if req.ValidFrom != "" {
		benefit.ValidFrom, _ = time.Parse("2006-01-02 15:04:05", req.ValidFrom)
	}
	if req.ValidUntil != "" {
		benefit.ValidUntil, _ = time.Parse("2006-01-02 15:04:05", req.ValidUntil)
	}

	err := s.repo.Create(benefit)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("CREATE", "member_benefit", benefit.ID, "Created member benefit")
	return benefit, nil
}

func (s *MemberBenefitService) GetByID(id int) (*models.MemberBenefit, error) {
	return s.repo.GetByID(id)
}

func (s *MemberBenefitService) GetAll(page, pageSize int) ([]models.MemberBenefit, int, error) {
	return s.repo.GetAll(page, pageSize)
}

func (s *MemberBenefitService) Update(id int, req dto.UpdateMemberBenefitRequest) (*models.MemberBenefit, error) {
	benefit, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	benefit.Name = req.Name
	benefit.Status = req.Status
	benefit.MemberLevel = req.MemberLevel
	benefit.RemainingQuota = req.RemainingQuota
	benefit.Owner = req.Owner
	benefit.AirportCode = req.AirportCode
	benefit.LoungeID = req.LoungeID
	benefit.Notes = req.Notes

	if req.ValidFrom != "" {
		benefit.ValidFrom, _ = time.Parse("2006-01-02 15:04:05", req.ValidFrom)
	}
	if req.ValidUntil != "" {
		benefit.ValidUntil, _ = time.Parse("2006-01-02 15:04:05", req.ValidUntil)
	}

	err = s.repo.Update(benefit)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("UPDATE", "member_benefit", id, "Updated member benefit")
	return benefit, nil
}

func (s *MemberBenefitService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	s.createAuditLog("DELETE", "member_benefit", id, "Deleted member benefit")
	return nil
}

func (s *MemberBenefitService) TransitionStatus(id int, req dto.StatusTransitionRequest) (*models.MemberBenefit, error) {
	benefit, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	validTransitions := map[string][]string{
		"draft":        {"pending_review", "archived"},
		"pending_review": {"approved", "rejected", "needs_info"},
		"approved":     {"active", "expired"},
		"needs_info":   {"pending_review", "rejected"},
		"active":       {"expired", "suspended"},
		"expired":      {"archived"},
		"suspended":    {"active", "expired"},
		"rejected":     {"draft"},
		"archived":     {},
	}

	allowedNext, exists := validTransitions[benefit.Status]
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
		return nil, errors.New("invalid status transition from " + benefit.Status + " to " + req.ToStatus)
	}

	if req.ToStatus == "rejected" && req.Reason == "" {
		return nil, errors.New("rejection reason is required")
	}

	oldStatus := benefit.Status
	benefit.Status = req.ToStatus

	err = s.repo.Update(benefit)
	if err != nil {
		return nil, err
	}

	statusRecord := &models.StatusTransitionRecord{
		EntityType: "member_benefit",
		EntityID:   id,
		FromStatus: oldStatus,
		ToStatus:   req.ToStatus,
		Action:     req.Action,
		Reason:     req.Reason,
		Operator:   "system",
	}
	s.statusRepo.Create(statusRecord)

	s.createAuditLog("STATUS_TRANSITION", "member_benefit", id, "Transitioned from "+oldStatus+" to "+req.ToStatus)
	return benefit, nil
}

func (s *MemberBenefitService) GetByStatus(status string) ([]models.MemberBenefit, error) {
	return s.repo.GetByStatus(status)
}

func (s *MemberBenefitService) createAuditLog(action string, entityType string, entityID int, details string) {
	log := &models.AuditLog{
		Action:    action,
		EntityType: entityType,
		EntityID:  entityID,
		UserID:    "system",
		Details:   details,
	}
	s.auditRepo.Create(log)
}
