package main

import (
	"database/sql"
	"log"
)

// Up is executed when this migration is applied
func Up_20160917010650(txn *sql.Tx) {
	if _, err := txn.Exec("CREATE SEQUENCE IF NOT EXISTS interviews_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;"); err != nil {
		log.Fatalf("index.creation.failed: %v", err)
	}

	if _, err := txn.Exec("CREATE TABLE IF NOT EXISTS interviews (id integer NOT NULL PRIMARY KEY DEFAULT nextval('interviews_id_seq'), expert_name varchar(20), candidate_name varchar(20), price integer)"); err != nil {
		log.Fatalf("table.create.failed: %v", err)
	}

}

// Down is executed when this migration is rolled back
func Down_20160917010650(txn *sql.Tx) {
	if _, err := txn.Exec("DROP TABLE INTERVIEWS"); err != nil {
		log.Fatalf("table.create.failed: %v", err)
	}
	if _, err := txn.Exec("DROP SEQUENCE interviews_id_seq ;"); err != nil {
		log.Fatalf("table.drop_sequence.failed: %v", err)
	}
}
