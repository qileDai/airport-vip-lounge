package repository

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/models"
)

type ExceptionEventRepository struct{}

func NewExceptionEventRepository() *ExceptionEventRepository {
	return &ExceptionEventRepository{}
}

func (r *ExceptionEventRepository) Create(event *models.ExceptionEvent) error {
	query := `INSERT INTO exception_events (code, name, status, exception_type, severity, entity_type, entity_id, trigger_field, threshold_value, actual_value, handler, deadline, resolution, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query,
		event.Code,
		event.Name,
		event.Status,
		event.ExceptionType,
		event.Severity,
		event.EntityType,
		event.EntityID,
		event.TriggerField,
		event.ThresholdValue,
		event.ActualValue,
		event.Handler,
		event.Deadline,
		event.Resolution,
		event.Owner,
		event.BatchID,
		event.Notes,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	event.ID = int(id)
	return nil
}

func (r *ExceptionEventRepository) GetByID(id int) (*models.ExceptionEvent, error) {
	query := `SELECT id, code, name, status, exception_type, severity, entity_type, entity_id, trigger_field, threshold_value, actual_value, handler, deadline, resolution, owner, batch_id, notes, created_at, updated_at FROM exception_events WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var event models.ExceptionEvent
	err := row.Scan(
		&event.ID,
		&event.Code,
		&event.Name,
		&event.Status,
		&event.ExceptionType,
		&event.Severity,
		&event.EntityType,
		&event.EntityID,
		&event.TriggerField,
		&event.ThresholdValue,
		&event.ActualValue,
		&event.Handler,
		&event.Deadline,
		&event.Resolution,
		&event.Owner,
		&event.BatchID,
		&event.Notes,
		&event.CreatedAt,
		&event.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *ExceptionEventRepository) GetAll(page, pageSize int) ([]models.ExceptionEvent, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM exception_events`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, code, name, status, exception_type, severity, entity_type, entity_id, trigger_field, threshold_value, actual_value, handler, deadline, resolution, owner, batch_id, notes, created_at, updated_at FROM exception_events ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var events []models.ExceptionEvent
	for rows.Next() {
		var event models.ExceptionEvent
		err := rows.Scan(
			&event.ID,
			&event.Code,
			&event.Name,
			&event.Status,
			&event.ExceptionType,
			&event.Severity,
			&event.EntityType,
			&event.EntityID,
			&event.TriggerField,
			&event.ThresholdValue,
			&event.ActualValue,
			&event.Handler,
			&event.Deadline,
			&event.Resolution,
			&event.Owner,
			&event.BatchID,
			&event.Notes,
			&event.CreatedAt,
			&event.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		events = append(events, event)
	}

	return events, totalCount, nil
}

func (r *ExceptionEventRepository) Update(event *models.ExceptionEvent) error {
	query := `UPDATE exception_events SET name=?, status=?, exception_type=?, severity=?, entity_type=?, entity_id=?, trigger_field=?, threshold_value=?, actual_value=?, handler=?, deadline=?, resolution=?, owner=?, notes=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := database.DB.Exec(query,
		event.Name,
		event.Status,
		event.ExceptionType,
		event.Severity,
		event.EntityType,
		event.EntityID,
		event.TriggerField,
		event.ThresholdValue,
		event.ActualValue,
		event.Handler,
		event.Deadline,
		event.Resolution,
		event.Owner,
		event.Notes,
		event.ID,
	)
	return err
}

func (r *ExceptionEventRepository) Delete(id int) error {
	query := `DELETE FROM exception_events WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}

func (r *ExceptionEventRepository) GetByStatus(status string) ([]models.ExceptionEvent, error) {
	query := `SELECT id, code, name, status, exception_type, severity, entity_type, entity_id, trigger_field, threshold_value, actual_value, handler, deadline, resolution, owner, batch_id, notes, created_at, updated_at FROM exception_events WHERE status = ? ORDER BY severity DESC, created_at DESC`
	rows, err := database.DB.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.ExceptionEvent
	for rows.Next() {
		var event models.ExceptionEvent
		err := rows.Scan(
			&event.ID,
			&event.Code,
			&event.Name,
			&event.Status,
			&event.ExceptionType,
			&event.Severity,
			&event.EntityType,
			&event.EntityID,
			&event.TriggerField,
			&event.ThresholdValue,
			&event.ActualValue,
			&event.Handler,
			&event.Deadline,
			&event.Resolution,
			&event.Owner,
			&event.BatchID,
			&event.Notes,
			&event.CreatedAt,
			&event.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

type AuditLogRepository struct{}

func NewAuditLogRepository() *AuditLogRepository {
	return &AuditLogRepository{}
}

func (r *AuditLogRepository) Create(log *models.AuditLog) error {
	query := `INSERT INTO audit_logs (action, entity_type, entity_id, user_id, details, ip_address, user_agent) VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query,
		log.Action,
		log.EntityType,
		log.EntityID,
		log.UserID,
		log.Details,
		log.IPAddress,
		log.UserAgent,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	log.ID = int(id)
	return nil
}

func (r *AuditLogRepository) GetAll(page, pageSize int) ([]models.AuditLog, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM audit_logs`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, action, entity_type, entity_id, user_id, details, ip_address, user_agent, created_at FROM audit_logs ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var logs []models.AuditLog
	for rows.Next() {
		var log models.AuditLog
		err := rows.Scan(
			&log.ID,
			&log.Action,
			&log.EntityType,
			&log.EntityID,
			&log.UserID,
			&log.Details,
			&log.IPAddress,
			&log.UserAgent,
			&log.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		logs = append(logs, log)
	}

	return logs, totalCount, nil
}

func (r *AuditLogRepository) GetByEntityType(entityType string, page, pageSize int) ([]models.AuditLog, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM audit_logs WHERE entity_type = ?`
	err := database.DB.QueryRow(countQuery, entityType).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, action, entity_type, entity_id, user_id, details, ip_address, user_agent, created_at FROM audit_logs WHERE entity_type = ? ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, entityType, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var logs []models.AuditLog
	for rows.Next() {
		var log models.AuditLog
		err := rows.Scan(
			&log.ID,
			&log.Action,
			&log.EntityType,
			&log.EntityID,
			&log.UserID,
			&log.Details,
			&log.IPAddress,
			&log.UserAgent,
			&log.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		logs = append(logs, log)
	}

	return logs, totalCount, nil
}
