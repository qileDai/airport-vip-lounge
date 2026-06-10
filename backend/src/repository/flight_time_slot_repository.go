package repository

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/models"
)

type FlightTimeSlotRepository struct{}

func NewFlightTimeSlotRepository() *FlightTimeSlotRepository {
	return &FlightTimeSlotRepository{}
}

func (r *FlightTimeSlotRepository) Create(slot *models.FlightTimeSlot) error {
	query := `INSERT INTO flight_time_slots (code, name, status, airport_code, gate, slot_start, slot_end, capacity, occupied_count, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query,
		slot.Code,
		slot.Name,
		slot.Status,
		slot.AirportCode,
		slot.Gate,
		slot.SlotStart,
		slot.SlotEnd,
		slot.Capacity,
		slot.OccupiedCount,
		slot.Owner,
		slot.BatchID,
		slot.Notes,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	slot.ID = int(id)
	return nil
}

func (r *FlightTimeSlotRepository) GetByID(id int) (*models.FlightTimeSlot, error) {
	query := `SELECT id, code, name, status, airport_code, gate, slot_start, slot_end, capacity, occupied_count, owner, batch_id, notes, created_at, updated_at FROM flight_time_slots WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var slot models.FlightTimeSlot
	err := row.Scan(
		&slot.ID,
		&slot.Code,
		&slot.Name,
		&slot.Status,
		&slot.AirportCode,
		&slot.Gate,
		&slot.SlotStart,
		&slot.SlotEnd,
		&slot.Capacity,
		&slot.OccupiedCount,
		&slot.Owner,
		&slot.BatchID,
		&slot.Notes,
		&slot.CreatedAt,
		&slot.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &slot, nil
}

func (r *FlightTimeSlotRepository) GetAll(page, pageSize int) ([]models.FlightTimeSlot, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM flight_time_slots`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, code, name, status, airport_code, gate, slot_start, slot_end, capacity, occupied_count, owner, batch_id, notes, created_at, updated_at FROM flight_time_slots ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var slots []models.FlightTimeSlot
	for rows.Next() {
		var slot models.FlightTimeSlot
		err := rows.Scan(
			&slot.ID,
			&slot.Code,
			&slot.Name,
			&slot.Status,
			&slot.AirportCode,
			&slot.Gate,
			&slot.SlotStart,
			&slot.SlotEnd,
			&slot.Capacity,
			&slot.OccupiedCount,
			&slot.Owner,
			&slot.BatchID,
			&slot.Notes,
			&slot.CreatedAt,
			&slot.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		slots = append(slots, slot)
	}

	return slots, totalCount, nil
}

func (r *FlightTimeSlotRepository) Update(slot *models.FlightTimeSlot) error {
	query := `UPDATE flight_time_slots SET name=?, status=?, airport_code=?, gate=?, slot_start=?, slot_end=?, capacity=?, occupied_count=?, owner=?, notes=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := database.DB.Exec(query,
		slot.Name,
		slot.Status,
		slot.AirportCode,
		slot.Gate,
		slot.SlotStart,
		slot.SlotEnd,
		slot.Capacity,
		slot.OccupiedCount,
		slot.Owner,
		slot.Notes,
		slot.ID,
	)
	return err
}

func (r *FlightTimeSlotRepository) Delete(id int) error {
	query := `DELETE FROM flight_time_slots WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}

func (r *FlightTimeSlotRepository) GetByAirportCode(airportCode string) ([]models.FlightTimeSlot, error) {
	query := `SELECT id, code, name, status, airport_code, gate, slot_start, slot_end, capacity, occupied_count, owner, batch_id, notes, created_at, updated_at FROM flight_time_slots WHERE airport_code = ?`
	rows, err := database.DB.Query(query, airportCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slots []models.FlightTimeSlot
	for rows.Next() {
		var slot models.FlightTimeSlot
		err := rows.Scan(
			&slot.ID,
			&slot.Code,
			&slot.Name,
			&slot.Status,
			&slot.AirportCode,
			&slot.Gate,
			&slot.SlotStart,
			&slot.SlotEnd,
			&slot.Capacity,
			&slot.OccupiedCount,
			&slot.Owner,
			&slot.BatchID,
			&slot.Notes,
			&slot.CreatedAt,
			&slot.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		slots = append(slots, slot)
	}
	return slots, nil
}
