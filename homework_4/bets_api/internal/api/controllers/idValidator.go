package controllers

type IdValidator interface {
	IdIsValid(id string) bool
}
