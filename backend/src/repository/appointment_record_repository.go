package repository

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/models"
)

type AppointmentRecordRepository struct{}

func NewAppointmentRecordRepository() *AppointmentRecordRepository {
	return &AppointmentRecordRepository{}
}

func (r *AppointmentRecordRepository) Create(record *models.AppointmentRecord) error {
	query := `INSERT INTO appointment_records (code, name, status, member_benefit_id, flight_number, flight_date, scheduled_time, actual_arrival_time, companion_count, owner, batch_id, verification_result_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query,
		record.Code,
		record.Name,
		record.Status,
		record.MemberBenefitID,
		record.FlightNumber,
		record.FlightDate,
		record.ScheduledTime,
		record.ActualArrivalTime,
		record.CompanionCount,
		record.Owner,
		record.BatchID,
		record.VerificationResultID,
		record.Notes,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	record.ID = int(id)
	return nil
}

func (r *AppointmentRecordRepository) GetByID(id int) (*models.AppointmentRecord, error) {
	query := `SELECT id, code, name, status, member_benefit_id, flight_number, flight_date, scheduled_time, actual_arrival_time, companion_count, owner, batch_id, verification_result_id, notes, created_at, updated_at FROM appointment_records WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var record models.AppointmentRecord
	err := row.Scan(
		&record.ID,
		&record.Code,
		&record.Name,
		&record.Status,
		&record.MemberBenefitID,
		&record.FlightNumber,
		&record.FlightDate,
		&record.ScheduledTime,
		&record.ActualArrivalTime,
		&record.CompanionCount,
		&record.Owner,
		&record.BatchID,
		&record.VerificationResultID,
		&record.Notes,
		&record.CreatedAt,
		&record.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *AppointmentRecordRepository) GetAll(page, pageSize int) ([]models.AppointmentRecord, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM appointment_records`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, code, name, status, member_benefit_id, flight_number, flight_date, scheduled_time, actual_arrival_time, companion_count, owner, batch_id, verification_result_id, notes, created_at, updated_at FROM appointment_records ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var records []models.AppointmentRecord
	for rows.Next() {
		var record models.AppointmentRecord
		err := rows.Scan(
			&record.ID,
			&record.Code,
			&record.Name,
			&record.Status,
			&record.MemberBenefitID,
			&record.FlightNumber,
			&record.FlightDate,
			&record.ScheduledTime,
			&record.ActualArrivalTime,
			&record.CompanionCount,
			&record.Owner,
			&record.BatchID,
			&record.VerificationResultID,
			&record.Notes,
			&record.CreatedAt,
			&record.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		records = append(records, record)
	}

	return records, totalCount, nil
}

func (r *AppointmentRecordRepository) Update(record *models.AppointmentRecord) error {
	query := `UPDATE appointment_records SET name=?, status=?, member_benefit_id=?, flight_number=?, flight_date=?, scheduled_time=?, actual_arrival_time=?, companion_count=?, owner=?, verification_result_id=?, notes=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := database.DB.Exec(query,
		record.Name,
		record.Status,
		record.MemberBenefitID,
		record.FlightNumber,
		record.FlightDate,
		record.ScheduledTime,
		record.ActualArrivalTime,
		record.CompanionCount,
		record.Owner,
		record.VerificationResultID,
		record.Notes,
		record.ID,
	)
	return err
}

func (r *AppointmentRecordRepository) Delete(id int) error {
	query := `DELETE FROM appointment_records WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}

func (r *AppointmentRecordRepository) GetByMemberBenefitID(memberBenefitID int) ([]models.AppointmentRecord, error) {
	query := `SELECT id, code, name, status, member_benefit_id, flight_number, flight_date, scheduled_time, actual_arrival_time, companion_count, owner, batch_id, verification_result_id, notes, created_at, updated_at FROM appointment_records WHERE member_benefit_id = ?`
	rows, err := database.DB.Query(query, memberBenefitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.AppointmentRecord
	for rows.Next() {
		var record models.AppointmentRecord
		err := rows.Scan(
			&record.ID,
			&record.Code,
			&record.Name,
			&record.Status,
			&record.MemberBenefitID,
			&record.FlightNumber,
			&record.FlightDate,
			&record.ScheduledTime,
			&record.ActualArrivalTime,
			&record.CompanionCount,
			&record.Owner,
			&record.BatchID,
			&record.VerificationResultID,
			&record.Notes,
			&record.CreatedAt,
			&record.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
