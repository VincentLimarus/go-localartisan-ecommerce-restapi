package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"localArtisans/configs"
	"localArtisans/models/database"
	"log"
	"time"
)

const redisLogQueue = "log_activity_queue"

type LogData struct {
	Action    string    `json:"action"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

func StartLogWorker() {
	client := configs.GetRedis()
	ctx := context.Background()
	db := configs.GetDB()

	for {
		result, err := client.BLPop(ctx, 0, redisLogQueue).Result()
		if err != nil {
			log.Printf("Error fetching from Redis queue: %v", err)
			continue
		}

		if len(result) < 2 {
			continue
		}

		var logData LogData
		err = json.Unmarshal([]byte(result[1]), &logData)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			continue
		}

		logEntry := database.LogActivity{
			Action:    fmt.Sprintf("%s by %s at %s", logData.Action, logData.UpdatedBy, logData.UpdatedAt.Format("2006-01-02 15:04:05")),
			UpdatedBy: logData.UpdatedBy,
			UpdatedAt: logData.UpdatedAt,
		}

		if err := db.Create(&logEntry).Error; err != nil {
			log.Printf("Error inserting log activity: %v", err)
		} else {
			fmt.Printf("Log inserted: %s by %s at %s\n", logData.Action, logData.UpdatedBy, logData.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
	}
}
