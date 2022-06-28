package db

import (
	"testing"
)

func Test_NewSqliteInMemoryServicesPopulatesDatabase(t *testing.T) {
	_, err := NewSqliteInMemoryServices()
	if err != nil {
		t.Fatalf("Error creating SqliteInMemoryServices: %s", err)
	}
	
	t.Fatal("Need to use services to verify contents of database")
}
