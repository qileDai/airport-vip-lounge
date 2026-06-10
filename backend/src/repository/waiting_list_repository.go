package repository

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/models"
)

type WaitingListRepository struct{}

func NewWaitingListRepository() *WaitingListRepository {
	return &WaitingListRepository{}
}

func (r *WaitingListRepository) Create(entry *models.WaitingListEntry) error {
	query := `INSERT INTO waiting_list (code, name, status, member_benefit_id, priority_score, queue_position, wait_start_time, expected_entry_time, actual_entry_time, reason, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query,
		entry.Code,
		entry.Name,
		entry.Status,
		entry.MemberBenefitID,
		entry.PriorityScore,
		entry.QueuePosition,
		entry.WaitStartTime,
		entry.ExpectedEntryTime,
		entry.ActualEntryTime,
		entry.Reason,
		entry.Owner,
		entry.BatchID,
		entry.Notes,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	entry.ID = int(id)
	return nil
}

func (r *WaitingListRepository) GetByID(id int) (*models.WaitingListEntry, error) {
	query := `SELECT id, code, name, status, member_benefit_id, priority_score, queue_position, wait_start_time, expected_entry_time, actual_entry_time, reason, owner, batch_id, notes, created_at, updated_at FROM waiting_list WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var entry models.WaitingListEntry
	err := row.Scan(
		&entry.ID,
		&entry.Code,
		&entry.Name,
		&entry.Status,
		&entry.MemberBenefitID,
		&entry.PriorityScore,
		&entry.QueuePosition,
		&entry.WaitStartTime,
		&entry.ExpectedEntryTime,
		&entry.ActualEntryTime,
		&entry.Reason,
		&entry.Owner,
		&entry.BatchID,
		&entry.Notes,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *WaitingListRepository) GetAll(page, pageSize int) ([]models.WaitingListEntry, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM waiting_list`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, code, name, status, member_benefit_id, priority_score, queue_position, wait_start_time, expected_entry_time, actual_entry_time, reason, owner, batch_id, notes, created_at, updated_at FROM waiting_list ORDER BY priority_score DESC, queue_position ASC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var entries []models.WaitingListEntry
	for rows.Next() {
		var entry models.WaitingListEntry
		err := rows.Scan(
			&entry.ID,
			&entry.Code,
			&entry.Name,
			&entry.Status,
			&entry.MemberBenefitID,
			&entry.PriorityScore,
			&entry.QueuePosition,
			&entry.WaitStartTime,
			&entry.ExpectedEntryTime,
			&entry.ActualEntryTime,
			&entry.Reason,
			&entry.Owner,
			&entry.BatchID,
			&entry.Notes,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		entries = append(entries, entry)
	}

	return entries, totalCount, nil
}

func (r *WaitingListRepository) Update(entry *models.WaitingListEntry) error {
	query := `UPDATE waiting_list SET name=?, status=?, member_benefit_id=?, priority_score=?, queue_position=?, wait_start_time=?, expected_entry_time=?, actual_entry_time=?, reason=?, owner=?, notes=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := database.DB.Exec(query,
		entry.Name,
		entry.Status,
		entry.MemberBenefitID,
		entry.PriorityScore,
		entry.QueuePosition,
		entry.WaitStartTime,
		entry.ExpectedEntryTime,
		entry.ActualEntryTime,
		entry.Reason,
		entry.Owner,
		entry.Notes,
		entry.ID,
	)
	return err
}

func (r *WaitingListRepository) Delete(id int) error {
	query := `DELETE FROM waiting_list WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}

func (r *WaitingListRepository) GetByStatus(status string) ([]models.WaitingListEntry, error) {
	query := `SELECT id, code, name, status, member_benefit_id, priority_score, queue_position, wait_start_time, expected_entry_time, actual_entry_time, reason, owner, batch_id, notes, created_at, updated_at FROM waiting_list WHERE status = ? ORDER BY priority_score DESC`
	rows, err := database.DB.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.WaitingListEntry
	for rows.Next() {
		var entry models.WaitingListEntry
		err := rows.Scan(
			&entry.ID,
			&entry.Code,
			&entry.Name,
			&entry.Status,
			&entry.MemberBenefitID,
			&entry.PriorityScore,
			&entry.QueuePosition,
			&entry.WaitStartTime,
			&entry.ExpectedEntryTime,
			&entry.ActualEntryTime,
			&entry.Reason,
			&entry.Owner,
			&entry.BatchID,
			&entry.Notes,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (r *WaitingListRepository) GetTopEntries(limit int) ([]models.WaitingListEntry, error) {
	query := `SELECT id, code, name, status, member_benefit_id, priority_score, queue_position, wait_start_time, expected_entry_time, actual_entry_time, reason, owner, batch_id, notes, created_at, updated_at FROM waiting_list WHERE status = 'waiting' ORDER BY priority_score DESC, queue_position ASC LIMIT ?`
	rows, err := database.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.WaitingListEntry
	for rows.Next() {
		var entry models.WaitingListEntry
		err := rows.Scan(
			&entry.ID,
			&entry.Code,
			&entry.Name,
			&entry.Status,
			&entry.MemberBenefitID,
			&entry.PriorityScore,
			&entry.QueuePosition,
			&entry.WaitStartTime,
			&entry.ExpectedEntryTime,
			&entry.ActualEntryTime,
			&entry.Reason,
			&entry.Owner,
			&entry.BatchID,
			&entry.Notes,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
