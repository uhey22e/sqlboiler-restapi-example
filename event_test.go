package app

import (
	"context"
	"testing"
)

func TestListEvents(t *testing.T) {
	testConfigureDatabase(t)
	got, err := ListEvents(context.TODO())
	if err != nil {
		t.Errorf("ListEvent() error = %v", err)
		return
	}
	logJson(t, got)
}

func TestListPopularEvents(t *testing.T) {
	testConfigureDatabase(t)
	got, err := ListPopularEvents(context.TODO())
	if err != nil {
		t.Errorf("ListPopularEvents() error = %v", err)
		return
	}
	logJson(t, got)
}
