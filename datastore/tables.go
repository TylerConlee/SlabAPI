package datastore

import "database/sql"

import "log"

func createTables(db *sql.DB) {
	tables := make([]string, 0)
	tables = append(tables, `create table if not exists zendesk (name text not null, apikey text not null, url text not null);`)
	tables = append(tables, `create table if not exists slack (workspace text not null, apikey text not null, channelid text not null);`)
	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			return
		}
	}
	log.Printf("Postgres tables created successfully.")

}
