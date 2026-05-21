package main

import (
	"database/sql"
	"log"
	"os/user"

	_ "github.com/lib/pq"
)

const (
	dbUser      = "prachi"
	dbPass      = "prachi"
	dbName      = "prachidb"
	dbHost      = "localhost"
	dbPort      = "5432"
	dbTableName = "notify"
)

func main() {
	// connect as the OS superuser (no password needed via peer/trust locally)
	osUser, err := user.Current()
	if err != nil {
		log.Fatal("get os user:", err)
	}
	superDSN := "postgres://" + osUser.Username + "@" + dbHost + ":" + dbPort + "/postgres?sslmode=disable"

	super, err := sql.Open("postgres", superDSN)
	if err != nil {
		log.Fatal("open super:", err)
	}
	defer super.Close()

	if err = super.Ping(); err != nil {
		log.Fatal("ping super:", err)
	}
	log.Println("connected as superuser")

	// create user if not exists
	var exists bool
	err = super.QueryRow(`SELECT EXISTS(SELECT 1 FROM pg_roles WHERE rolname = $1)`, dbUser).Scan(&exists)
	if err != nil {
		log.Fatal("check user:", err)
	}
	if !exists {
		_, err = super.Exec(`CREATE USER ` + dbUser + ` WITH PASSWORD '` + dbPass + `'`)
		if err != nil {
			log.Fatal("create user:", err)
		}
		log.Printf("user %q created", dbUser)
	} else {
		log.Printf("user %q already exists", dbUser)
	}

	// create database if not exists
	err = super.QueryRow(`SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)`, dbName).Scan(&exists)
	if err != nil {
		log.Fatal("check db:", err)
	}
	if !exists {
		// CREATE DATABASE cannot run inside a transaction, exec directly
		if _, err = super.Exec(`CREATE DATABASE ` + dbName + ` OWNER ` + dbUser); err != nil {
			log.Fatal("create db:", err)
		}
		log.Printf("database %q created", dbName)
	} else {
		log.Printf("database %q already exists", dbName)
	}

	// grant privileges on the database
	if _, err = super.Exec(`GRANT ALL PRIVILEGES ON DATABASE ` + dbName + ` TO ` + dbUser); err != nil {
		log.Fatal("grant db:", err)
	}
	log.Printf("granted privileges on %q to %q", dbName, dbUser)

	// now connect to the target database as superuser to create the table
	targetDSN := "postgres://" + osUser.Username + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	target, err := sql.Open("postgres", targetDSN)
	if err != nil {
		log.Fatal("open target:", err)
	}
	defer target.Close()

	_, err = target.Exec(`
		CREATE TABLE IF NOT EXISTS ` + dbTableName + ` (
			id         SERIAL PRIMARY KEY,
			user_id    INT NOT NULL,
			message    TEXT NOT NULL,
			status     VARCHAR(20) DEFAULT 'PENDING',
			created_at TIMESTAMP DEFAULT NOW(),
			read_at    TIMESTAMP NULL
		)`)
	if err != nil {
		log.Fatal("create table:", err)
	}
	log.Printf("table %q ready", dbTableName)

	for _, stmt := range []string{
		`ALTER TABLE ` + dbTableName + ` OWNER TO ` + dbUser,
		`GRANT ALL PRIVILEGES ON TABLE ` + dbTableName + ` TO ` + dbUser,
		`GRANT ALL PRIVILEGES ON SEQUENCE ` + dbTableName + `_id_seq TO ` + dbUser,
	} {
		if _, err = target.Exec(stmt); err != nil {
			log.Fatal("ownership/grant:", err)
		}
	}
	log.Printf("table ownership and grants set to %q", dbUser)
	log.Printf("setup complete — connect with: postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}
