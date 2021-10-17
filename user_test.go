package app

import (
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testConfigureDatabase(t)
	d := `{"name":"Clara Clayton"}`
	got, err := CreateUser(context.TODO(), []byte(d))
	if err != nil {
		t.Errorf("CreateUser() error = %v", err)
		return
	}
	logJson(t, got)
}

func TestListUsers(t *testing.T) {
	testConfigureDatabase(t)
	got, err := ListUsers(context.TODO())
	if err != nil {
		t.Errorf("ListUsers() error = %v", err)
		return
	}
	logJson(t, got)
}
