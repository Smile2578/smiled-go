// database_test.go

package main

import (
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	ConnectDatabase() // Ensure this is safe to call multiple times

	err := PingDatabase()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
}
