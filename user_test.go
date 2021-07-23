package app

import (
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// DB接続Helper
	testConfigureDatabase(t)
	d := `{"organizationID":"00000001-0000-0000-0000-000000000000","name":"Emmett Brown"}`
	got, err := CreateUser(context.TODO(), []byte(d))
	if err != nil {
		t.Errorf("CreateUser() error = %v", err)
		return
	}
	// JSONでログ出力するHelper
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

func TestListUsersAPI(t *testing.T) {
	testConfigureDatabase(t)
	got, err := ListUsersAPI(context.TODO())
	if err != nil {
		t.Errorf("ListUsersAPI() error = %v", err)
		return
	}
	logJson(t, got)
}
