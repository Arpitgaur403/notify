package main

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// const dsn = "postgres://admin:admin@localhost:5432/postgres?sslmode=disable"

// const createDB = `CREATE DATABASE notify;`

// const createTable = `
// CREATE TABLE IF NOT EXISTS notifications (
//     id         SERIAL PRIMARY KEY,
//     user_id    INT NOT NULL,
//     message    TEXT NOT NULL,
//     status     VARCHAR(20) DEFAULT 'PENDING',
//     created_at TIMESTAMP DEFAULT NOW(),
//     read_at    TIMESTAMP NULL
// );`

// func main() {
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatal("open:", err)
// 	}
// 	defer db.Close()

// 	if err = db.Ping(); err != nil {
// 		log.Fatal("ping:", err)
// 	}

// 	_, err = db.Exec(createDB)
// 	if err != nil {
// 		log.Printf("create db (may already exist): %v", err)
// 	} else {
// 		log.Println("database 'notify' created")
// 	}

// 	notifyDB, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/notify?sslmode=disable")
// 	if err != nil {
// 		log.Fatal("open notify db:", err)
// 	}
// 	defer notifyDB.Close()

// 	if _, err = notifyDB.Exec(createTable); err != nil {
// 		log.Fatal("create table:", err)
// 	}
// 	log.Println("table 'notifications' ready in 'notify' db")
// }
