package repository

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/models"
)

type CompanionRepository struct{}

func NewCompanionRepository() *CompanionRepository {
	return &CompanionRepository{}
}

func (r *CompanionRepository) Create(companion *models.Companion) error {
	query := `INSERT INTO companions (code, name, status, appointment_record_id, id_type, id_number, member_level, relation_type, age, owner, batch_id, verification_status, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query,
		companion.Code,
		companion.Name,
		companion.Status,
		companion.AppointmentRecordID,
		companion.IDType,
		companion.IDNumber,
		companion.MemberLevel,
		companion.RelationType,
		companion.Age,
		companion.Owner,
		companion.BatchID,
		companion.VerificationStatus,
		companion.Notes,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	companion.ID = int(id)
	return nil
}

func (r *CompanionRepository) GetByID(id int) (*models.Companion, error) {
	query := `SELECT id, code, name, status, appointment_record_id, id_type, id_number, member_level, relation_type, age, owner, batch_id, verification_status, notes, created_at, updated_at FROM companions WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var companion models.Companion
	err := row.Scan(
		&companion.ID,
		&companion.Code,
		&companion.Name,
		&companion.Status,
		&companion.AppointmentRecordID,
		&companion.IDType,
		&companion.IDNumber,
		&companion.MemberLevel,
		&companion.RelationType,
		&companion.Age,
		&companion.Owner,
		&companion.BatchID,
		&companion.VerificationStatus,
		&companion.Notes,
		&companion.CreatedAt,
		&companion.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &companion, nil
}

func (r *CompanionRepository) GetAll(page, pageSize int) ([]models.Companion, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM companions`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, code, name, status, appointment_record_id, id_type, id_number, member_level, relation_type, age, owner, batch_id, verification_status, notes, created_at, updated_at FROM companions ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var companions []models.Companion
	for rows.Next() {
		var companion models.Companion
		err := rows.Scan(
			&companion.ID,
			&companion.Code,
			&companion.Name,
			&companion.Status,
			&companion.AppointmentRecordID,
			&companion.IDType,
			&companion.IDNumber,
			&companion.MemberLevel,
			&companion.RelationType,
			&companion.Age,
			&companion.Owner,
			&companion.BatchID,
			&companion.VerificationStatus,
			&companion.Notes,
			&companion.CreatedAt,
			&companion.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		companions = append(companions, companion)
	}

	return companions, totalCount, nil
}

func (r *CompanionRepository) Update(companion *models.Companion) error {
	query := `UPDATE companions SET name=?, status=?, appointment_record_id=?, id_type=?, id_number=?, member_level=?, relation_type=?, age=?, owner=?, verification_status=?, notes=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := database.DB.Exec(query,
		companion.Name,
		companion.Status,
		companion.AppointmentRecordID,
		companion.IDType,
		companion.IDNumber,
		companion.MemberLevel,
		companion.RelationType,
		companion.Age,
		companion.Owner,
		companion.VerificationStatus,
		companion.Notes,
		companion.ID,
	)
	return err
}

func (r *CompanionRepository) Delete(id int) error {
	query := `DELETE FROM companions WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}

func (r *CompanionRepository) GetByAppointmentRecordID(appointmentRecordID int) ([]models.Companion, error) {
	query := `SELECT id, code, name, status, appointment_record_id, id_type, id_number, member_level, relation_type, age, owner, batch_id, verification_status, notes, created_at, updated_at FROM companions WHERE appointment_record_id = ?`
	rows, err := database.DB.Query(query, appointmentRecordID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companions []models.Companion
	for rows.Next() {
		var companion models.Companion
		err := rows.Scan(
			&companion.ID,
			&companion.Code,
			&companion.Name,
			&companion.Status,
			&companion.AppointmentRecordID,
			&companion.IDType,
			&companion.IDNumber,
			&companion.MemberLevel,
			&companion.RelationType,
			&companion.Age,
			&companion.Owner,
			&companion.BatchID,
			&companion.VerificationStatus,
			&companion.Notes,
			&companion.CreatedAt,
			&companion.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		companions = append(companions, companion)
	}
	return companions, nil
}

func (r *CompanionRepository) BatchCreate(companions []models.Companion) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO companions (code, name, status, appointment_record_id, id_type, id_number, member_level, relation_type, age, owner, batch_id, verification_status, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	
	for _, companion := range companions {
		_, err = tx.Exec(query,
			companion.Code,
			companion.Name,
			companion.Status,
			companion.AppointmentRecordID,
			companion.IDType,
			companion.IDNumber,
			companion.MemberLevel,
			companion.RelationType,
			companion.Age,
			companion.Owner,
			companion.BatchID,
			companion.VerificationStatus,
			companion.Notes,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
