package db

import (
	"log"
	"testing"
)

func clearTables() {
	log.Print("xx")
	dbConn.Exec("truncate users")
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

//
//func TestMain(m *testing.M) {
//	clearTables()
//	m.Run()
//	clearTables()
//}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUserCredential)
}
