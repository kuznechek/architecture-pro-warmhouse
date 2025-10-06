package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"database/sql"
	"sensors-telemetry-service/models"

	_ "github.com/lib/pq"
)

// DB represents the database connection
var db *sql.DB

func New(connString string) {
	var err error
	db, err = sql.Open("postgres", connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Create pool failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK!")

	// Test the connection
	if err := db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "unable to ping database: %w", err)
	}
}

func GetSensors(ctx context.Context) ([]models.Sensor, error) {
	query := `
		SELECT id, name, type, location, value, unit, status, last_updated, created_at
		FROM sensors
		ORDER BY id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying sensors: %w", err)
	}
	defer rows.Close()

	var sensors []models.Sensor
	for rows.Next() {
		var s models.Sensor
		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Type,
			&s.Location,
			&s.Value,
			&s.Unit,
			&s.Status,
			&s.LastUpdated,
			&s.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning sensor row: %w", err)
		}
		sensors = append(sensors, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating sensor rows: %w", err)
	}

	return sensors, nil
}

func UpdateSensorValue(ctx context.Context, id int, value float64, status string) error {
	query := `
		UPDATE sensors
		SET value = $1, status = $2, last_updated = $3
		WHERE id = $4
	`

	rows, err := db.Query(query, value, status, time.Now(), id)
	if err != nil {
		return fmt.Errorf("error updating sensor value: %w", err)
	}
	defer rows.Close()

	return nil
}
