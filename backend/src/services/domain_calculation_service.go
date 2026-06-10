package services

import (
	"airport-vip-lounge/src/models"
	"airport-vip-lounge/src/repository"
	"math"
	"time"
)

type DomainCalculationService struct {
	memberRepo       *repository.MemberBenefitRepository
	appointmentRepo  *repository.AppointmentRecordRepository
	flightSlotRepo   *repository.FlightTimeSlotRepository
	waitingListRepo  *repository.WaitingListRepository
	verificationRepo *repository.VerificationResultRepository
}

func NewDomainCalculationService() *DomainCalculationService {
	return &DomainCalculationService{
		memberRepo:       repository.NewMemberBenefitRepository(),
		appointmentRepo:  repository.NewAppointmentRecordRepository(),
		flightSlotRepo:   repository.NewFlightTimeSlotRepository(),
		waitingListRepo:  repository.NewWaitingListRepository(),
		verificationRepo: repository.NewVerificationResultRepository(),
	}
}

func (s *DomainCalculationService) CalculateEntryPriority(memberBenefitID int, flightTime time.Time) (float64, error) {
	benefit, err := s.memberRepo.GetByID(memberBenefitID)
	if err != nil {
		return 0, err
	}

	levelScore := s.calculateMemberLevelScore(benefit.MemberLevel)
	quotaScore := s.calculateQuotaScore(benefit.RemainingQuota)
	timeScore := s.calculateTimeScore(flightTime)

	priorityScore := levelScore*0.4 + quotaScore*0.3 + timeScore*0.3
	return math.Round(priorityScore*100) / 100, nil
}

func (s *DomainCalculationService) calculateMemberLevelScore(level string) float64 {
	levelScores := map[string]float64{
		"platinum": 100,
		"diamond":  90,
		"gold":     75,
		"silver":   60,
		"bronze":   40,
		"basic":    20,
	}

	if score, exists := levelScores[level]; exists {
		return score
	}
	return 0
}

func (s *DomainCalculationService) calculateQuotaScore(remainingQuota int) float64 {
	if remainingQuota <= 0 {
		return 0
	} else if remainingQuota >= 10 {
		return 100
	} else {
		return float64(remainingQuota) * 10
	}
}

func (s *DomainCalculationService) calculateTimeScore(flightTime time.Time) float64 {
	now := time.Now()
	duration := flightTime.Sub(now).Hours()

	if duration < 0 {
		return 0
	} else if duration <= 2 {
		return 100
	} else if duration <= 6 {
		return 80
	} else if duration <= 12 {
		return 60
	} else if duration <= 24 {
		return 40
	} else if duration <= 48 {
		return 20
	} else {
		return 10
	}
}

func (s *DomainCalculationService) ProcessWaitingList(slotID int, availableSlots int) ([]models.WaitingListEntry, error) {
	slot, err := s.flightSlotRepo.GetByID(slotID)
	if err != nil {
		return nil, err
	}

	if slot.OccupiedCount+availableSlots > slot.Capacity {
		availableSlots = slot.Capacity - slot.OccupiedCount
	}

	if availableSlots <= 0 {
		return nil, nil
	}

	entries, err := s.waitingListRepo.GetTopEntries(availableSlots)
	if err != nil {
		return nil, err
	}

	var processedEntries []models.WaitingListEntry
	for _, entry := range entries {
		entry.Status = "admitted"
		entry.ActualEntryTime = time.Now()
		err = s.waitingListRepo.Update(&entry)
		if err != nil {
			continue
		}
		processedEntries = append(processedEntries, entry)
	}

	slot.OccupiedCount += len(processedEntries)
	s.flightSlotRepo.Update(&slot)

	return processedEntries, nil
}

func (s *DomainCalculationService) CheckConsistency(appointmentID int) (map[string]interface{}, error) {
	appointment, err := s.appointmentRepo.GetByID(appointmentID)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	result["is_consistent"] = true
	result["missing_fields"] = []string{}
	result["warnings"] = []string{}

	if appointment.MemberBenefitID == 0 {
		result["is_consistent"] = false
		result["missing_fields"] = append(result["missing_fields"].([]string), "member_benefit_id")
	}

	if appointment.FlightNumber == "" {
		result["is_consistent"] = false
		result["missing_fields"] = append(result["missing_fields"].([]string), "flight_number")
	}

	if appointment.FlightDate == "" {
		result["is_consistent"] = false
		result["missing_fields"] = append(result["missing_fields"].([]string), "flight_date")
	}

	if appointment.ScheduledTime.IsZero() {
		result["warnings"] = append(result["warnings"].([]string), "scheduled_time not set")
	}

	if !result["is_consistent"].(bool) {
		result["can_save_as_draft"] = true
		result["error_message"] = "Missing required fields for verification"
	}

	return result, nil
}

func (s *DomainCalculationService) DetectExpiredBenefits() ([]models.MemberBenefit, error) {
	benefits, err := s.memberRepo.GetByStatus("active")
	if err != nil {
		return nil, err
	}

	now := time.Now()
	var expiredBenefits []models.MemberBenefit

	for _, benefit := range benefits {
		if !benefit.ValidUntil.IsZero() && benefit.ValidUntil.Before(now) {
			expiredBenefits = append(expiredBenefits, benefit)
		}
	}

	return expiredBenefits, nil
}

func (s *DomainCalculationService) CalculateStatistics(dateFrom, dateTo string) (map[string]interface{}, error) {
	stats, err := s.verificationRepo.GetStatistics(dateFrom, dateTo)
	if err != nil {
		return nil, err
	}

	waitingEntries, _, _ := s.waitingListRepo.GetAll(1, 1000)
	totalWaiting := 0
	admittedCount := 0
	for _, entry := range waitingEntries {
		if entry.Status == "waiting" {
			totalWaiting++
		} else if entry.Status == "admitted" {
			admittedCount++
		}
	}

	stats["waiting_list_count"] = totalWaiting
	stats["admitted_from_waiting"] = admittedCount

	if totalWaiting > 0 {
		stats["admission_rate"] = math.Round(float64(admittedCount)/float64(totalWaiting)*10000) / 100
	} else {
		stats["admission_rate"] = 0.0
	}

	activeBenefits, _ := s.memberRepo.GetByStatus("active")
	stats["active_benefits_count"] = len(activeBenefits)

	var usedVouchers int
	for _, benefit := range activeBenefits {
		if benefit.RemainingQuota == 0 {
			usedVouchers++
		}
	}

	if len(activeBenefits) > 0 {
		stats["usage_rate"] = math.Round(float64(usedVouchers)/float64(len(activeBenefits))*10000) / 100
	} else {
		stats["usage_rate"] = 0.0
	}

	return stats, nil
}
