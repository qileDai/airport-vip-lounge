package repository

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/models"
)

type MemberBenefitRepository struct{}

func NewMemberBenefitRepository() *MemberBenefitRepository {
	return &MemberBenefitRepository{}
}

func (r *MemberBenefitRepository) Create(benefit *models.MemberBenefit) error {
	query := `INSERT INTO member_benefits (code, name, status, member_level, remaining_quota, valid_from, valid_until, owner, batch_id, airport_code, lounge_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query,
		benefit.Code,
		benefit.Name,
		benefit.Status,
		benefit.MemberLevel,
		benefit.RemainingQuota,
		benefit.ValidFrom,
		benefit.ValidUntil,
		benefit.Owner,
		benefit.BatchID,
		benefit.AirportCode,
		benefit.LoungeID,
		benefit.Notes,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	benefit.ID = int(id)
	return nil
}

func (r *MemberBenefitRepository) GetByID(id int) (*models.MemberBenefit, error) {
	query := `SELECT id, code, name, status, member_level, remaining_quota, valid_from, valid_until, owner, batch_id, airport_code, lounge_id, notes, created_at, updated_at FROM member_benefits WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var benefit models.MemberBenefit
	err := row.Scan(
		&benefit.ID,
		&benefit.Code,
		&benefit.Name,
		&benefit.Status,
		&benefit.MemberLevel,
		&benefit.RemainingQuota,
		&benefit.ValidFrom,
		&benefit.ValidUntil,
		&benefit.Owner,
		&benefit.BatchID,
		&benefit.AirportCode,
		&benefit.LoungeID,
		&benefit.Notes,
		&benefit.CreatedAt,
		&benefit.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &benefit, nil
}

func (r *MemberBenefitRepository) GetAll(page, pageSize int) ([]models.MemberBenefit, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM member_benefits`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, code, name, status, member_level, remaining_quota, valid_from, valid_until, owner, batch_id, airport_code, lounge_id, notes, created_at, updated_at FROM member_benefits ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var benefits []models.MemberBenefit
	for rows.Next() {
		var benefit models.MemberBenefit
		err := rows.Scan(
			&benefit.ID,
			&benefit.Code,
			&benefit.Name,
			&benefit.Status,
			&benefit.MemberLevel,
			&benefit.RemainingQuota,
			&benefit.ValidFrom,
			&benefit.ValidUntil,
			&benefit.Owner,
			&benefit.BatchID,
			&benefit.AirportCode,
			&benefit.LoungeID,
			&benefit.Notes,
			&benefit.CreatedAt,
			&benefit.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		benefits = append(benefits, benefit)
	}

	return benefits, totalCount, nil
}

func (r *MemberBenefitRepository) Update(benefit *models.MemberBenefit) error {
	query := `UPDATE member_benefits SET name=?, status=?, member_level=?, remaining_quota=?, valid_from=?, valid_until=?, owner=?, airport_code=?, lounge_id=?, notes=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := database.DB.Exec(query,
		benefit.Name,
		benefit.Status,
		benefit.MemberLevel,
		benefit.RemainingQuota,
		benefit.ValidFrom,
		benefit.ValidUntil,
		benefit.Owner,
		benefit.AirportCode,
		benefit.LoungeID,
		benefit.Notes,
		benefit.ID,
	)
	return err
}

func (r *MemberBenefitRepository) Delete(id int) error {
	query := `DELETE FROM member_benefits WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}

func (r *MemberBenefitRepository) GetByStatus(status string) ([]models.MemberBenefit, error) {
	query := `SELECT id, code, name, status, member_level, remaining_quota, valid_from, valid_until, owner, batch_id, airport_code, lounge_id, notes, created_at, updated_at FROM member_benefits WHERE status = ?`
	rows, err := database.DB.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var benefits []models.MemberBenefit
	for rows.Next() {
		var benefit models.MemberBenefit
		err := rows.Scan(
			&benefit.ID,
			&benefit.Code,
			&benefit.Name,
			&benefit.Status,
			&benefit.MemberLevel,
			&benefit.RemainingQuota,
			&benefit.ValidFrom,
			&benefit.ValidUntil,
			&benefit.Owner,
			&benefit.BatchID,
			&benefit.AirportCode,
			&benefit.LoungeID,
			&benefit.Notes,
			&benefit.CreatedAt,
			&benefit.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		benefits = append(benefits, benefit)
	}
	return benefits, nil
}
