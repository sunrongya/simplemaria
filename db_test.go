package db

import (
	"database/sql"
	"testing"
)

func TestLocalConnection(t *testing.T) {
	if err := TestConnection(); err != nil {
		t.Errorf(err.Error())
	}
}

var host *sql.DB

func TestNew(t *testing.T) {
	host = New()
}

func TestList(t *testing.T) {
	const (
		listname = "abc123_test_test_test_123abc"
		testdata = "123abc"
	)
	list := NewList(host, listname, "main")
	if err := list.Add(testdata); err != nil {
		t.Errorf("Error, could not add item to list! %s", err.Error())
	}
	items, err := list.GetAll()
	if err != nil {
		t.Errorf("Error when retrieving list! %s", err.Error())
	}
	if len(items) != 1 {
		t.Error("Error, wrong list length!")
	}
	if (len(items) > 0) && (items[0] != testdata) {
		t.Error("Error, wrong list contents!")
	}
	err = list.Remove()
	if err != nil {
		t.Errorf("Error, could not remove list! %s", err.Error())
	}
}

func TestTwoFields(t *testing.T) {
	test, test23, ok := twoFields("test1@test2@test3", "@")
	if ok && ((test != "test1") || (test23 != "test2@test3")) {
		t.Error("Error in twoFields functions")
	}
}
