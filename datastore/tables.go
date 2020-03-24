package datastore

import "database/sql"

import "log"

func createTables(db *sql.DB) {
	tables := make([]string, 0)
	// Create configuration table
	tables = append(tables, `create table if not exists config (zen_user text not null, zen_apikey text not null, zen_url text not null, slack_apikey text not null, slack_channel text not null);`)
	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			return
		}
	}
	log.Printf("Postgres tables created successfully.")

}
