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

func (s *Store) GetFile() *os.File {
	csv_file, err := os.OpenFile(UNITED_CSV_FILENAME, os.O_CREATE|os.O_RDWR, os.ModePerm)

	if err != nil {
		panic(err)
	}

	return csv_file
}

func (s *Store) Read() []models.EventModel {
	csv_file := s.GetFile()
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
	csv_file := s.GetFile()
	writer := csv.NewWriter(csv_file)

	defer csv_file.Close()

	for _, model := range events {
		line := []string{
			model.Uuid,
			model.Name,
			strconv.Itoa(int(model.MinPrice)),
			strconv.Itoa(int(model.MaxPrice)),
			model.NotificationSentAt.Format(time.RFC3339),
		}

		if err := writer.Write(line); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
