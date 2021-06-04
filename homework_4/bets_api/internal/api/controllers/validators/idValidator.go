package validators

type IdValidator struct {}

func NewIdValidator() *IdValidator {
	return &IdValidator{}
}

func (i *IdValidator) IdIsValid(id string) bool {
	return id != ""
}
