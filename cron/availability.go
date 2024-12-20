// cron/availability.go

package cron

import (
	"log"
	"os"
	"rent-a-girlfriend/db"
	"time"

	"github.com/robfig/cron/v3"
)

func InitAvailabilityCron() {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	c := cron.New(cron.WithLocation(jakartaTime))

	// Run every day at midnight
	_, err := c.AddFunc("0 0 * * *", updateAvailabilityStatus)
	if err != nil {
		panic("Failed to initialize availability cron job: " + err.Error())
	}

	c.Start()
}

func updateAvailabilityStatus() {
	logger := log.New(os.Stdout, "AVAILABILITY-CRON: ", log.LstdFlags)
	now := time.Now()

	logger.Printf("Starting availability update at %v", now)

	err := db.GormDB.Table("availabilities").
		Where("end_date < ? AND is_available = ?", now, false).
		Update("is_available", true).Error

	if err != nil {
		logger.Printf("Error updating availability status: %v", err)
		return
	}

	logger.Printf("Completed availability update")
}
