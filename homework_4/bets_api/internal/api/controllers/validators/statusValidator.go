package validators

type StatusValidator struct {}

const statusActive = "active"
const statusWon = "won"
const statusLost = "lost"

func NewStatusValidator() *StatusValidator {
	return &StatusValidator{}
}

func (s *StatusValidator) StatusIsValid(status string) bool {
	return status == statusActive || status == statusLost || status == statusWon
}
