package models

import "time"

type MemberBenefit struct {
	ID             int        `json:"id"`
	Code           string     `json:"code"`
	Name           string     `json:"name"`
	Status         string     `json:"status"`
	MemberLevel    string     `json:"member_level"`
	RemainingQuota int        `json:"remaining_quota"`
	ValidFrom      time.Time  `json:"valid_from"`
	ValidUntil     time.Time  `json:"valid_until"`
	Owner          string     `json:"owner"`
	BatchID        string     `json:"batch_id"`
	AirportCode    string     `json:"airport_code"`
	LoungeID       string     `json:"lounge_id"`
	Notes          string     `json:"notes"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type AppointmentRecord struct {
	ID                 int       `json:"id"`
	Code               string    `json:"code"`
	Name               string    `json:"name"`
	Status             string    `json:"status"`
	MemberBenefitID    int       `json:"member_benefit_id"`
	FlightNumber       string    `json:"flight_number"`
	FlightDate         string    `json:"flight_date"`
	ScheduledTime      time.Time `json:"scheduled_time"`
	ActualArrivalTime  time.Time `json:"actual_arrival_time"`
	CompanionCount     int       `json:"companion_count"`
	Owner              string    `json:"owner"`
	BatchID            string    `json:"batch_id"`
	VerificationResultID int     `json:"verification_result_id"`
	Notes              string    `json:"notes"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type FlightTimeSlot struct {
	ID            int       `json:"id"`
	Code          string    `json:"code"`
	Name          string    `json:"name"`
	Status        string    `json:"status"`
	AirportCode   string    `json:"airport_code"`
	Gate          string    `json:"gate"`
	SlotStart     time.Time `json:"slot_start"`
	SlotEnd       time.Time `json:"slot_end"`
	Capacity      int       `json:"capacity"`
	OccupiedCount int       `json:"occupied_count"`
	Owner         string    `json:"owner"`
	BatchID       string    `json:"batch_id"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Companion struct {
	ID                int       `json:"id"`
	Code              string    `json:"code"`
	Name              string    `json:"name"`
	Status            string    `json:"status"`
	AppointmentRecordID int     `json:"appointment_record_id"`
	IDType            string    `json:"id_type"`
	IDNumber          string    `json:"id_number"`
	MemberLevel       string    `json:"member_level"`
	RelationType      string    `json:"relation_type"`
	Age               int       `json:"age"`
	Owner             string    `json:"owner"`
	BatchID           string    `json:"batch_id"`
	VerificationStatus string   `json:"verification_status"`
	Notes             string    `json:"notes"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type UsageVoucher struct {
	ID                int       `json:"id"`
	Code              string    `json:"code"`
	Name              string    `json:"name"`
	Status            string    `json:"status"`
	VoucherType       string    `json:"voucher_type"`
	MemberBenefitID   int       `json:"member_benefit_id"`
	AppointmentRecordID int     `json:"appointment_record_id"`
	IssueTime         time.Time `json:"issue_time"`
	ExpireTime        time.Time `json:"expire_time"`
	UseTime           time.Time `json:"use_time"`
	QRCode            string    `json:"qr_code"`
	Owner             string    `json:"owner"`
	BatchID           string    `json:"batch_id"`
	Notes             string    `json:"notes"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type WaitingListEntry struct {
	ID               int       `json:"id"`
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	Status           string    `json:"status"`
	MemberBenefitID  int       `json:"member_benefit_id"`
	PriorityScore    float64   `json:"priority_score"`
	QueuePosition    int       `json:"queue_position"`
	WaitStartTime    time.Time `json:"wait_start_time"`
	ExpectedEntryTime time.Time `json:"expected_entry_time"`
	ActualEntryTime  time.Time `json:"actual_entry_time"`
	Reason           string    `json:"reason"`
	Owner            string    `json:"owner"`
	BatchID          string    `json:"batch_id"`
	Notes            string    `json:"notes"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type VerificationResult struct {
	ID                 int       `json:"id"`
	Code               string    `json:"code"`
	Name               string    `json:"name"`
	Status             string    `json:"status"`
	AppointmentRecordID int      `json:"appointment_record_id"`
	MemberBenefitID    int       `json:"member_benefit_id"`
	VerificationTime   time.Time `json:"verification_time"`
	Verifier           string    `json:"verifier"`
	Result             string    `json:"result"`
	RejectionReason    string    `json:"rejection_reason"`
	RuleHits           string    `json:"rule_hits"`
	Score              float64   `json:"score"`
	Owner              string    `json:"owner"`
	BatchID            string    `json:"batch_id"`
	Notes              string    `json:"notes"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type StatusTransitionRecord struct {
	ID        int    `json:"id"`
	EntityType string `json:"entity_type"`
	EntityID  int    `json:"entity_id"`
	FromStatus string `json:"from_status"`
	ToStatus  string `json:"to_status"`
	Action    string `json:"action"`
	Reason    string `json:"reason"`
	Operator  string `json:"operator"`
	BatchID   string `json:"batch_id"`
	CreatedAt time.Time `json:"created_at"`
}

type RuleConfiguration struct {
	ID                   int       `json:"id"`
	Code                 string    `json:"code"`
	Name                 string    `json:"name"`
	Status               string    `json:"status"`
	RuleType             string    `json:"rule_type"`
	Priority             int       `json:"priority"`
	ConditionExpression  string    `json:"condition_expression"`
	ActionExpression     string    `json:"action_expression"`
	Description          string    `json:"description"`
	Owner                string    `json:"owner"`
	BatchID              string    `json:"batch_id"`
	Notes                string    `json:"notes"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type ExceptionEvent struct {
	ID             int       `json:"id"`
	Code           string    `json:"code"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	ExceptionType  string    `json:"exception_type"`
	Severity       string    `json:"severity"`
	EntityType     string    `json:"entity_type"`
	EntityID       int       `json:"entity_id"`
	TriggerField   string    `json:"trigger_field"`
	ThresholdValue string    `json:"threshold_value"`
	ActualValue    string    `json:"actual_value"`
	Handler        string    `json:"handler"`
	Deadline       time.Time `json:"deadline"`
	Resolution     string    `json:"resolution"`
	Owner          string    `json:"owner"`
	BatchID        string    `json:"batch_id"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type AuditLog struct {
	ID        int    `json:"id"`
	Action    string `json:"action"`
	EntityType string `json:"entity_type"`
	EntityID  int    `json:"entity_id"`
	UserID    string `json:"user_id"`
	Details   string `json:"details"`
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	CreatedAt time.Time `json:"created_at"`
}
