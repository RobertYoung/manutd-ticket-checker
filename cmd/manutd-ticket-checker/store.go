package mutc

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	models "github.com/robertyoung/manutd-ticket-checker/v2/cmd/manutd-ticket-checker/models"
)

type Store struct {
}

const UNITED_CSV_FILENAME = "united_events.csv"

func (s *Store) Read() []models.EventModel {
	csv_file, err := os.Open(UNITED_CSV_FILENAME)

	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(csv_file)
	records, err := reader.ReadAll()

	defer csv_file.Close()

	if err != nil {
		panic(err)
	}

	var events []models.EventModel

	for _, item := range records {
		min_int, _ := strconv.Atoi(item[2])
		max_int, _ := strconv.Atoi(item[3])
		notification_sent_at, _ := time.Parse(time.RFC3339, item[4])

		events = append(events, models.EventModel{
			Uuid:               item[0],
			Name:               item[1],
			MinPrice:           uint16(min_int),
			MaxPrice:           uint16(max_int),
			NotificationSentAt: notification_sent_at,
		})
	}

	return events
}

func (s *Store) Save(events []models.EventModel) {
	csv_file, err := os.Create(UNITED_CSV_FILENAME)

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	writer := csv.NewWriter(csv_file)
	var rows [][]string

	defer csv_file.Close()

	for _, model := range events {
		row := []string{
			model.Uuid,
			model.Name,
			strconv.Itoa(int(model.MinPrice)),
			strconv.Itoa(int(model.MaxPrice)),
			model.NotificationSentAt.Format(time.RFC3339),
		}

		rows = append(rows, row)
	}

	writer.WriteAll(rows)

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
