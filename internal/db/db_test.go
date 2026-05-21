package db

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

const testDSN = "postgres://arpit:arpit@localhost:5432/test?sslmode=disable"

func testDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("postgres", testDSN)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if err = db.Ping(); err != nil {
		t.Fatalf("ping: %v", err)
	}
	return db
}

func TestConnect(t *testing.T) {
	db := testDB(t)
	defer db.Close()
}

func TestInsertAndQueryNotification(t *testing.T) {
	db := testDB(t)
	defer db.Close()

	var id int
	err := db.QueryRow(`
		INSERT INTO notifications (user_id, message, status)
		VALUES ($1, $2, $3)
		RETURNING id`,
		1, "test message", "PENDING",
	).Scan(&id)
	if err != nil {
		t.Fatalf("insert: %v", err)
	}

	var message, status string
	err = db.QueryRow(`SELECT message, status FROM notifications WHERE id = $1`, id).
		Scan(&message, &status)
	if err != nil {
		t.Fatalf("query: %v", err)
	}

	if message != "test message" {
		t.Errorf("message: got %q, want %q", message, "test message")
	}
	if status != "PENDING" {
		t.Errorf("status: got %q, want %q", status, "PENDING")
	}

	// cleanup
	if _, err = db.Exec(`DELETE FROM notifications WHERE id = $1`, id); err != nil {
		t.Errorf("cleanup: %v", err)
	}
}
