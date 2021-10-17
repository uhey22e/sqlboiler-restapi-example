// Package restapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package restapi

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Event defines model for Event.
type Event struct {
	Date         openapi_types.Date `json:"date" boil:"date"`
	Description  string             `json:"description" boil:"description"`
	Id           string             `json:"id" boil:"id"`
	Name         string             `json:"name" boil:"name"`
	Participants []*User            `json:"participants" boil:"participants"`
}

// PopularEvent defines model for PopularEvent.
type PopularEvent struct {
	Date                       openapi_types.Date `json:"date" boil:"date"`
	Description                string             `json:"description" boil:"description"`
	Id                         string             `json:"id" boil:"id"`
	Name                       string             `json:"name" boil:"name"`
	Participants               int                `json:"participants" boil:"participants"`
	RegistrationsInThePastWeek int                `json:"registrationsInThePastWeek" boil:"registrations_in_the_past_week"`
}

// User defines model for User.
type User struct {
	Id   string `json:"id" boil:"id"`
	Name string `json:"name" boil:"name"`
}

