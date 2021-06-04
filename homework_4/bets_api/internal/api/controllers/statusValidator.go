package controllers

type StatusValidator interface {
	StatusIsValid(status string) bool
}
