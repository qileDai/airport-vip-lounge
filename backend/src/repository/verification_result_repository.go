package repository

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/models"
)

type VerificationResultRepository struct{}

func NewVerificationResultRepository() *VerificationResultRepository {
	return &VerificationResultRepository{}
}

func (r *VerificationResultRepository) Create(result *models.VerificationResult) error {
	query := `INSERT INTO verification_results (code, name, status, appointment_record_id, member_benefit_id, verification_time, verifier, result, rejection_reason, rule_hits, score, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	resultObj, err := database.DB.Exec(query,
		result.Code,
		result.Name,
		result.Status,
		result.AppointmentRecordID,
		result.MemberBenefitID,
		result.VerificationTime,
		result.Verifier,
		result.Result,
		result.RejectionReason,
		result.RuleHits,
		result.Score,
		result.Owner,
		result.BatchID,
		result.Notes,
	)
	if err != nil {
		return err
	}
	id, _ := resultObj.LastInsertId()
	result.ID = int(id)
	return nil
}

func (r *VerificationResultRepository) GetByID(id int) (*models.VerificationResult, error) {
	query := `SELECT id, code, name, status, appointment_record_id, member_benefit_id, verification_time, verifier, result, rejection_reason, rule_hits, score, owner, batch_id, notes, created_at, updated_at FROM verification_results WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var result models.VerificationResult
	err := row.Scan(
		&result.ID,
		&result.Code,
		&result.Name,
		&result.Status,
		&result.AppointmentRecordID,
		&result.MemberBenefitID,
		&result.VerificationTime,
		&result.Verifier,
		&result.Result,
		&result.RejectionReason,
		&result.RuleHits,
		&result.Score,
		&result.Owner,
		&result.BatchID,
		&result.Notes,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *VerificationResultRepository) GetAll(page, pageSize int) ([]models.VerificationResult, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM verification_results`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, code, name, status, appointment_record_id, member_benefit_id, verification_time, verifier, result, rejection_reason, rule_hits, score, owner, batch_id, notes, created_at, updated_at FROM verification_results ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var results []models.VerificationResult
	for rows.Next() {
		var result models.VerificationResult
		err := rows.Scan(
			&result.ID,
			&result.Code,
			&result.Name,
			&result.Status,
			&result.AppointmentRecordID,
			&result.MemberBenefitID,
			&result.VerificationTime,
			&result.Verifier,
			&result.Result,
			&result.RejectionReason,
			&result.RuleHits,
			&result.Score,
			&result.Owner,
			&result.BatchID,
			&result.Notes,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, result)
	}

	return results, totalCount, nil
}

func (r *VerificationResultRepository) Update(result *models.VerificationResult) error {
	query := `UPDATE verification_results SET name=?, status=?, appointment_record_id=?, member_benefit_id=?, verification_time=?, verifier=?, result=?, rejection_reason=?, rule_hits=?, score=?, owner=?, notes=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := database.DB.Exec(query,
		result.Name,
		result.Status,
		result.AppointmentRecordID,
		result.MemberBenefitID,
		result.VerificationTime,
		result.Verifier,
		result.Result,
		result.RejectionReason,
		result.RuleHits,
		result.Score,
		result.Owner,
		result.Notes,
		result.ID,
	)
	return err
}

func (r *VerificationResultRepository) Delete(id int) error {
	query := `DELETE FROM verification_results WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}

func (r *VerificationResultRepository) GetByAppointmentRecordID(appointmentRecordID int) (*models.VerificationResult, error) {
	query := `SELECT id, code, name, status, appointment_record_id, member_benefit_id, verification_time, verifier, result, rejection_reason, rule_hits, score, owner, batch_id, notes, created_at, updated_at FROM verification_results WHERE appointment_record_id = ?`
	row := database.DB.QueryRow(query, appointmentRecordID)

	var result models.VerificationResult
	err := row.Scan(
		&result.ID,
		&result.Code,
		&result.Name,
		&result.Status,
		&result.AppointmentRecordID,
		&result.MemberBenefitID,
		&result.VerificationTime,
		&result.Verifier,
		&result.Result,
		&result.RejectionReason,
		&result.RuleHits,
		&result.Score,
		&result.Owner,
		&result.BatchID,
		&result.Notes,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *VerificationResultRepository) GetStatistics(dateFrom, dateTo string) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	var totalVerifications int
	countQuery := `SELECT COUNT(*) FROM verification_results WHERE verification_time BETWEEN ? AND ?`
	err := database.DB.QueryRow(countQuery, dateFrom, dateTo).Scan(&totalVerifications)
	if err != nil {
		return nil, err
	}
	stats["total_verifications"] = totalVerifications

	var passedCount int
	passedQuery := `SELECT COUNT(*) FROM verification_results WHERE result = 'passed' AND verification_time BETWEEN ? AND ?`
	err = database.DB.QueryRow(passedQuery, dateFrom, dateTo).Scan(&passedCount)
	if err != nil {
		return nil, err
	}
	stats["passed_count"] = passedCount

	var failedCount int
	failedQuery := `SELECT COUNT(*) FROM verification_results WHERE result = 'failed' AND verification_time BETWEEN ? AND ?`
	err = database.DB.QueryRow(failedQuery, dateFrom, dateTo).Scan(&failedCount)
	if err != nil {
		return nil, err
	}
	stats["failed_count"] = failedCount

	if totalVerifications > 0 {
		passRate := float64(passedCount) / float64(totalVerifications) * 100
		stats["pass_rate"] = passRate
	} else {
		stats["pass_rate"] = 0.0
	}

	return stats, nil
}
