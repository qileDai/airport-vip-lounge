package handlers

import (
	"airport-vip-lounge/src/database"
	"airport-vip-lounge/src/dto"
	"airport-vip-lounge/src/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ResetSeedData(c *fiber.Ctx) error {
	database.DB.Exec("DELETE FROM status_transition_records")
	database.DB.Exec("DELETE FROM audit_logs")
	database.DB.Exec("DELETE FROM companions")
	database.DB.Exec("DELETE FROM usage_vouchers")
	database.DB.Exec("DELETE FROM verification_results")
	database.DB.Exec("DELETE FROM waiting_list")
	database.DB.Exec("DELETE FROM exception_events")
	database.DB.Exec("DELETE FROM appointment_records")
	database.DB.Exec("DELETE FROM flight_time_slots")
	database.DB.Exec("DELETE FROM rule_configurations")
	database.DB.Exec("DELETE FROM member_benefits")

	insertMemberBenefits()
	insertFlightTimeSlots()
	insertAppointmentRecords()
	insertCompanions()
	insertUsageVouchers()
	insertWaitingList()
	insertVerificationResults()
	insertRuleConfigurations()
	insertExceptionEvents()

	return c.JSON(dto.APIResponse{
		Success: true,
		Message: "Seed data reset successfully",
		Data: map[string]int{
			"member_benefits":      12,
			"flight_time_slots":    8,
			"appointment_records":  15,
			"companions":           18,
			"usage_vouchers":       15,
			"waiting_list":         10,
			"verification_results": 15,
			"rule_configurations":  6,
			"exception_events":     8,
		},
	})
}

