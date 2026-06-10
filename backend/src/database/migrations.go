package database

func createMemberBenefitsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS member_benefits (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(200) NOT NULL,
		status VARCHAR(20) DEFAULT 'draft',
		member_level VARCHAR(20) NOT NULL,
		remaining_quota INTEGER DEFAULT 0,
		valid_from DATETIME,
		valid_until DATETIME,
		owner VARCHAR(100),
		batch_id VARCHAR(50),
		airport_code VARCHAR(10),
		lounge_id VARCHAR(50),
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	DB.Exec(query)
}

func createAppointmentRecordsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS appointment_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(200) NOT NULL,
		status VARCHAR(20) DEFAULT 'draft',
		member_benefit_id INTEGER,
		flight_number VARCHAR(20),
		flight_date DATE,
		scheduled_time DATETIME,
		actual_arrival_time DATETIME,
		companion_count INTEGER DEFAULT 0,
		owner VARCHAR(100),
		batch_id VARCHAR(50),
		verification_result_id INTEGER,
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (member_benefit_id) REFERENCES member_benefits(id)
	)`
	DB.Exec(query)
}

func createFlightTimeSlotsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS flight_time_slots (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(200) NOT NULL,
		status VARCHAR(20) DEFAULT 'active',
		airport_code VARCHAR(10) NOT NULL,
		gate VARCHAR(10),
		slot_start DATETIME,
		slot_end DATETIME,
		capacity INTEGER DEFAULT 0,
		occupied_count INTEGER DEFAULT 0,
		owner VARCHAR(100),
		batch_id VARCHAR(50),
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	DB.Exec(query)
}

func createCompanionsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS companions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(200) NOT NULL,
		status VARCHAR(20) DEFAULT 'pending',
		appointment_record_id INTEGER,
		id_type VARCHAR(20),
		id_number VARCHAR(50),
		member_level VARCHAR(20),
		relation_type VARCHAR(20),
		age INTEGER,
		owner VARCHAR(100),
		batch_id VARCHAR(50),
		verification_status VARCHAR(20),
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (appointment_record_id) REFERENCES appointment_records(id)
	)`
	DB.Exec(query)
}

func createUsageVouchersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS usage_vouchers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(200) NOT NULL,
		status VARCHAR(20) DEFAULT 'active'),
		voucher_type VARCHAR(30),
		member_benefit_id INTEGER,
		appointment_record_id INTEGER,
		issue_time DATETIME,
		expire_time DATETIME,
		use_time DATETIME,
		qr_code VARCHAR(200),
		owner VARCHAR(100),
		batch_id VARCHAR(50),
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (member_benefit_id) REFERENCES member_benefits(id),
		FOREIGN KEY (appointment_record_id) REFERENCES appointment_records(id)
	)`
	DB.Exec(query)
}

func createWaitingListTable() {
	query := `
	CREATE TABLE IF NOT EXISTS waiting_list (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(200) NOT NULL,
		status VARCHAR(20) DEFAULT 'waiting'),
		member_benefit_id INTEGER,
		priority_score REAL DEFAULT 0,
		queue_position INTEGER DEFAULT 0,
		wait_start_time DATETIME,
		expected_entry_time DATETIME,
		actual_entry_time DATETIME,
		reason VARCHAR(200),
		owner VARCHAR(100),
		batch_id VARCHAR(50),
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (member_benefit_id) REFERENCES member_benefits(id)
	)`
	DB.Exec(query)
}

func createVerificationResultsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS verification_results (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(200) NOT NULL,
		status VARCHAR(20) DEFAULT 'pending'),
		appointment_record_id INTEGER,
		member_benefit_id INTEGER,
		verification_time DATETIME,
		verifier VARCHAR(100),
		result VARCHAR(20),
		rejection_reason TEXT,
		rule_hits JSON,
		score REAL DEFAULT 0,
		owner VARCHAR(100),
		batch_id VARCHAR(50),
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (appointment_record_id) REFERENCES appointment_records(id),
		FOREIGN KEY (member_benefit_id) REFERENCES member_benefits(id)
	)`
	DB.Exec(query)
}

func createStatusTransitionRecordsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS status_transition_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		entity_type VARCHAR(50) NOT NULL,
		entity_id INTEGER NOT NULL,
		from_status VARCHAR(20),
		to_status VARCHAR(20) NOT NULL,
		action VARCHAR(50) NOT NULL,
		reason TEXT,
		operator VARCHAR(100),
		batch_id VARCHAR(50),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	DB.Exec(query)
}

func createRuleConfigurationsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS rule_configurations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(200) NOT NULL,
		status VARCHAR(20) DEFAULT 'active'),
		rule_type VARCHAR(30),
		priority INTEGER DEFAULT 0,
		condition_expression TEXT,
		action_expression TEXT,
		description TEXT,
		owner VARCHAR(100),
		batch_id VARCHAR(50),
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	DB.Exec(query)
}

func createExceptionEventsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS exception_events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(200) NOT NULL,
		status VARCHAR(20) DEFAULT 'open'),
		exception_type VARCHAR(30),
		severity VARCHAR(20),
		entity_type VARCHAR(50),
		entity_id INTEGER,
		trigger_field VARCHAR(50),
		threshold_value VARCHAR(100),
		actual_value VARCHAR(100),
		handler VARCHAR(100),
		deadline DATETIME,
		resolution TEXT,
		owner VARCHAR(100),
		batch_id VARCHAR(50),
		notes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	DB.Exec(query)
}

func createAuditLogsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS audit_logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		action VARCHAR(50) NOT NULL,
		entity_type VARCHAR(50),
		entity_id INTEGER,
		user_id VARCHAR(100),
		details TEXT,
		ip_address VARCHAR(50),
		user_agent TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	DB.Exec(query)
}
