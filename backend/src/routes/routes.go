package routes

import (
	"airport-vip-lounge/src/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api *fiber.App) {
	memberBenefitHandler := handlers.NewMemberBenefitHandler()
	appointmentHandler := handlers.NewAppointmentRecordHandler()
	flightSlotHandler := handlers.NewFlightTimeSlotHandler()
	companionHandler := handlers.NewCompanionHandler()
	statisticsHandler := handlers.NewStatisticsHandler()
	domainCalcHandler := handlers.NewDomainCalculationHandler()

	api.Get("/health", handlers.HealthCheck)

	memberBenefits := api.Group("/member-benefits")
	memberBenefits.Post("", memberBenefitHandler.Create)
	memberBenefits.Get("", memberBenefitHandler.GetAll)
	memberBenefits.Get("/:id", memberBenefitHandler.GetByID)
	memberBenefits.Put("/:id", memberBenefitHandler.Update)
	memberBenefits.Delete("/:id", memberBenefitHandler.Delete)
	memberBenefits.Post("/:id/transition", memberBenefitHandler.TransitionStatus)
	memberBenefits.Get("/status/:status", memberBenefitHandler.GetByStatus)

	appointments := api.Group("/appointments")
	appointments.Post("", appointmentHandler.Create)
	appointments.Get("", appointmentHandler.GetAll)
	appointments.Get("/:id", appointmentHandler.GetByID)
	appointments.Put("/:id", appointmentHandler.Update)
	appointments.Delete("/:id", appointmentHandler.Delete)
	appointments.Post("/:id/transition", appointmentHandler.TransitionStatus)

	flightSlots := api.Group("/flight-slots")
	flightSlots.Post("", flightSlotHandler.Create)
	flightSlots.Get("", flightSlotHandler.GetAll)
	flightSlots.Get("/:id", flightSlotHandler.GetByID)
	flightSlots.Put("/:id", flightSlotHandler.Update)
	flightSlots.Delete("/:id", flightSlotHandler.Delete)
	flightSlots.Get("/airport/:airport_code", flightSlotHandler.GetByAirportCode)

	companions := api.Group("/companions")
	companions.Post("", companionHandler.Create)
	companions.Get("", companionHandler.GetAll)
	companions.Get("/:id", companionHandler.GetByID)
	companions.Post("/batch", companionHandler.BatchCreate)
	companions.Get("/appointment/:appointment_id", companionHandler.GetByAppointmentID)

	statistics := api.Group("/statistics")
	statistics.Get("", statisticsHandler.GetStatistics)

	domainCalculations := api.Group("/domain-calculations")
	domainCalculations.Get("/priority/:member_benefit_id", domainCalcHandler.CalculatePriority)
	domainCalculations.Get("/consistency/:appointment_id", domainCalcHandler.CheckConsistency)
	domainCalculations.Post("/waiting-list/:slot_id/process", domainCalcHandler.ProcessWaitingList)
	domainCalculations.Get("/expired-benefits", domainCalcHandler.DetectExpiredBenefits)

	api.Post("/seed/reset", handlers.ResetSeedData)
}
