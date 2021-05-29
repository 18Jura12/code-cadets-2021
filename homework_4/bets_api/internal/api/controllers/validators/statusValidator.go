package validators

type StatusValidator struct {}

func NewStatusValidator() *StatusValidator {
	return &StatusValidator{}
}

func StatusIsValid(status string) bool {
	return status == "active" || status == "lost" || status == "won"
}
