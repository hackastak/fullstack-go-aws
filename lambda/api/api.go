package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty parameters")
	}

	// check if user already exists
	userExists, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("Internal was an error checking if user exists %w", err)
	}

	// The user already exists
	if userExists {
		return fmt.Errorf("User already exists")
	}

	// The user does not exist, insert the user
	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("Internal error inserting user %w", err)
	}

	return nil
}
