package services_test

import (
	"context"
	"testing"

	"github.com/emwalker/digraph/models"
	"github.com/emwalker/digraph/services"
	"github.com/markbates/goth"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type cleanupFunc func() error

func fetchOrMakeSession(
	t *testing.T, gothUser goth.User,
) (*services.FetchOrMakeSessionSessionResult, cleanupFunc, error) {
	ctx := context.Background()

	tx, err := testDB.Begin()
	result, err := services.FetchOrMakeSession(ctx, tx, gothUser)
	tx.Commit()

	if err != nil {
		return nil, nil, err
	}

	cleanup := func() error {
		if _, err = result.Session.Delete(ctx, testDB); err != nil {
			return err
		}

		if _, err = result.User.Delete(ctx, testDB); err != nil {
			return err
		}

		return nil
	}

	return result, cleanup, nil
}

func TestFetchOrMakeSession(t *testing.T) {
	email := "gnusto@gnusto.com"

	count, err := models.Users(qm.Where("primary_email like ?", email)).Count(context.Background(), testDB)
	if err != nil {
		t.Fatal(err)
	}

	if count > 0 {
		t.Fatalf("Did not expect to find a user with email %s", email)
	}

	gothUser := goth.User{
		Provider: "github",
		NickName: "gnusto",
		Email:    "gnusto@gnusto.com",
	}

	result, cleanup, err := fetchOrMakeSession(t, gothUser)
	defer cleanup()

	if err != nil {
		t.Fatal(err)
	}

	if !result.SessionCreated {
		t.Fatal("Expected a session to be created")
	}

	if !result.UserCreated {
		t.Fatal("Expected a user to be created")
	}
}
