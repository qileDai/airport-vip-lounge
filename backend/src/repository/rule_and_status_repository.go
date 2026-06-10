package repository

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/models"
)

type StatusTransitionRecordRepository struct{}

func NewStatusTransitionRecordRepository() *StatusTransitionRecordRepository {
	return &StatusTransitionRecordRepository{}
}

func (r *StatusTransitionRecordRepository) Create(record *models.StatusTransitionRecord) error {
	query := `INSERT INTO status_transition_records (entity_type, entity_id, from_status, to_status, action, reason, operator, batch_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query,
		record.EntityType,
		record.EntityID,
		record.FromStatus,
		record.ToStatus,
		record.Action,
		record.Reason,
		record.Operator,
		record.BatchID,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	record.ID = int(id)
	return nil
}

func (r *StatusTransitionRecordRepository) GetByEntity(entityType string, entityID int) ([]models.StatusTransitionRecord, error) {
	query := `SELECT id, entity_type, entity_id, from_status, to_status, action, reason, operator, batch_id, created_at FROM status_transition_records WHERE entity_type = ? AND entity_id = ? ORDER BY created_at DESC`
	rows, err := database.DB.Query(query, entityType, entityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.StatusTransitionRecord
	for rows.Next() {
		var record models.StatusTransitionRecord
		err := rows.Scan(
			&record.ID,
			&record.EntityType,
			&record.EntityID,
			&record.FromStatus,
			&record.ToStatus,
			&record.Action,
			&record.Reason,
			&record.Operator,
			&record.BatchID,
			&record.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (r *StatusTransitionRecordRepository) GetAll(page, pageSize int) ([]models.StatusTransitionRecord, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM status_transition_records`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, entity_type, entity_id, from_status, to_status, action, reason, operator, batch_id, created_at FROM status_transition_records ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var records []models.StatusTransitionRecord
	for rows.Next() {
		var record models.StatusTransitionRecord
		err := rows.Scan(
			&record.ID,
			&record.EntityType,
			&record.EntityID,
			&record.FromStatus,
			&record.ToStatus,
			&record.Action,
			&record.Reason,
			&record.Operator,
			&record.BatchID,
			&record.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		records = append(records, record)
	}

	return records, totalCount, nil
}

type RuleConfigurationRepository struct{}

func NewRuleConfigurationRepository() *RuleConfigurationRepository {
	return &RuleConfigurationRepository{}
}

func (r *RuleConfigurationRepository) Create(rule *models.RuleConfiguration) error {
	query := `INSERT INTO rule_configurations (code, name, status, rule_type, priority, condition_expression, action_expression, description, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query,
		rule.Code,
		rule.Name,
		rule.Status,
		rule.RuleType,
		rule.Priority,
		rule.ConditionExpression,
		rule.ActionExpression,
		rule.Description,
		rule.Owner,
		rule.BatchID,
		rule.Notes,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	rule.ID = int(id)
	return nil
}

func (r *RuleConfigurationRepository) GetByID(id int) (*models.RuleConfiguration, error) {
	query := `SELECT id, code, name, status, rule_type, priority, condition_expression, action_expression, description, owner, batch_id, notes, created_at, updated_at FROM rule_configurations WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var rule models.RuleConfiguration
	err := row.Scan(
		&rule.ID,
		&rule.Code,
		&rule.Name,
		&rule.Status,
		&rule.RuleType,
		&rule.Priority,
		&rule.ConditionExpression,
		&rule.ActionExpression,
		&rule.Description,
		&rule.Owner,
		&rule.BatchID,
		&rule.Notes,
		&rule.CreatedAt,
		&rule.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &rule, nil
}

func (r *RuleConfigurationRepository) GetAll(page, pageSize int) ([]models.RuleConfiguration, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM rule_configurations`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, code, name, status, rule_type, priority, condition_expression, action_expression, description, owner, batch_id, notes, created_at, updated_at FROM rule_configurations ORDER BY priority DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var rules []models.RuleConfiguration
	for rows.Next() {
		var rule models.RuleConfiguration
		err := rows.Scan(
			&rule.ID,
			&rule.Code,
			&rule.Name,
			&rule.Status,
			&rule.RuleType,
			&rule.Priority,
			&rule.ConditionExpression,
			&rule.ActionExpression,
			&rule.Description,
			&rule.Owner,
			&rule.BatchID,
			&rule.Notes,
			&rule.CreatedAt,
			&rule.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		rules = append(rules, rule)
	}

	return rules, totalCount, nil
}

func (r *RuleConfigurationRepository) Update(rule *models.RuleConfiguration) error {
	query := `UPDATE rule_configurations SET name=?, status=?, rule_type=?, priority=?, condition_expression=?, action_expression=?, description=?, owner=?, notes=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := database.DB.Exec(query,
		rule.Name,
		rule.Status,
		rule.RuleType,
		rule.Priority,
		rule.ConditionExpression,
		rule.ActionExpression,
		rule.Description,
		rule.Owner,
		rule.Notes,
		rule.ID,
	)
	return err
}

func (r *RuleConfigurationRepository) Delete(id int) error {
	query := `DELETE FROM rule_configurations WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}

func (r *RuleConfigurationRepository) GetActiveRules() ([]models.RuleConfiguration, error) {
	query := `SELECT id, code, name, status, rule_type, priority, condition_expression, action_expression, description, owner, batch_id, notes, created_at, updated_at FROM rule_configurations WHERE status = 'active' ORDER BY priority DESC`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []models.RuleConfiguration
	for rows.Next() {
		var rule models.RuleConfiguration
		err := rows.Scan(
			&rule.ID,
			&rule.Code,
			&rule.Name,
			&rule.Status,
			&rule.RuleType,
			&rule.Priority,
			&rule.ConditionExpression,
			&rule.ActionExpression,
			&rule.Description,
			&rule.Owner,
			&rule.BatchID,
			&rule.Notes,
			&rule.CreatedAt,
			&rule.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, nil
}
