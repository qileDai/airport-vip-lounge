package dto

type CreateMemberBenefitRequest struct {
	Code           string `json:"code"`
	Name           string `json:"name"`
	MemberLevel    string `json:"member_level"`
	RemainingQuota int    `json:"remaining_quota"`
	ValidFrom      string `json:"valid_from"`
	ValidUntil     string `json:"valid_until"`
	Owner          string `json:"owner"`
	BatchID        string `json:"batch_id"`
	AirportCode    string `json:"airport_code"`
	LoungeID       string `json:"lounge_id"`
	Notes          string `json:"notes"`
}

type UpdateMemberBenefitRequest struct {
	Name           string `json:"name"`
	Status         string `json:"status"`
	MemberLevel    string `json:"member_level"`
	RemainingQuota int    `json:"remaining_quota"`
	ValidFrom      string `json:"valid_from"`
	ValidUntil     string `json:"valid_until"`
	Owner          string `json:"owner"`
	AirportCode    string `json:"airport_code"`
	LoungeID       string `json:"lounge_id"`
	Notes          string `json:"notes"`
}

type StatusTransitionRequest struct {
	ToStatus string `json:"to_status"`
	Action   string `json:"action"`
	Reason   string `json:"reason"`
}

type BatchImportRequest struct {
	Records []interface{} `json:"records"`
}

type VerificationRequest struct {
	AppointmentRecordID int `json:"appointment_record_id"`
	MemberBenefitID     int `json:"member_benefit_id"`
	Verifier            string `json:"verifier"`
}

type FilterRequest struct {
	Page         int                    `json:"page"`
	PageSize     int                    `json:"page_size"`
	Filters      map[string]interface{} `json:"filters"`
	SortBy       string                 `json:"sort_by"`
	SortOrder    string                 `json:"sort_order"`
	Search       string                 `json:"search"`
	DateFrom     string                 `json:"date_from"`
	DateTo       string                 `json:"date_to"`
}

type StatisticsRequest struct {
	GroupBy      string `json:"group_by"`
	DateFrom     string `json:"date_from"`
	DateTo       string `json:"date_to"`
	BatchID      string `json:"batch_id"`
	Role         string `json:"role"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
}
