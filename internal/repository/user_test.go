package repository

import (
	"errors"
	"testing"
)

func TestFetchUser_NotFoundWrapsDatabaseError(t *testing.T) {
	r := &UserRepo{}
	_, err := r.FetchUser(7)
	if err == nil {
		t.Fatal("expected error")
	}
	var dbErr *DatabaseError
	if !errors.As(err, &dbErr) {
		t.Fatalf("want *DatabaseError, got %T (%v)", err, err)
	}
	if dbErr.Query == "" {
		t.Error("Query should not be empty")
	}
}

func TestFetchUser_Found(t *testing.T) {
	r := &UserRepo{}
	u, err := r.FetchUser(1)
	if err != nil {
		t.Fatal(err)
	}
	if u.ID != 1 {
		t.Errorf("id = %d", u.ID)
	}
}
