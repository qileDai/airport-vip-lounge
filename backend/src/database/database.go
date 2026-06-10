package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./airport_vip_lounge.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	fmt.Println("Database connection established")
}

func RunMigrations() {
	createMemberBenefitsTable()
	createAppointmentRecordsTable()
	createFlightTimeSlotsTable()
	createCompanionsTable()
	createUsageVouchersTable()
	createWaitingListTable()
	createVerificationResultsTable()
	createStatusTransitionRecordsTable()
	createRuleConfigurationsTable()
	createExceptionEventsTable()
	createAuditLogsTable()

	fmt.Println("Database migrations completed")
}
