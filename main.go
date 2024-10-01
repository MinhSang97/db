package main

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ChargeHistory struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	BoxID             int       `json:"box_id"`
	OutletID          int       `json:"outlet_id"`
	InvoiceID         string    `json:"invoice_id"`
	ActivationFee     int       `json:"activation_fee"`
	StartTime         int64     `json:"start_time"` // Change to int64 for timestamp
	EndTime           int64     `json:"end_time"`   // Change to int64 for timestamp
	WattageConsumed   float64   `json:"wattage_consumed"`
	Status            int       `json:"status"`
	DiscountAmount    float64   `json:"discount_amount"`
	PromotionCode     string    `json:"promotion_code"`
	PromotionDiscount int       `json:"promotion_discount"`
	TotalConsumedFee  int       `json:"total_consumed_fee"`
	ClientID          string    `json:"client_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ReasonClosed      string    `json:"reason_closed"`
}

// TableName sets the insert table name for this struct type
func (ChargeHistory) TableName() string {
	return "charge_history"
}

func main() {
	// Copy dns and paste here
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//// Query all records
	//var histories []ChargeHistory
	//db.First(&histories)
	//
	//// Print result
	//for _, history := range histories {
	//	println(history)
	//}

	// Create seed for random function based on current time
	rand.Seed(time.Now().UnixNano())

	randomData(db)
}

// Function random data ChargeHistory
func randomData(db *gorm.DB) {
	// Set time range from 2020-01-01 00:00:00 to 2024-01-10 00:00:00
	startRange := time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)
	endRange := time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)

	// Create an array containing times of day: 4h, 10h, and 18h
	timesOfDay := []int{4, 10, 18}

	// Loop through each day in the time range
	for currentDate := startRange; currentDate.Before(endRange); currentDate = currentDate.Add(24 * time.Hour) {
		for _, hour := range timesOfDay {
			// Create startTime with fixed hour (4h, 10h, 18h)
			startTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), hour, 0, 0, 0, currentDate.Location())

			// Calculate endTime by adding 2 hours to startTime
			endTime := startTime.Add(2 * time.Hour)

			// Print startTime and endTime in timestamp
			fmt.Println("Start Time (timestamp):", startTime.Unix())
			fmt.Println("End Time (timestamp):", endTime.Unix())

			// Create ChargeHistory record
			history := ChargeHistory{
				UserID:            1464,
				BoxID:             101,
				OutletID:          1001,
				InvoiceID:         "EVD660393C092AB6",
				ActivationFee:     3000,
				StartTime:         startTime.Unix(),
				EndTime:           endTime.Unix(),
				WattageConsumed:   rand.Float64() * 100, // Value of energy consumption randomly
				Status:            1,
				DiscountAmount:    0.0,
				PromotionCode:     "",
				PromotionDiscount: 0,
				TotalConsumedFee:  rand.Intn(1000000),
				ClientID:          "44E3E887-8A3F-40D3-9022-47201B101E84-1711510464359",
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
				ReasonClosed:      "Notice: Your e-wallet is out of money, please top-up and restart again. Thanks",
			}

			// Add record to database
			db.Create(&history)
			fmt.Println("Inserted record:", history)

			// Query all records
			var histories []ChargeHistory
			db.First(&histories)

			// Print result
			for _, history = range histories {
				fmt.Println(history)
			}
		}
	}
}
