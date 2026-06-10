package services

import (
	"airport-vip-lounge/src/models"
	"airport-vip-lounge/src/repository"
)

type CompanionService struct {
	repo      *repository.CompanionRepository
	auditRepo *repository.AuditLogRepository
}

func NewCompanionService() *CompanionService {
	return &CompanionService{
		repo:      repository.NewCompanionRepository(),
		auditRepo: repository.NewAuditLogRepository(),
	}
}

func (s *CompanionService) Create(companion *models.Companion) error {
	companion.Status = "pending"
	companion.VerificationStatus = "pending"
	err := s.repo.Create(companion)
	if err != nil {
		return err
	}
	s.createAuditLog("CREATE", "companion", companion.ID, "Created companion")
	return nil
}

func (s *CompanionService) GetByID(id int) (*models.Companion, error) {
	return s.repo.GetByID(id)
}

func (s *CompanionService) GetAll(page, pageSize int) ([]models.Companion, int, error) {
	return s.repo.GetAll(page, pageSize)
}

func (s *CompanionService) Update(companion *models.Companion) error {
	err := s.repo.Update(companion)
	if err != nil {
		return err
	}
	s.createAuditLog("UPDATE", "companion", companion.ID, "Updated companion")
	return nil
}

func (s *CompanionService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	s.createAuditLog("DELETE", "companion", id, "Deleted companion")
	return nil
}

func (s *CompanionService) GetByAppointmentRecordID(appointmentRecordID int) ([]models.Companion, error) {
	return s.repo.GetByAppointmentRecordID(appointmentRecordID)
}

func (s *CompanionService) BatchCreate(companions []models.Companion) error {
	err := s.repo.BatchCreate(companions)
	if err != nil {
		return err
	}
	for _, companion := range companions {
		s.createAuditLog("BATCH_CREATE", "companion", companion.ID, "Batch created companion")
	}
	return nil
}

func (s *CompanionService) createAuditLog(action string, entityType string, entityID int, details string) {
	log := &models.AuditLog{
		Action:    action,
		EntityType: entityType,
		EntityID:  entityID,
		UserID:    "system",
		Details:   details,
	}
	s.auditRepo.Create(log)
}

type UsageVoucherService struct {
	repo      *repository.UsageVoucherRepository
	auditRepo *repository.AuditLogRepository
}

func NewUsageVoucherService() *UsageVoucherService {
	return &UsageVoucherService{
		repo:      repository.NewUsageVoucherRepository(),
		auditRepo: repository.NewAuditLogRepository(),
	}
}

func (s *UsageVoucherService) Create(voucher *models.UsageVoucher) error {
	voucher.Status = "active"
	err := s.repo.Create(voucher)
	if err != nil {
		return err
	}
	s.createAuditLog("CREATE", "usage_voucher", voucher.ID, "Created usage voucher")
	return nil
}

func (s *UsageVoucherService) GetByID(id int) (*models.UsageVoucher, error) {
	return s.repo.GetByID(id)
}

func (s *UsageVoucherService) GetAll(page, pageSize int) ([]models.UsageVoucher, int, error) {
	return s.repo.GetAll(page, pageSize)
}

func (s *UsageVoucherService) Update(voucher *models.UsageVoucher) error {
	err := s.repo.Update(voucher)
	if err != nil {
		return err
	}
	s.createAuditLog("UPDATE", "usage_voucher", voucher.ID, "Updated usage voucher")
	return nil
}

func (s *UsageVoucherService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	s.createAuditLog("DELETE", "usage_voucher", id, "Deleted usage voucher")
	return nil
}

func (s *UsageVoucherService) GetByStatus(status string) ([]models.UsageVoucher, error) {
	return s.repo.GetByStatus(status)
}

func (s *UsageVoucherService) createAuditLog(action string, entityType string, entityID int, details string) {
	log := &models.AuditLog{
		Action:    action,
		EntityType: entityType,
		EntityID:  entityID,
		UserID:    "system",
		Details:   details,
	}
	s.auditRepo.Create(log)
}

type WaitingListService struct {
	repo           *repository.WaitingListRepository
	domainCalcSvc  *DomainCalculationService
	auditRepo      *repository.AuditLogRepository
}

func NewWaitingListService() *WaitingListService {
	return &WaitingListService{
		repo:          repository.NewWaitingListRepository(),
		domainCalcSvc: NewDomainCalculationService(),
		auditRepo:     repository.NewAuditLogRepository(),
	}
}

func (s *WaitingListService) Create(entry *models.WaitingListEntry) error {
	entry.Status = "waiting"
	entry.WaitStartTime = time.Now()
	
	priorityScore, _ := s.domainCalcSvc.CalculateEntryPriority(entry.MemberBenefitID, time.Now().Add(2*time.Hour))
	entry.PriorityScore = priorityScore
	
	err := s.repo.Create(entry)
	if err != nil {
		return err
	}
	s.createAuditLog("CREATE", "waiting_list", entry.ID, "Added to waiting list")
	return nil
}

func (s *WaitingListService) GetByID(id int) (*models.WaitingListEntry, error) {
	return s.repo.GetByID(id)
}

func (s *WaitingListService) GetAll(page, pageSize int) ([]models.WaitingListEntry, int, error) {
	return s.repo.GetAll(page, pageSize)
}

func (s *WaitingListService) Update(entry *models.WaitingListEntry) error {
	err := s.repo.Update(entry)
	if err != nil {
		return err
	}
	s.createAuditLog("UPDATE", "waiting_list", entry.ID, "Updated waiting list entry")
	return nil
}

func (s *WaitingListService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	s.createAuditLog("DELETE", "waiting_list", id, "Deleted from waiting list")
	return nil
}

func (s *WaitingListService) ProcessAdmissions(slotID int, availableSlots int) ([]models.WaitingListEntry, error) {
	entries, err := s.domainCalcSvc.ProcessWaitingList(slotID, availableSlots)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		s.createAuditLog("ADMIT", "waiting_list", entry.ID, "Admitted from waiting list")
	}
	return entries, nil
}

func (s *WaitingListService) createAuditLog(action string, entityType string, entityID int, details string) {
	log := &models.AuditLog{
		Action:    action,
		EntityType: entityType,
		EntityID:  entityID,
		UserID:    "system",
		Details:   details,
	}
	s.auditRepo.Create(log)
}
