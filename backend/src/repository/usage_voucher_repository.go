package repository

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/models"
)

type UsageVoucherRepository struct{}

func NewUsageVoucherRepository() *UsageVoucherRepository {
	return &UsageVoucherRepository{}
}

func (r *UsageVoucherRepository) Create(voucher *models.UsageVoucher) error {
	query := `INSERT INTO usage_vouchers (code, name, status, voucher_type, member_benefit_id, appointment_record_id, issue_time, expire_time, use_time, qr_code, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query,
		voucher.Code,
		voucher.Name,
		voucher.Status,
		voucher.VoucherType,
		voucher.MemberBenefitID,
		voucher.AppointmentRecordID,
		voucher.IssueTime,
		voucher.ExpireTime,
		voucher.UseTime,
		voucher.QRCode,
		voucher.Owner,
		voucher.BatchID,
		voucher.Notes,
	)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	voucher.ID = int(id)
	return nil
}

func (r *UsageVoucherRepository) GetByID(id int) (*models.UsageVoucher, error) {
	query := `SELECT id, code, name, status, voucher_type, member_benefit_id, appointment_record_id, issue_time, expire_time, use_time, qr_code, owner, batch_id, notes, created_at, updated_at FROM usage_vouchers WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var voucher models.UsageVoucher
	err := row.Scan(
		&voucher.ID,
		&voucher.Code,
		&voucher.Name,
		&voucher.Status,
		&voucher.VoucherType,
		&voucher.MemberBenefitID,
		&voucher.AppointmentRecordID,
		&voucher.IssueTime,
		&voucher.ExpireTime,
		&voucher.UseTime,
		&voucher.QRCode,
		&voucher.Owner,
		&voucher.BatchID,
		&voucher.Notes,
		&voucher.CreatedAt,
		&voucher.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}

func (r *UsageVoucherRepository) GetAll(page, pageSize int) ([]models.UsageVoucher, int, error) {
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM usage_vouchers`
	err := database.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, code, name, status, voucher_type, member_benefit_id, appointment_record_id, issue_time, expire_time, use_time, qr_code, owner, batch_id, notes, created_at, updated_at FROM usage_vouchers ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := database.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var vouchers []models.UsageVoucher
	for rows.Next() {
		var voucher models.UsageVoucher
		err := rows.Scan(
			&voucher.ID,
			&voucher.Code,
			&voucher.Name,
			&voucher.Status,
			&voucher.VoucherType,
			&voucher.MemberBenefitID,
			&voucher.AppointmentRecordID,
			&voucher.IssueTime,
			&voucher.ExpireTime,
			&voucher.UseTime,
			&voucher.QRCode,
			&voucher.Owner,
			&voucher.BatchID,
			&voucher.Notes,
			&voucher.CreatedAt,
			&voucher.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		vouchers = append(vouchers, voucher)
	}

	return vouchers, totalCount, nil
}

func (r *UsageVoucherRepository) Update(voucher *models.UsageVoucher) error {
	query := `UPDATE usage_vouchers SET name=?, status=?, voucher_type=?, member_benefit_id=?, appointment_record_id=?, issue_time=?, expire_time=?, use_time=?, qr_code=?, owner=?, notes=?, updated_at=CURRENT_TIMESTAMP WHERE id=?`
	_, err := database.DB.Exec(query,
		voucher.Name,
		voucher.Status,
		voucher.VoucherType,
		voucher.MemberBenefitID,
		voucher.AppointmentRecordID,
		voucher.IssueTime,
		voucher.ExpireTime,
		voucher.UseTime,
		voucher.QRCode,
		voucher.Owner,
		voucher.Notes,
		voucher.ID,
	)
	return err
}

func (r *UsageVoucherRepository) Delete(id int) error {
	query := `DELETE FROM usage_vouchers WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}

func (r *UsageVoucherRepository) GetByStatus(status string) ([]models.UsageVoucher, error) {
	query := `SELECT id, code, name, status, voucher_type, member_benefit_id, appointment_record_id, issue_time, expire_time, use_time, qr_code, owner, batch_id, notes, created_at, updated_at FROM usage_vouchers WHERE status = ?`
	rows, err := database.DB.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vouchers []models.UsageVoucher
	for rows.Next() {
		var voucher models.UsageVoucher
		err := rows.Scan(
			&voucher.ID,
			&voucher.Code,
			&voucher.Name,
			&voucher.Status,
			&voucher.VoucherType,
			&voucher.MemberBenefitID,
			&voucher.AppointmentRecordID,
			&voucher.IssueTime,
			&voucher.ExpireTime,
			&voucher.UseTime,
			&voucher.QRCode,
			&voucher.Owner,
			&voucher.BatchID,
			&voucher.Notes,
			&voucher.CreatedAt,
			&voucher.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		vouchers = append(vouchers, voucher)
	}
	return vouchers, nil
}
