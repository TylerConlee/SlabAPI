package datastore

import "database/sql"

func createTables(db *sql.DB) {
	tables := make([]string, 0)
	// Create configuration table
	tables = append(tables, `create table if not exists config (zen_user text not null, zen_apikey text not null, zen_url text not null);`)
	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			return
		}
	}
	log.Info("Postgres tables created successfully.")

}
