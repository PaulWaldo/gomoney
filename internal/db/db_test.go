package db

import (
	"testing"
)

func Test_NewSqliteInMemoryServices(t *testing.T) {
	db, err := connectToDatabase("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Unable to create database: %s", err)
	}
	if db != nil {
		t.Fatal("This is a test to see if test on gitlab cause failure")
	}
	// services := domain.Services{Account: NewAccountSvc(db)}
}