func insertMemberBenefits() {
	benefits := []models.MemberBenefit{
		{Code: "MB-202601-PKU001", Name: "北京大学钻石会员权益包", Status: "active", MemberLevel: "diamond", RemainingQuota: 8, ValidFrom: parseTime("2026-01-01 00:00:00"), ValidUntil: parseTime("2026-12-31 23:59:59"), Owner: "张明远", BatchID: "BATCH-202601", AirportCode: "PEK", LoungeID: "PEK-T3-DOM-01", Notes: "年度钻石会员，含贵宾厅使用权益"},
		{Code: "MB-202601-SHFT002", Name: "上海浦东白金卡会员权益", Status: "active", MemberLevel: "platinum", RemainingQuota: 5, ValidFrom: parseTime("2026-02-01 00:00:00"), ValidUntil: parseTime("2027-01-31 23:59:59"), Owner: "李婷婷", BatchID: "BATCH-202602", AirportCode: "PVG", LoungeID: "PVG-T2-INT-03", Notes: "白金卡季度权益包"},
		{Code: "MB-202601-GZCN003", Name: "广州白云金卡会员服务", Status: "active", MemberLevel: "gold", RemainingQuota: 3, ValidFrom: parseTime("2026-03-01 00:00:00"), ValidUntil: parseTime("2026-09-30 23:59:59"), Owner: "王建国", BatchID: "BATCH-202603", AirportCode: "CAN", LoungeID: "CAN-T2-A-02", Notes: "金卡半年期服务"},
		{Code: "MB-202601-SZXZ004", Name: "深圳宝安银卡基础权益", Status: "active", MemberLevel: "silver", RemainingQuota: 2, ValidFrom: parseTime("2026-04-01 00:00:00"), ValidUntil: parseTime("2026-10-31 23:59:59"), Owner: "陈晓华", BatchID: "BATCH-202604", AirportCode: "SZX", LoungeID: "SZX-T3-05", Notes: "银卡基础版权益"},
		{Code: "MB-202601-CTUU005", Name: "成都天府铜卡体验券", Status: "active", MemberLevel: "bronze", RemainingQuota: 1, ValidFrom: parseTime("2026-05-01 00:00:00"), ValidUntil: parseTime("2026-07-31 23:59:59"), Owner: "刘思雨", BatchID: "BATCH-202605", AirportCode: "TFU", LoungeID: "TFU-T2-B-01", Notes: "新客户体验权益"},
		{Code: "MB-202601-HGHQ006", Name: "杭州萧山基本卡权益", Status: "expired", MemberLevel: "basic", RemainingQuota: 0, ValidFrom: parseTime("2025-06-01 00:00:00"), ValidUntil: parseTime("2026-05-20 23:59:59"), Owner: "赵文博", BatchID: "BATCH-202506", AirportCode: "HGH", LoungeID: "HGH-T3-C-04", Notes: "已过期的基础权益包"},
		{Code: "MB-202601-XIYN007", Name: "咸阳机场企业协议权益", Status: "pending_review", MemberLevel: "gold", RemainingQuota: 10, ValidFrom: parseTime("2026-06-01 00:00:00"), ValidUntil: parseTime("2026-12-31 23:59:59"), Owner: "孙丽萍", BatchID: "BATCH-202606", AirportCode: "XIY", LoungeID: "XIY-T3-A-01", Notes: "华为技术企业协议待审核"},
		{Code: "MB-202601-NKGW008", Name: "南京禄口高端商务权益", Status: "needs_info", MemberLevel: "diamond", RemainingQuota: 12, ValidFrom: parseTime("2026-07-01 00:00:00"), ValidUntil: parseTime("2027-06-30 23:59:59"), Owner: "周志强", BatchID: "BATCH-202607", AirportCode: "NKG", LoungeID: "NKG-T2-VIP-01", Notes: "缺少企业资质证明文件"},
		{Code: "MB-202601-WUSN009", Name: "无锡硕放区域联运权益", Status: "rejected", MemberLevel: "silver", RemainingQuota: 0, ValidFrom: parseTime("2026-08-01 00:00:00"), ValidUntil: parseTime("2027-01-31 23:59:59"), Owner: "吴美玲", BatchID: "BATCH-202608", AirportCode: "WUX", LoungeID: "WUX-T2-02", Notes: "不符合区域联运规则要求"},
		{Code: "MB-202601-DLCY010", Name: "大连周水子航空联盟权益", Status: "archived", MemberLevel: "gold", RemainingQuota: 0, ValidFrom: parseTime("2025-09-01 00:00:00"), ValidUntil: parseTime("2026-03-31 23:59:59"), Owner: "郑海涛", BatchID: "BATCH-202509", AirportCode: "DLC", LoungeID: "DLC-T2-INT-01", Notes: "已归档的历史权益记录"},
		{Code: "MB-202601-KMGZ011", name: "昆明长水旅游季特惠权益", Status: "suspended", MemberLevel: "bronze", RemainingQuota: 0, ValidFrom: parseTime("2026-10-01 00:00:00"), ValidUntil: parseTime("2026-11-30 23:59:59"), owner: "冯雪梅", BatchID: "BATCH-202610", AirportCode: "KMG", LoungeID: "KMG-T3-B-03", Notes: "因违规使用被暂停"},
		{Code: "MB-202601-CKGR012", Name: "重庆江北政企合作权益", Status: "active", MemberLevel: "platinum", RemainingQuota: 20, ValidFrom: parseTime("2026-11-01 00:00:00"), ValidUntil: parseTime("2027-10-31 23:59:59"), Owner: "杨建华", BatchID: "BATCH-202611", AirportCode: "CKG", LoungeID: "CKG-T3-VIP-02", Notes: "重庆市政务服务中心合作项目"},
	}

	for _, benefit := range benefits {
		database.DB.Exec(
			`INSERT INTO member_benefits (code, name, status, member_level, remaining_quota, valid_from, valid_until, owner, batch_id, airport_code, lounge_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			benefit.Code, benefit.Name, benefit.Status, benefit.MemberLevel, benefit.RemainingQuota,
			benefit.ValidFrom, benefit.ValidUntil, benefit.Owner, benefit.BatchID,
			benefit.AirportCode, benefit.LoungeID, benefit.Notes,
		)
	}
}

func insertFlightTimeSlots() {
	slots := []models.FlightTimeSlot{
		{Code: "FS-PEK-T3-20260525-M01", Name: "首都机场T3国内出发早高峰时段A", Status: "active", AirportCode: "PEK", Gate: "E23-E28", SlotStart: parseTime("2026-05-25 06:00:00"), SlotEnd: parseTime("2026-05-25 09:00:00"), Capacity: 50, OccupiedCount: 42, Owner: "值班经理王磊", BatchID: "SLOT-20260525", Notes: "早高峰核心时段"},
		{Code: "FS-PEK-T3-20260525-A01", Name: "首都机场T3国内到达午间时段B", Status: "active", AirportCode: "PEK", Gate: "E15-E20", SlotStart: parseTime("2026-05-25 11:00:00"), SlotEnd: parseTime("2026-05-25 14:00:00"), Capacity: 35, OccupiedCount: 28, Owner: "值班主管李娜", BatchID: "SLOT-20260525", Notes: "午间平峰时段"},
		{Code: "FS-PVG-T2-20260526-I01", Name: "浦东机场T2国际出发晚间时段C", Status: "active", AirportCode: "PVG", Gate: "G75-G82", SlotStart: parseTime("2026-05-26 17:00:00"), SlotEnd: parseTime("2026-05-26 21:00:00"), Capacity: 40, OccupiedCount: 38, Owner: "国际部张伟", BatchID: "SLOT-20260526", Notes: "国际航班集中出发时段"},
		{Code: "FS-CAN-T2-20260527-D01", Name: "白云机场T2国内中转时段D", Status: "active", AirportCode: "CAN", Gate: "A201-A210", SlotStart: parseTime("2026-05-27 10:00:00"), SlotEnd: parseTime("2026-05-27 13:00:00"), Capacity: 30, OccupiedCount: 22, Owner: "中转组陈芳", BatchID: "SLOT-20260527", Notes: "国内中转高峰时段"},
		{Code: "FS-SZX-T3-20260528-N01", Name: "宝安机场T3夜间红眼航班时段E", Status: "active", AirportCode: "SZX", Gate: "N301-N308", SlotStart: parseTime("2026-05-28 22:00:00"), SlotEnd: parseTime("2026-05-29 02:00:00"), Capacity: 25, OccupiedCount: 18, Owner: "夜班组长刘洋", BatchID: "SLOT-20260528", Notes: "红眼航班保障时段"},
		{Code: "FS-TFU-T2-20260529-F01", Name: "天府机场T2首航纪念时段F", Status: "archived", AirportCode: "TFU", Gate: "B101-B108", SlotStart: parseTime("2026-05-29 08:00:00"), SlotEnd: parseTime("2026-05-29 11:00:00"), Capacity: 45, OccupiedCount: 45, Owner: "运营总监赵敏", BatchID: "SLOT-ARCHIVE-001", Notes: "已满员的归档时段"},
		{Code: "FS-XIY-T3-20260530-G01", Name: "咸阳机场T3丝路经济带专享时段G", Status: "suspended", AirportCode: "XIY", Gate: "A501-A510", SlotStart: parseTime("2026-05-30 14:00:00"), SlotEnd: parseTime("2026-05-30 17:00:00"), Capacity: 38, OccupiedCount: 0, Owner: "丝路项目部马超", BatchID: "SLOT-20260530", Notes: "因设备维护暂停开放"},
		{Code: "FS-HGH-T3-20260531-H01", Name: "萧山机场T3亚运会保障预留时段H", Status: "active", AirportCode: "HGH", Gate: "C301-C310", SlotStart: parseTime("2026-05-31 09:00:00"), SlotEnd: parseTime("2026-05-31 12:00:00"), Capacity: 55, OccupiedCount: 35, Owner: "亚运保障组钱进", BatchID: "SLOT-20260531", Notes: "重大活动保障预留时段"},
	}

	for _, slot := range slots {
		database.DB.Exec(
			`INSERT INTO flight_time_slots (code, name, status, airport_code, gate, slot_start, slot_end, capacity, occupied_count, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			slot.Code, slot.Name, slot.Status, slot.AirportCode, slot.Gate,
			slot.SlotStart, slot.SlotEnd, slot.Capacity, slot.OccupiedCount,
			slot.Owner, slot.BatchID, slot.Notes,
		)
	}
}

func insertAppointmentRecords() {
	appointments := []models.AppointmentRecord{
		{Code: "AP-20260525-PEK001", Name: "张明远首都机场T3贵宾厅预约", Status: "confirmed", MemberBenefitID: 1, FlightNumber: "CA1234", FlightDate: "2026-05-25", ScheduledTime: parseTime("2026-05-25 07:30:00"), CompanionCount: 2, Owner: "张明远", BatchID: "APPT-20260525", Notes: "CA1234北京飞上海"},
		{Code: "AP-20260525-PEK002", Name: "李婷婷浦东机场预约申请", Status: "pending_review", MemberBenefitID: 2, FlightNumber: "MU5678", FlightDate: "2026-05-26", ScheduledTime: parseTime("2026-05-26 18:30:00"), CompanionCount: 1, Owner: "李婷婷", BatchID: "APPT-20260526", Notes: "MU5678上海飞东京"},
		{Code: "AP-20260525-CAN003", Name: "王建国广州白云预约确认", Status: "checked_in", MemberBenefitID: 3, FlightNumber: "CZ3456", FlightDate: "2026-05-27", ScheduledTime: parseTime("2026-05-27 11:00:00"), ActualArrivalTime: parseTime("2026-05-27 10:45:00"), CompanionCount: 0, Owner: "王建国", BatchID: "APPT-20260527", Notes: "已签到入场"},
		{Code: "AP-20260525-SZX004", Name: "陈晓华深圳宝安预约", Status: "confirmed", MemberBenefitID: 4, FlightNumber: "ZH9012", FlightDate: "2026-05-28", ScheduledTime: parseTime("2026-05-28 23:30:00"), CompanionCount: 1, Owner: "陈晓华", BatchID: "APPT-20260528", Notes: "夜间航班预约"},
		{Code: "AP-20260525-TFU005", Name: "刘思雨成都天府首次体验", Status: "needs_info", MemberBenefitID: 5, FlightNumber: "CA6789", FlightDate: "2026-05-29", ScheduledTime: parseTime("2026-05-29 09:00:00"), CompanionCount: 0, Owner: "刘思雨", BatchID: "APPT-20260529", Notes: "需要补充身份证明"},
		{Code: "AP-20260525-HGH006", Name: "赵文博杭州萧山过期预约", Status: "cancelled", MemberBenefitID: 6, FlightNumber: "HU2345", FlightDate: "2026-05-21", ScheduledTime: parseTime("2026-05-21 14:00:00"), CompanionCount: 0, Owner: "赵文博", BatchID: "APPT-EXPIRED-001", Notes: "权益过期自动取消"},
		{Code: "AP-20260525-XIY007", Name: "孙丽萍咸阳机场企业预约", Status: "draft", MemberBenefitID: 7, FlightNumber: "MU8901", FlightDate: "2026-06-05", ScheduledTime: parseTime("2026-06-05 15:00:00"), CompanionCount: 3, Owner: "孙丽萍", BatchID: "APPT-20260605", Notes: "企业批量预约草稿"},
		{Code: "AP-20260525-NKG008", Name: "周志强南京禄口VIP预约", Status: "rejected", MemberBenefitID: 8, FlightNumber: "MU3456", FlightDate: "2026-07-10", ScheduledTime: parseTime("2026-07-10 16:00:00"), CompanionCount: 2, Owner: "周志强", BatchID: "APPT-REJECTED-001", Notes: "资料不全被驳回"},
		{Code: "AP-20260525-WUX009", Name: "吴美玲无锡硕放预约", Status: "no_show", MemberBenefitID: 9, FlightNumber: "HO1234", FlightDate: "2026-08-15", ScheduledTime: parseTime("2026-08-15 13:00:00"), CompanionCount: 0, Owner: "吴美玲", BatchID: "APPT-NOSHOW-001", Notes: "未按时到场"},
		{Code: "AP-20260525-DLC010", Name: "郑海涛大连周水子历史预约", Status: "completed", MemberBenefitID: 10, FlightNumber: "CA5678", FlightDate: "2026-03-20", ScheduledTime: parseTime("2026-03-20 10:00:00"), ActualArrivalTime: parseTime("2026-03-20 09:55:00"), CompanionCount: 1, Owner: "郑海涛", BatchID: "APPT-ARCHIVE-001", Notes: "已完成的服务记录"},
		{Code: "AP-20260525-KMG011", Name: "冯雪梅昆明长水违规预约", Status: "cancelled", MemberBenefitID: 11, FlightNumber: "8L9901", FlightDate: "2026-10-20", ScheduledTime: parseTime("2026-10-20 11:00:00"), CompanionCount: 0, Owner: "冯雪梅", BatchID: "APPT-SUSPENDED-001", Notes: "因账户暂停而取消"},
		{Code: "AP-20260525-CKG012", Name: "杨建华重庆江北政府接待", Status: "confirmed", MemberBenefitID: 12, FlightNumber: "CA4567", FlightDate: "2026-11-15", ScheduledTime: parseTime("2026-11-15 14:30:00"), CompanionCount: 4, Owner: "杨建华", BatchID: "APPT-20261115", Notes: "政务接待重要预约"},
		{Code: "AP-20260525-PEK013", Name: "李明北京转机快速通道", Status: "pending_review", MemberBenefitID: 1, FlightNumber: "CA1357", FlightDate: "2026-06-01", ScheduledTime: parseTime("2026-06-01 08:00:00"), CompanionCount: 0, Owner: "李明", BatchID: "APPT-20260601", Notes: "转机快速通道申请"},
		{Code: "AP-20260525-PVG014", Name: "王芳上海浦东商务出行", Status: "confirmed", MemberBenefitID: 2, FlightNumber: "FM9345", FlightDate: "2026-06-10", ScheduledTime: parseTime("2026-06-10 19:00:00"), CompanionCount: 1, Owner: "王芳", BatchID: "APPT-20260610", Notes: "商务出差预约"},
		{Code: "AP-20260525-CAN015", Name: "赵强广州白云家庭出行", Status: "draft", MemberBenefitID: 3, FlightNumber: "CZ6789", FlightDate: "2026-07-01", ScheduledTime: parseTime("2026-07-01 12:00:00"), CompanionCount: 3, Owner: "赵强", BatchID: "APPT-20260701", Notes: "暑期家庭出行草稿"},
	}

	for _, appt := range appointments {
		database.DB.Exec(
			`INSERT INTO appointment_records (code, name, status, member_benefit_id, flight_number, flight_date, scheduled_time, actual_arrival_time, companion_count, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			appt.Code, appt.Name, appt.Status, appt.MemberBenefitID, appt.FlightNumber,
			appt.FlightDate, appt.ScheduledTime, appt.ActualArrivalTime, appt.CompanionCount,
			appt.Owner, appt.BatchID, appt.Notes,
		)
	}
}

func insertCompanions() {
	companions := []models.Companion{
		{Code: "CP-20260525-001", Name: "张小明", Status: "verified", AppointmentRecordID: 1, IDType: "身份证", IDNumber: "110101199001011234", MemberLevel: "basic", RelationType: "子女", Age: 16, Owner: "张明远", BatchID: "CP-20260525", VerificationStatus: "passed", Notes: "未成年子女同行"},
		{Code: "CP-20260525-002", Name: "张小红", Status: "verified", AppointmentRecordID: 1, IDType: "身份证", IDNumber: "110101199505052345", MemberLevel: "basic", RelationType: "配偶", Age: 31, Owner: "张明远", BatchID: "CP-20260525", VerificationStatus: "passed", Notes: "配偶同行人"},
		{Code: "CP-20260525-003", Name: "李大伟", Status: "pending", AppointmentRecordID: 2, IDType: "护照", IDNumber: "E12345678", MemberLevel: "basic", RelationType: "同事", Age: 35, Owner: "李婷婷", BatchID: "CP-20260526", VerificationStatus: "pending", Notes: "商务伙伴同行"},
		{Code: "CP-20260525-004", Name: "陈小芳", Status: "verified", AppointmentRecordID: 4, IDType: "身份证", IDNumber: "440305198808084567", MemberLevel: "basic", RelationType: "朋友", Age: 38, Owner: "陈晓华", BatchID: "CP-20260528", VerificationStatus: "passed", Notes: "朋友同行"},
		{Code: "CP-20260525-005", Name: "刘小军", Status: "rejected", AppointmentRecordID: 7, IDType: "身份证", IDNumber: "610115200001016789", MemberLevel: "basic", RelationType: "下属", Age: 26, Owner: "孙丽萍", BatchID: "CP-20260605", VerificationStatus: "failed", Notes: "年龄超过同行人限制"},
		{Code: "CP-20260525-006", Name: "王小燕", Status: "pending", AppointmentRecordID: 7, IDType: "身份证", IDNumber: "610115199203027890", MemberLevel: "silver", RelationType: "下属", Age: 34, Owner: "孙丽萍", BatchID: "CP-20260605", VerificationStatus: "pending", Notes: "等待核验"},
		{Code: "CP-20260525-007", Name: "赵小刚", Status: "pending", AppointmentRecordID: 7, IDType: "身份证", IDNumber: "610115199512038901", MemberLevel: "basic", RelationType: "客户", Age: 31, Owner: "孙丽萍", BatchID: "CP-20260605", VerificationStatus: "pending", Notes: "客户陪同"},
		{Code: "CP-20260525-008", Name: "周小丽", Status: "verified", AppointmentRecordID: 8, IDType: "护照", IDNumber: "G87654321", MemberLevel: "basic", RelationType: "配偶", Age: 33, Owner: "周志强", BatchID: "CP-REJECTED-001", VerificationStatus: "passed", Notes: "外籍配偶"},
		{Code: "CP-20260525-009", Name: "周小强", Status: "rejected", AppointmentRecordID: 8, IDType: "身份证", IDNumber: "320106199809091012", MemberLevel: "basic", RelationType: "子女", Age: 28, Owner: "周志强", BatchID: "CP-REJECTED-001", VerificationStatus: "failed", Notes: "成年子女超出免费额度"},
		{Code: "CP-20260525-010", Name: "郑小梅", Status: "verified", AppointmentRecordID: 10, IDType: "身份证", IDNumber: "210202198707071123", MemberLevel: "basic", RelationType: "配偶", Age: 39, Owner: "郑海涛", BatchID: "CP-ARCHIVE-001", VerificationStatus: "passed", Notes: "历史同行记录"},
		{Code: "CP-20260525-011", Name: "杨小华", Status: "verified", AppointmentRecordID: 12, IDType: "身份证", IDNumber: "500105197505052345", MemberLevel: "gold", RelationType: "同事", Age: 51, Owner: "杨建华", BatchID: "CP-20261115", VerificationStatus: "passed", Notes: "政务随行人员"},
		{Code: "CP-20260525-012", Name: "杨小军", Status: "verified", AppointmentRecordID: 12, IDType: "身份证", IDNumber: "500105198006063456", MemberLevel: "gold", RelationType: "同事", Age: 46, Owner: "杨建华", BatchID: "CP-20261115", VerificationStatus: "passed", Notes: "政务随行人员"},
		{Code: "CP-20260525-013", Name: "杨小丽", Status: "pending", AppointmentRecordID: 12, IDType: "身份证", IDNumber: "500105199010104567", MemberLevel: "basic", RelationType: "下属", Age: 36, Owner: "杨建华", BatchID: "CP-20261115", VerificationStatus: "pending", Notes: "工作人员"},
		{Code: "CP-20260525-014", Name: "杨小强", Status: "pending", AppointmentRecordID: 12, IDType: "身份证", IDNumber: "500105199211156789", MemberLevel: "basic", RelationType: "司机", Age: 34, Owner: "杨建华", BatchID: "CP-20261115", VerificationStatus: "pending", Notes: "司机人员"},
		{Code: "CP-20260525-015", Name: "李小娟", Status: "verified", AppointmentRecordID: 14, IDType: "身份证", IDNumber: "310115198812126789", MemberLevel: "basic", RelationType: "配偶", Age: 37, Owner: "王芳", BatchID: "CP-20260610", VerificationStatus: "passed", Notes: "商务伴侣"},
		{Code: "CP-20260525-016", Name: "赵小宝", Status: "pending", AppointmentRecordID: 15, IDType: "身份证", IDNumber: "440106201001018901", MemberLevel: "basic", RelationType: "子女", Age: 16, Owner: "赵强", BatchID: "CP-20260701", VerificationStatus: "pending", Notes: "未成年子女"},
		{Code: "CP-20260525-017", Name: "赵小贝", Status: "pending", AppointmentRecordID: 15, IDType: "身份证", IDNumber: "440106201303029012", MemberLevel: "basic", RelationType: "子女", Age: 13, Owner: "赵强", BatchID: "CP-20260701", VerificationStatus: "pending", Notes: "未成年子女"},
		{Code: "CP-20260525-018", Name: "赵小美", Status: "pending", AppointmentRecordID: 15, IDType: "身份证", IDNumber: "440106200505031234", MemberLevel: "basic", RelationType: "子女", Age: 21, Owner: "赵强", BatchID: "CP-20260701", VerificationStatus: "pending", Notes: "成年子女需额外付费"},
	}

	for _, companion := range companions {
		database.DB.Exec(
			`INSERT INTO companions (code, name, status, appointment_record_id, id_type, id_number, member_level, relation_type, age, owner, batch_id, verification_status, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			companion.Code, companion.Name, companion.Status, companion.AppointmentRecordID,
			companion.IDType, companion.IDNumber, companion.MemberLevel, companion.RelationType,
			companion.Age, companion.Owner, companion.BatchID, companion.VerificationStatus, companion.Notes,
		)
	}
}

func insertUsageVouchers() {
	vouchers := []models.UsageVoucher{
		{Code: "UV-QR-20260525001", Name: "张明远贵宾厅入场凭证", Status: "used", VoucherType: "entry", MemberBenefitID: 1, AppointmentRecordID: 1, IssueTime: parseTime("2026-05-25 07:00:00"), ExpireTime: parseTime("2026-05-25 10:00:00"), UseTime: parseTime("2026-05-25 07:35:00"), QRCode: "QR-ENTRY-MB001-AP001-20260525", Owner: "张明远", BatchID: "UV-20260525", Notes: "正常使用入场凭证"},
		{Code: "UV-QR-20260526001", Name: "李婷婷国际出发凭证", Status: "active", VoucherType: "entry", MemberBenefitID: 2, AppointmentRecordID: 2, IssueTime: parseTime("2026-05-26 17:30:00"), ExpireTime: parseTime("2026-05-26 21:00:00"), QRCode: "QR-ENTRY-MB002-AP002-20260526", Owner: "李婷婷", BatchID: "UV-20260526", Notes: "待使用的国际航班凭证"},
		{Code: "UV-QR-20260527001", Name: "王建国已使用凭证", Status: "used", VoucherType: "entry", MemberBenefitID: 3, AppointmentRecordID: 3, IssueTime: parseTime("2026-05-27 10:30:00"), ExpireTime: parseTime("2026-05-27 14:00:00"), UseTime: parseTime("2026-05-27 10:48:00"), QRCode: "QR-ENTRY-MB003-AP003-20260527", Owner: "王建国", BatchID: "UV-20260527", Notes: "已签到使用"},
		{Code: "UV-QR-20260528001", Name: "陈晓华夜间凭证", Status: "active", VoucherType: "entry", MemberBenefitID: 4, AppointmentRecordID: 4, IssueTime: parseTime("2026-05-28 22:30:00"), ExpireTime: parseTime("2026-05-29 02:00:00"), QRCode: "QR-ENTRY-MB004-AP004-20260528", Owner: "陈晓华", BatchID: "UV-20260528", Notes: "夜间航班专用凭证"},
		{Code: "UV-QR-20260529001", Name: "刘思雨首次体验凭证", Status: "expired", VoucherType: "trial", MemberBenefitID: 5, AppointmentRecordID: 5, IssueTime: parseTime("2026-05-29 08:00:00"), ExpireTime: parseTime("2026-05-29 11:00:00"), QRCode: "QR-TRIAL-MB005-AP005-20260529", Owner: "刘思雨", BatchID: "UV-20260529", Notes: "未在有效期内使用"},
		{Code: "UV-QR-20260530001", Name: "赵文博过期作废凭证", Status: "voided", VoucherType: "entry", MemberBenefitID: 6, AppointmentRecordID: 6, IssueTime: parseTime("2026-05-20 13:00:00"), ExpireTime: parseTime("2026-05-20 16:00:00"), QRCode: "QR-VOID-MB006-AP006-20260520", Owner: "赵文博", BatchID: "UV-VOID-001", Notes: "权益过期自动作废"},
		{Code: "UV-QR-20260605001", Name: "孙丽萍企业批量凭证A", Status: "active", VoucherType: "corporate", MemberBenefitID: 7, AppointmentRecordID: 7, IssueTime: parseTime("2026-06-05 14:00:00"), ExpireTime: parseTime("2026-06-05 18:00:00"), QRCode: "QR-CORP-MB007-AP007-20260605-A", Owner: "孙丽萍", BatchID: "UV-20260605", Notes: "企业主凭证"},
		{Code: "UV-QR-20260605002", Name: "孙丽萍企业批量凭证B", Status: "active", VoucherType: "companion", MemberBenefitID: 7, AppointmentRecordID: 7, IssueTime: parseTime("2026-06-05 14:00:00"), ExpireTime: parseTime("2026-06-05 18:00:00"), QRCode: "QR-COMP-MB007-AP007-20260605-B", Owner: "孙丽萍", BatchID: "UV-20260605", Notes: "同行人凭证"},
		{Code: "UV-QR-20260710001", Name: "周志强被驳回凭证", Status: "cancelled", VoucherType: "entry", MemberBenefitID: 8, AppointmentRecordID: 8, IssueTime: parseTime("2026-07-10 15:00:00"), ExpireTime: parseTime("2026-07-10 19:00:00"), QRCode: "QR-CANC-MB008-AP008-20260710", Owner: "周志强", BatchID: "UV-REJECTED-001", Notes: "预约被驳回凭证作废"},
		{Code: "UV-QR-20260815001", Name: "吴美玲未到场凭证", Status: "no_show", VoucherType: "entry", MemberBenefitID: 9, AppointmentRecordID: 9, IssueTime: parseTime("2026-08-15 12:00:00"), ExpireTime: parseTime("2026-08-15 16:00:00"), QRCode:QR-NOSHOW-MB009-AP009-20260815", Owner: "吴美玲", BatchID: "UV-NOSHOW-001", Notes: "未按时到场标记"},
		{Code: "UV-QR-20260320001", Name: "郑海涛历史完成凭证", Status: "archived", VoucherType: "entry", MemberBenefitID: 10, AppointmentRecordID: 10, IssueTime: parseTime("2026-03-20 09:00:00"), ExpireTime: parseTime("2026-03-20 13:00:00"), UseTime: parseTime("2026-03-20 09:55:00"), QRCode: "QR-ARCH-MB010-AP010-20260320", Owner: "郑海涛", BatchID: "UV-ARCHIVE-001", Notes: "已完成归档凭证"},
		{Code: "UV-QR-20261020001", Name: "冯雪梅暂停期间凭证", Status: "frozen", VoucherType: "entry", MemberBenefitID: 11, AppointmentRecordID: 11, IssueTime: parseTime("2026-10-20 10:00:00"), ExpireTime: parseTime("2026-10-20 14:00:00"), QRCode: "QR-FRZN-MB011-AP011-20261020", Owner: "冯雪梅", BatchID: "UV-SUSPENDED-001", Notes: "账户暂停期间冻结"},
		{Code: "UV-QR-20261115001", Name: "杨建华政务主凭证", Status: "active", VoucherType: "government", MemberBenefitID: 12, AppointmentRecordID: 12, IssueTime: parseTime("2026-11-15 13:30:00"), ExpireTime: parseTime("2026-11-15 18:00:00"), QRCode: "QR-GOVN-MB012-AP012-20261115-MAIN", Owner: "杨建华", BatchID: "UV-20261115", Notes: "政务接待主凭证"},
		{Code: "UV-QR-20261115002", Name: "杨建华政务随行凭证A", Status: "active", VoucherType: "companion", MemberBenefitID: 12, AppointmentRecordID: 12, IssueTime: parseTime("2026-11-15 13:30:00"), ExpireTime: parseTime("2026-11-15 18:00:00"), QRCode: "QR-GOVN-MB012-AP012-20261115-COMP-A", Owner: "杨建华", BatchID: "UV-20261115", Notes: "随行人员凭证A"},
		{Code: "UV-QR-20261115003", Name: "杨建华政务随行凭证B", Status: "active", VoucherType: "companion", MemberBenefitID: 12, AppointmentRecordID: 12, IssueTime: parseTime("2026-11-15 13:30:00"), ExpireTime: parseTime("2026-11-15 18:00:00"), QRCode: "QR-GOVN-MB012-AP012-20261115-COMP-B", Owner: "杨建华", BatchID: "UV-20261115", Notes: "随行人员凭证B"},
	}

	for _, voucher := range vouchers {
		database.DB.Exec(
			`INSERT INTO usage_vouchers (code, name, status, voucher_type, member_benefit_id, appointment_record_id, issue_time, expire_time, use_time, qr_code, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			voucher.Code, voucher.Name, voucher.Status, voucher.VoucherType, voucher.MemberBenefitID,
			voucher.AppointmentRecordID, voucher.IssueTime, voucher.ExpireTime, voucher.UseTime,
			voucher.QRCode, voucher.Owner, voucher.BatchID, voucher.Notes,
		)
	}
}

func insertWaitingList() {
	waitingList := []models.WaitingListEntry{
		{Code: "WL-20260525-001", Name: "马小波候补请求", Status: "waiting", MemberBenefitID: 1, PriorityScore: 92.5, QueuePosition: 1, WaitStartTime: parseTime("2026-05-24 15:00:00"), ExpectedEntryTime: parseTime("2026-05-26 07:00:00"), Reason: "原定航班延误需改签", Owner: "马小波", BatchID: "WL-20260525", Notes: "高优先级候补"},
		{Code: "WL-20260525-002", Name: "朱小丽候补申请", Status: "waiting", MemberBenefitID: 2, PriorityScore: 88.0, QueuePosition: 2, WaitStartTime: parseTime("2026-05-24 16:30:00"), ExpectedEntryTime: parseTime("2026-05-26 10:00:00"), Reason: "临时增加商务需求", Owner: "朱小丽", BatchID: "WL-20260525", Notes: "中等优先级"},
		{Code: "WL-20260525-003", Name: "胡小强候补排队", Status: "waiting", MemberBenefitID: 3, PriorityScore: 75.5, QueuePosition: 3, WaitStartTime: parseTime("2026-05-25 08:00:00"), ExpectedEntryTime: parseTime("2026-05-27 11:00:00"), Reason: "名额已满申请候补", Owner: "胡小强", BatchID: "WL-20260525", Notes: "普通候补"},
		{Code: "WL-20260525-004", Name: "郭小芳候补等待", Status: "waiting", MemberBenefitID: 4, PriorityScore: 68.0, QueuePosition: 4, WaitStartTime: parseTime("2026-05-25 09:15:00"), ExpectedEntryTime: parseTime("2026-05-27 15:00:00"), Reason: "团队行程变更", Owner: "郭小芳", BatchID: "WL-20260525", Notes: "低优先级候补"},
		{Code: "WL-20260525-005", Name: "林小刚候补登记", Status: "admitted", MemberBenefitID: 1, PriorityScore: 95.0, QueuePosition: 1, WaitStartTime: parseTime("2026-05-23 10:00:00"), ExpectedEntryTime: parseTime("2026-05-25 06:00:00"), ActualEntryTime: parseTime("2026-05-25 05:45:00"), Reason: "紧急公务出行", Owner: "林小刚", BatchID: "WL-ADMITTED-001", Notes: "已成功入场"},
		{Code: "WL-20260525-006", Name: "何小燕候补成功", Status: "admitted", MemberBenefitID: 2, PriorityScore: 89.5, QueuePosition: 2, WaitStartTime: parseTime("2026-05-23 11:30:00"), ExpectedEntryTime: parseTime("2026-05-25 08:00:00"), ActualEntryTime: parseTime("2026-05-25 07:50:00"), Reason: "有临时取消名额", Owner: "何小燕", BatchID: "WL-ADMITTED-002", Notes: "候补成功转入"},
		{Code: "WL-20260525-007", Name: "罗小军候补取消", Status: "cancelled", MemberBenefitID: 5, PriorityScore: 55.0, QueuePosition: 5, WaitStartTime: parseTime("2026-05-25 12:00:00"), ExpectedEntryTime: parseTime("2026-05-28 09:00:00"), Reason: "主动取消候补", Owner: "罗小军", BatchID: "WL-CANCELLED-001", Notes: "用户主动撤销"},
		{Code: "WL-20260525-008", Name: "梁小梅候补超时", Status: "expired", MemberBenefitID: 6, PriorityScore: 42.0, QueuePosition: 8, WaitStartTime: parseTime("2026-05-20 14:00:00"), ExpectedEntryTime: parseTime("2026-05-25 10:00:00"), Reason: "候补超时自动失效", Owner: "梁小梅", BatchID: "WL-EXPIRED-001", Notes: "超过最大等待时间"},
		{Code: "WL-20260525-009", Name: "宋小华候补排队中", Status: "waiting", MemberBenefitID: 7, PriorityScore: 78.5, QueuePosition: 5, WaitStartTime: parseTime("2026-05-25 10:30:00"), ExpectedEntryTime: parseTime("2026-05-27 14:00:00"), Reason: "特殊日期需求", Owner: "宋小华", BatchID: "WL-20260525", Notes: "企业客户候补"},
		{Code: "WL-20260525-010", Name: "唐小龙候补登记", Status: "waiting", MemberBenefitID: 8, PriorityScore: 62.0, QueuePosition: 6, WaitStartTime: parseTime("2026-05-25 11:45:00"), ExpectedEntryTime: parseTime("2026-05-28 11:00:00"), Reason: "节假日高峰期", Owner: "唐小龙", BatchID: "WL-20260525", Notes: "普通旅客候补"},
	}

	for _, entry := range waitingList {
		database.DB.Exec(
			`INSERT INTO waiting_list (code, name, status, member_benefit_id, priority_score, queue_position, wait_start_time, expected_entry_time, actual_entry_time, reason, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			entry.Code, entry.Name, entry.Status, entry.MemberBenefitID, entry.PriorityScore,
			entry.QueuePosition, entry.WaitStartTime, entry.ExpectedEntryTime, entry.ActualEntryTime,
			entry.Reason, entry.Owner, entry.BatchID, entry.Notes,
		)
	}
}

func insertVerificationResults() {
	results := []models.VerificationResult{
		{Code: "VR-20260525-001", Name: "张明远权益核验通过", Status: "completed", AppointmentRecordID: 1, MemberBenefitID: 1, VerificationTime: parseTime("2026-05-25 07:32:00"), Verifier: "值班员李华", Result: "passed", RuleHits: `[{"rule":"member_level_check","result":"pass"},{"rule":"quota_check","result":"pass"},{"rule":"time_validity","result":"pass"}]`, Score: 98.5, Owner: "张明远", BatchID: "VR-20260525", Notes: "完全符合所有核验规则"},
		{Code: "VR-20260526-001", Name: "李婷婷待核验", Status: "pending", AppointmentRecordID: 2, MemberBenefitID: 2, Verifier: "", Result: "", Score: 0, Owner: "李婷婷", BatchID: "VR-20260526", Notes: "等待现场核验"},
		{Code: "VR-20260527-001", Name: "王建国核验通过", Status: "completed", AppointmentRecordID: 3, MemberBenefitID: 3, VerificationTime: parseTime("2026-05-27 10:49:00"), Verifier: "值班员张伟", Result: "passed", RuleHits: `[{"rule":"member_level_check","result":"pass"},{"rule":"quota_check","result":"pass"}]`, Score: 95.0, Owner: "王建国", BatchID: "VR-20260527", Notes: "提前到场核验通过"},
		{Code: "VR-20260528-001", Name: "陈晓华待核验", Status: "pending", AppointmentRecordID: 4, MemberBenefitID: 4, Verifier: "", Result: "", Score: 0, Owner: "陈晓华", BatchID: "VR-20260528", Notes: "夜间航班待核验"},
		{Code: "VR-20260529-001", Name: "刘思雨信息不完整", Status: "needs_info", AppointmentRecordID: 5, MemberBenefitID: 5, VerificationTime: parseTime("2026-05-29 09:05:00"), Verifier: "值班员王芳", Result: "pending", RejectionReason: "缺少有效身份证明文件", RuleHits: `[{"rule":"identity_verification","result":"fail","reason":"missing_id_document"}]`, Score: 45.0, Owner: "刘思雨", BatchID: "VR-20260529", Notes: "需要补充身份证明"},
		{Code: "VR-20260521-001", Name: "赵文博权益过期拒绝", Status: "completed", AppointmentRecordID: 6, MemberBenefitID: 6, VerificationTime: parseTime("2026-05-21 13:55:00"), Verifier: "系统自动", Result: "failed", RejectionReason: "会员权益已过期", RuleHits: `[{"rule":"validity_check","result":"fail","reason":"benefit_expired"}]`, Score: 0, Owner: "赵文博", BatchID: "VR-EXPIRED-001", Notes: "系统自动检测过期"},
		{Code: "VR-20260605-001", Name: "孙丽萍企业待复核", Status: "pending_review", AppointmentRecordID: 7, MemberBenefitID: 7, Verifier: "", Result: "", Score: 0, Owner: "孙丽萍", BatchID: "VR-20260605", Notes: "企业批量预约待人工复核"},
		{Code: "VR-20260710-001", Name: "周志强资料驳回", Status: "completed", AppointmentRecordID: 8, MemberBenefitID: 8, VerificationTime: parseTime("2026-07-10 15:30:00"), Verifier: "审核主管刘强", Result: "failed", RejectionReason: "企业资质证明文件缺失且同行人数量超标", RuleHits: `[{"rule":"corporate_verification","result":"fail"},{"rule":"companion_limit","result":"fail"}]`, Score: 25.0, Owner: "周志强", BatchID: "VR-REJECTED-001", Notes: "多项规则不通过"},
		{Code: "VR-20260815-001", Name: "吴美玲未到场", Status: "no_show", AppointmentRecordID: 9, MemberBenefitID: 9, VerificationTime: parseTime("2026-08-15 13:05:00"), Verifier: "值班员陈明", Result: "absent", RejectionReason: "预约时间未到场", RuleHits: `[{"rule":"attendance_check","result":"absent"}]`, Score: 0, Owner: "吴美玲", BatchID: "VR-NOSHOW-001", Notes: "标记为未到场"},
		{Code: "VR-20260320-001", Name: "郑海涛历史核验", Status: "archived", AppointmentRecordID: 10, MemberBenefitID: 10, VerificationTime: parseTime("2026-03-20 09:56:00"), Verifier: "值班员赵丽", Result: "passed", RuleHits: `[{"rule":"all_checks","result":"pass"}]`, Score: 100.0, Owner: "郑海涛", BatchID: "VR-ARCHIVE-001", Notes: "历史归档记录"},
		{Code: "VR-20261020-001", Name: "冯雪梅账户暂停", Status: "blocked", AppointmentRecordID: 11, MemberBenefitID: 11, VerificationTime: parseTime("2026-10-20 10:30:00"), Verifier: "风控系统", Result: "blocked", RejectionReason: "账户因违规使用已被暂停", RuleHits: `[{"rule":"account_status","result":"blocked","reason":"account_suspended"}]`, Score: 0, Owner: "冯雪梅", BatchID: "VR-SUSPENDED-001", Notes: "风控拦截"},
		{Code: "VR-20261115-001", Name: "杨建华政务核验", Status: "pending", AppointmentRecordID: 12, MemberBenefitID: 12, Verifier: "", Result: "", Score: 0, Owner: "杨建华", BatchID: "VR-20261115", Notes: "政务接待待核验"},
		{Code: "VR-20260601-001", Name: "李明转机核验", Status: "pending", AppointmentRecordID: 13, MemberBenefitID: 1, Verifier: "", Result: "", Score: 0, Owner: "李明", BatchID: "VR-20260601", Notes: "转机快速通道待核验"},
		{Code: "VR-20260610-001", Name: "王芳商务核验", Status: "pending", AppointmentRecordID: 14, MemberBenefitID: 2, Verifier: "", Result: "", Score: 0, Owner: "王芳", BatchID: "VR-20260610", Notes: "商务出行待核验"},
		{Code: "VR-20260701-001", Name: "赵强家庭核验", Status: "pending", AppointmentRecordID: 15, MemberBenefitID: 3, Verifier: "", Result: "", Score: 0, Owner: "赵强", BatchID: "VR-20260701", Notes: "家庭出行待核验"},
	}

	for _, result := range results {
		database.DB.Exec(
			`INSERT INTO verification_results (code, name, status, appointment_record_id, member_benefit_id, verification_time, verifier, result, rejection_reason, rule_hits, score, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			result.Code, result.Name, result.Status, result.AppointmentRecordID, result.MemberBenefitID,
			result.VerificationTime, result.Verifier, result.Result, result.RejectionReason,
			result.RuleHits, result.Score, result.Owner, result.BatchID, result.Notes,
		)
	}
}

func insertRuleConfigurations() {
	rules := []models.RuleConfiguration{
		{Code: "RULE-001", Name: "会员等级准入规则", Status: "active", RuleType: "eligibility", Priority: 100, ConditionExpression: "member_level IN ['diamond', 'platinum', 'gold', 'silver', 'bronze']", ActionExpression: "ALLOW_ENTRY", Description: "根据会员等级判断是否具备贵宾厅准入资格", Owner: "产品团队", BatchID: "RULE-INIT", Notes: "核心准入规则"},
		{Code: "RULE-002", Name: "剩余额度检查规则", Status: "active", RuleType: "quota", Priority: 90, ConditionExpression: "remaining_quota > 0", ActionExpression: "CHECK_QUOTA", Description: "检查会员剩余使用额度是否充足", Owner: "运营团队", BatchID: "RULE-INIT", Notes: "额度控制规则"},
		{Code: "RULE-003", Name: "同行人数量限制规则", Status: "active", RuleType: "companion", Priority: 80, ConditionExpression: "companion_count <= max_companions_by_level(member_level)", ActionExpression: "VALIDATE_COMPANIONS", Description: "根据会员等级限制同行人数量", Owner: "合规团队", BatchID: "RULE-INIT", Notes: "同行人管控规则"},
		{Code: "RULE-004", Name: "有效期验证规则", Status: "active", RuleType: "validity", Priority: 95, ConditionExpression: "current_date BETWEEN valid_from AND valid_until", ActionExpression: "CHECK_VALIDITY", Description: "验证会员权益是否在有效期内", Owner: "风控团队", BatchID: "RULE-INIT", Notes: "时效性检查规则"},
		{Code: "RULE-005", Name: "航班时段匹配规则", Status: "active", RuleType: "scheduling", Priority: 70, ConditionExpression: "flight_time WITHIN slot_time_range +/- 2hours", ActionExpression: "MATCH_SLOT", Description: "确保预约时间与航班时段合理匹配", Owner: "调度团队", BatchID: "RULE-INIT", Notes: "时段匹配规则"},
		{Code: "RULE-006", Name: "黑名单检查规则", Status: "active", RuleType: "risk_control", Priority: 100, ConditionExpression: "member_id NOT IN blacklist", ActionExpression: "ALLOW_ACCESS", Description: "检查会员是否在黑名单中", Owner: "安全团队", BatchID: "RULE-INIT", Notes: "安全风控规则"},
	}

	for _, rule := range rules {
		database.DB.Exec(
			`INSERT INTO rule_configurations (code, name, status, rule_type, priority, condition_expression, action_expression, description, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			rule.Code, rule.Name, rule.Status, rule.RuleType, rule.Priority,
			rule.ConditionExpression, rule.ActionExpression, rule.Description,
			rule.Owner, rule.BatchID, rule.Notes,
		)
	}
}

func insertExceptionEvents() {
	exceptions := []models.ExceptionEvent{
		{Code: "EX-20260520-001", Name: "赵文博权益过期异常", Status: "resolved", ExceptionType: "benefit_expiry", Severity: "medium", EntityType: "member_benefit", EntityID: 6, TriggerField: "valid_until", ThresholdValue: "2026-05-20 23:59:59", ActualValue: "2026-05-20 23:59:59", Handler: "系统自动", Deadline: parseTime("2026-05-21 00:00:00"), Resolution: "自动标记为过期状态并通知用户", Owner: "赵文博", BatchID: "EX-20260520", Notes: "正常到期处理"},
		{Code: "EX-20260525-001", Name: "刘思雨身份证明缺失异常", Status: "open", ExceptionType: "missing_information", Severity: "low", EntityType: "appointment_record", EntityID: 5, TriggerField: "id_documents", ThresholdValue: "required", ActualValue: "missing", Handler: "值班员王芳", Deadline: parseTime("2026-05-30 00:00:00"), Owner: "刘思雨", BatchID: "EX-20260525", Notes: "首次使用用户缺少身份证明"},
		{Code: "EX-20260527-002", Name: "周志强资料不全异常", Status: "open", ExceptionType: "incomplete_data", Severity: "high", EntityType: "appointment_record", EntityID: 8, TriggerField: "corporate_documents", ThresholdValue: "required", ActualValue: "missing", Handler: "审核主管刘强", Deadline: parseTime("2026-07-15 00:00:00"), Owner: "周志强", BatchID: "EX-20260710", Notes: "企业客户缺少资质证明"},
		{Code: "EX-20260815-001", Name: "吴美玲未到场异常", Status: "closed", ExceptionType: "no_show", Severity: "medium", EntityType: "appointment_record", EntityID: 9, TriggerField: "attendance", ThresholdValue: "present", ActualValue: "absent", Handler: "值班员陈明", Deadline: parseTime("2026-08-16 00:00:00"), Resolution: "标记为未到场并记录信用影响", Owner: "吴美玲", BatchID: "EX-20260815", Notes: "未按时到场处理"},
		{Code: "EX-20261020-001", Name: "冯雪梅账户异常", Status: "investigating", ExceptionType: "policy_violation", Severity: "critical", EntityType: "member_benefit", EntityID: 11, TriggerField: "account_status", ThresholdValue: "normal", ActualValue: "suspended", Handler: "风控专员张磊", Deadline: parseTime("2026-10-25 00:00:00"), Owner: "冯雪梅", BatchID: "EX-20261020", Notes: "疑似违规使用正在调查"},
		{Code: "EX-20260525-002", Name: "首都机场T3容量预警", Status: "resolved", ExceptionType: "capacity_warning", Severity: "high", EntityType: "flight_time_slot", EntityID: 1, TriggerField: "occupancy_rate", ThresholdValue: "< 85%", ActualValue: "84%", Handler: "调度中心", Deadline: parseTime("2026-05-25 06:00:00"), Resolution: "启动候补队列释放名额", Owner: "调度中心", BatchID: "EX-CAPACITY-001", Notes: "接近满员预警"},
		{Code: "EX-20260526-001", Name: "浦东机场国际时段超售风险", Status: "monitoring", ExceptionType: "overbooking_risk", Severity: "medium", EntityType: "flight_time_slot", EntityID: 3, TriggerField: "occupancy_rate", ThresholdValue: "> 95%", ActualValue: "95%", Handler: "国际部张伟", Deadline: parseTime("2026-05-26 17:00:00"), Owner: "国际部", BatchID: "EX-OVERBOOK-001", Notes: "国际航班高峰期监控"},
		{Code: "EX-20260529-003", Name: "天府机场设备维护影响", Status: "scheduled", ExceptionType: "maintenance", Severity: "low", EntityType: "flight_time_slot", EntityID: 7, TriggerField: "facility_status", ThresholdValue: "operational", ActualValue: "maintenance", Owner: "运营团队", BatchID: "EX-MAINT-001", Notes: "计划性维护不影响其他时段"},
	}

	for _, exc := range exceptions {
		database.DB.Exec(
			`INSERT INTO exception_events (code, name, status, exception_type, severity, entity_type, entity_id, trigger_field, threshold_value, actual_value, handler, deadline, resolution, owner, batch_id, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			exc.Code, exc.Name, exc.Status, exc.ExceptionType, exc.Severity, exc.EntityType, exc.EntityID,
			exc.TriggerField, exc.ThresholdValue, exc.ActualValue, exc.Handler, exc.Deadline,
			exc.Resolution, exc.Owner, exc.BatchID, exc.Notes,
		)
	}
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	return t
}
